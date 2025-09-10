import json
import os
import time
import subprocess
import argparse


PROFILES = {
    'mainnet': ["--network", "ALGORAND_MAINNET"],
    'testnet': ["--network", "ALGORAND_TESTNET"],
    'suppranet': [
        "--network", "ALGORAND_SUPPRANET",
        "--addr-dial-type", "private"
    ]
}

def run_go_program(network_args, results_dir):
    """
    Calls the nebula Go program to crawl the network and generate new results files.
    'network_args' is a list of strings for the specific crawl command.
    """
    print("Running nebula crawl command...")
    base_command = ["nebula", "--json-out", f"./{results_dir}/", "crawl"]
    command = base_command + network_args

    print(f"Executing: {' '.join(command)}")
    try:
        result = subprocess.run(command, capture_output=True, text=True, check=True)
        print("Nebula crawl finished successfully.")
        return get_latest_file(results_dir)
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
    """Reads a neighbors.ndjson file and returns a dictionary mapping PeerID to its full JSON data object."""
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

    parser = argparse.ArgumentParser(
        description="Monitor a libp2p network by repeatedly crawling it and aggregating peer data."
    )

    subparsers = parser.add_subparsers(dest='network_profile', required=True, help='Available network profiles')

    parser_mainnet = subparsers.add_parser('mainnet', help='Monitor the ALGORAND_MAINNET network')

    parser_testnet = subparsers.add_parser('testnet', help='Monitor the ALGORAND_TESTNET network')

    parser_suppranet = subparsers.add_parser('suppranet', help='Monitor a custom SUPPRANET network')

    parser_suppranet.add_argument(
        '--bootstrap-peers',
        required=True,
        help='A mandatory bootstrap peer string for the suppranet profile (e.g., /dns4/node.example.com/...)',
        type=str
    )

    args = parser.parse_args()
    profile_name = args.network_profile

    network_args = PROFILES[profile_name]

    if profile_name == 'suppranet':
        network_args.extend(["--bootstrap-peers", args.bootstrap_peers])

    timestamp = time.strftime("%Y%m%d_%H%M%S")
    final_output_file = f"{profile_name}_neighbors_{timestamp}.ndjson"

    results_dir = f"results_{profile_name}"
    wait_interval = 60

    print(f"--- Starting monitor for profile: {profile_name} ---")
    print(f"Output file for this run: {final_output_file}")
    print(f"Results dir: {results_dir}")
    print("-------------------------------------------------")

    if not os.path.exists(results_dir):
        os.makedirs(results_dir)


    consecutive_no_change_runs = 0

    aggregated_peers_data = {}

    while consecutive_no_change_runs < 3:
        new_crawl_file = run_go_program(network_args, results_dir)

        if not new_crawl_file:
            print(f"Failed to get new results from nebula. Retrying in {wait_interval} seconds...")
            time.sleep(wait_interval)
            continue

        print(f"Aggregating new results from '{new_crawl_file}'...")

        new_peers_data = read_peers_to_dict(new_crawl_file)

        has_changed = False

        for peer_id, new_peer_info in new_peers_data.items():
            if peer_id not in aggregated_peers_data:
                print(f"Found new PeerID: {peer_id}. Adding to records.")
                aggregated_peers_data[peer_id] = new_peer_info
                has_changed = True
                continue

            existing_peer_info = aggregated_peers_data[peer_id]
            existing_neighbors = set(existing_peer_info.get("NeighborIDs", []))
            new_neighbors = set(new_peer_info.get("NeighborIDs", []))
            added_neighbors = new_neighbors - existing_neighbors

            if added_neighbors:
                print(f"Found {len(added_neighbors)} new neighbor(s) for PeerID: {peer_id}.")
                existing_peer_info["NeighborIDs"].extend(list(added_neighbors))
                existing_peer_info["ErrorBits"] = new_peer_info.get("ErrorBits", "0")
                has_changed = True

        if has_changed:
            print(f"Network state changed. Updating in-memory records.")
            consecutive_no_change_runs = 0
        else:
            print("No new peers or neighbors found in this crawl.")
            consecutive_no_change_runs += 1


        print(f"Writing current aggregated data to '{final_output_file}'.")
        write_peers_from_dict(final_output_file, aggregated_peers_data)

        print(f"Consecutive crawls with no changes: {consecutive_no_change_runs}")
        if consecutive_no_change_runs < 3:
            print(f"Waiting for {wait_interval} seconds before next crawl...")
            time.sleep(wait_interval)

    print("\n-----------------------------------------------------")
    print(f"Network '{profile_name}' stabilized. No additive changes found in 3 consecutive crawls.")
    print(f"The final aggregated data for this run is saved in '{final_output_file}'.")
    print("-----------------------------------------------------")

if __name__ == "__main__":
    main()