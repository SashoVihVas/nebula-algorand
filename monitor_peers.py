import json
import os
import time
import subprocess

def run_go_program():
    """
    Calls the nebula Go program to crawl the network and generate new results files.
    """
    print("Running nebula crawl command...")
    command = ["nebula", "--json-out", "./results/", "crawl", "--network", "ALGORAND_MAINNET"]
    try:
        result = subprocess.run(command, capture_output=True, text=True, check=True)
        print("Nebula crawl finished successfully.")
        return get_latest_file("results")
    except FileNotFoundError:
        print("\n---")
        print("Error: 'nebula' command not found.")
        print("Please make sure the 'nebula' executable is in your system's PATH,")
        print("or that you have built the program (e.g., 'go build ./cmd/nebula').")
        print("---\n")
        return None
    except subprocess.CalledProcessError as e:
        print(f"\n---")
        print(f"Error running nebula command (return code {e.returncode}):")
        print(f"Stderr: {e.stderr}")
        print(f"---")
        return None
    except Exception as e:
        print(f"An unexpected error occurred while running the Go program: {e}")
        return None

def get_latest_file(directory):
    """Finds the most recently created '*_neighbors.ndjson' file."""
    files = [os.path.join(directory, f) for f in os.listdir(directory) if f.endswith("_neighbors.ndjson")]
    if not files:
        return None
    return max(files, key=os.path.getctime)

def read_peers_to_dict(filepath):
    """
    Reads a neighbors.ndjson file and returns a dictionary
    mapping PeerID to its full JSON data object.
    """
    peers_data = {}
    try:
        with open(filepath, "r") as f:
            for line in f:
                if line.strip():
                    try:
                        data = json.loads(line)
                        peer_id = data.get("PeerID")
                        if peer_id:
                            peers_data[peer_id] = data
                    except json.JSONDecodeError:
                        print(f"Warning: Skipping malformed JSON line in {filepath}: {line.strip()}")
    except FileNotFoundError:
        pass
    return peers_data

def write_peers_from_dict(filepath, peers_data):
    """Writes the peer data from a dictionary to a file in ndjson format."""
    with open(filepath, "w") as f:
        for peer_data in peers_data.values():
            f.write(json.dumps(peer_data) + "\n")

def main():
    """Main function to orchestrate the monitoring and updating process."""
    main_file = "neighbors.ndjson"
    results_dir = "results"
    wait_interval = 60  # seconds to wait between crawls

    if not os.path.exists(results_dir):
        os.makedirs(results_dir)

    # --- Bootstrap Logic ---
    # If the main file doesn't exist, we need to create it first.
    if not os.path.exists(main_file):
        print(f"'{main_file}' not found. Running initial crawl...")
        initial_crawl_file = run_go_program()
        if initial_crawl_file and os.path.exists(initial_crawl_file):
            print(f"Creating '{main_file}' from '{initial_crawl_file}'.")
            with open(initial_crawl_file, "r") as f_in, open(main_file, "w") as f_out:
                f_out.write(f_in.read())

            # Wait for the full interval after the initial crawl before starting the monitor loop.
            print(f"Initial file created. Waiting for {wait_interval} seconds before starting monitoring...")
            time.sleep(wait_interval)
        else:
            print("Could not generate an initial data file. Exiting.")
            return

    # --- Monitoring Loop ---
    consecutive_no_change_runs = 0
    while consecutive_no_change_runs < 3:
        # Now, the first crawl inside this loop is the *second* overall crawl,
        # which is the first meaningful comparison.
        new_crawl_file = run_go_program()

        if not new_crawl_file:
            print(f"Failed to get new results from nebula. Retrying in {wait_interval} seconds...")
            time.sleep(wait_interval)
            continue

        print(f"Comparing '{main_file}' with new results from '{new_crawl_file}'...")

        main_peers_data = read_peers_to_dict(main_file)
        new_peers_data = read_peers_to_dict(new_crawl_file)

        has_changed = False

        for peer_id, new_peer_info in new_peers_data.items():
            # Case 1: The peer is entirely new.
            if peer_id not in main_peers_data:
                print(f"Found new PeerID: {peer_id}. Adding to records.")
                main_peers_data[peer_id] = new_peer_info
                has_changed = True
                continue

            # Case 2: The peer exists. Check for new neighbors in an additive way.
            existing_peer_info = main_peers_data[peer_id]
            existing_neighbors = set(existing_peer_info.get("NeighborIDs", []))
            new_neighbors = set(new_peer_info.get("NeighborIDs", []))

            # Find neighbors that are in the new set but not the existing one.
            added_neighbors = new_neighbors - existing_neighbors

            if added_neighbors:
                print(f"Found {len(added_neighbors)} new neighbor(s) for PeerID: {peer_id}.")
                # Add the new neighbors to the existing list.
                existing_peer_info["NeighborIDs"].extend(list(added_neighbors))
                # Also update the ErrorBits, in case it changed (e.g., from error to success)
                existing_peer_info["ErrorBits"] = new_peer_info.get("ErrorBits", "0")
                has_changed = True

        if has_changed:
            print(f"Network state changed. Updating '{main_file}'.")
            write_peers_from_dict(main_file, main_peers_data)
            consecutive_no_change_runs = 0
        else:
            print("No new peers or neighbors found in this run.")
            consecutive_no_change_runs += 1

        print(f"Consecutive runs with no changes: {consecutive_no_change_runs}")
        if consecutive_no_change_runs < 3:
            # Also updated this to use the variable for consistency.
            print(f"Waiting for {wait_interval} seconds before next crawl...")
            time.sleep(wait_interval)

    print("\n-----------------------------------------------------")
    print("Network stabilized. No additive changes found in 3 consecutive runs.")
    print(f"The final aggregated data is in '{main_file}'.")
    print("-----------------------------------------------------")

if __name__ == "__main__":
    main()