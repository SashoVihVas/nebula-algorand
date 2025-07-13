import pandas as pd
import networkx as nx
import matplotlib.pyplot as plt
import json
import os

def find_neighbors_file(directory):
    """
    Finds a file in the given directory that ends with 'neighbors.ndjson'.

    Args:
        directory (str): The directory to search in.

    Returns:
        str: The full path to the found file, or None if not found.
    """
    try:
        for filename in os.listdir(directory):
            if filename.endswith("neighbors.ndjson"):
                return os.path.join(directory, filename)
    except FileNotFoundError:
        print(f"Error: The directory '{directory}' was not found.")
        return None
    return None

def create_network_visualization(file_path):
    """
    Reads a newline-delimited JSON file of network neighbors
    and creates a visual representation of the network topology
    for a limited number of peers.

    Args:
        file_path (str): The path to the .ndjson file.
    """
    if not file_path:
        print("No matching neighbors file found.")
        return

    # Read the ndjson file line by line
    data = []
    print(f"Reading network data from '{file_path}'")
    try:
        with open(file_path, 'r') as f:
            for line in f:
                data.append(json.loads(line))
    except FileNotFoundError:
        print(f"Error: The file '{file_path}' was not found.")
        return

    # Create a DataFrame from the data
    df = pd.DataFrame(data)

    # Create a new directed graph
    G = nx.DiGraph()

    # Add edges to the graph
    for index, row in df.iterrows():
        peer_id = row.get('PeerID')
        if peer_id is None:
            continue  # Skip rows without a PeerID

        # Ensure the PeerID node exists, even if it has no neighbors
        if not G.has_node(peer_id):
            G.add_node(peer_id)

        neighbor_ids = row.get('NeighborIDs', [])
        for neighbor_id in neighbor_ids:
            # Ensure the neighbor node exists
            if not G.has_node(neighbor_id):
                G.add_node(neighbor_id)
            G.add_edge(peer_id, neighbor_id)

    # --- Visualization ---
    if not G.nodes():
        print("Graph is empty. No visualization will be generated.")
        return

    print("Calculating graph layout...")
    plt.figure(figsize=(18, 18))

    # Use a layout that works well for larger networks
    pos = nx.spring_layout(G, k=0.15, iterations=50)

    # Draw the nodes and edges
    nx.draw_networkx_nodes(G, pos, node_size=70, node_color='skyblue')
    nx.draw_networkx_edges(G, pos, edgelist=G.edges(), edge_color='gray', alpha=0.6, arrows=True)

    # Improve plot aesthetics
    plt.title("Network Topology Visualization", size=20)
    plt.axis('off')  # Turn off the axis
    plt.show()


if __name__ == '__main__':
    # Directory where the neighbors file is located
    search_directory = 'results/'

    # Find the specific neighbors file
    neighbors_file = find_neighbors_file(search_directory)

    # Create the visualization if a file was found
    create_network_visualization(neighbors_file)