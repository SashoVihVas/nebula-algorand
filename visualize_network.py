import pandas as pd
import networkx as nx
import matplotlib.pyplot as plt
import json


def create_network_visualization(file_path):
    """
    Reads a newline-delimited JSON file of network neighbors
    and creates a visual representation of the network topology
    for a limited number of peers.

    Args:
        file_path (str): The path to the .ndjson file.
        peer_limit (int): The maximum number of peers to process.
    """
    # Read the ndjson file line by line
    data = []
    peers_processed = 0
    print(f"Reading network data from '{file_path}'")
    try:
        with open(file_path, 'r') as f:
            for line in f:
                data.append(json.loads(line))
                peers_processed += 1
    except FileNotFoundError:
        print(f"Error: The file '{file_path}' was not found.")
        return

    # Create a DataFrame from the limited data
    df = pd.DataFrame(data)

    # Create a new directed graph
    G = nx.DiGraph()

    # Add edges to the graph
    for index, row in df.iterrows():
        peer_id = row['PeerID']
        # Ensure the PeerID node exists, even if it has no neighbors
        if not G.has_node(peer_id):
            G.add_node(peer_id)
        for neighbor_id in row['NeighborIDs']:
            # Ensure the neighbor node exists
            if not G.has_node(neighbor_id):
                G.add_node(neighbor_id)
            G.add_edge(peer_id, neighbor_id)

    # --- Visualization ---
    print("Calculating graph layout...")
    plt.figure(figsize=(18, 18))

    # Use a layout that works well for larger networks
    pos = nx.spring_layout(G, k=0.15, iterations=50)

    # Draw the nodes and edges
    nx.draw_networkx_nodes(G, pos, node_size=70, node_color='skyblue')
    nx.draw_networkx_edges(G, pos, edgelist=G.edges(), edge_color='gray', alpha=0.6, arrows=True)

    # The line that draws labels has been removed.

    # Improve plot aesthetics
    plt.title("Network Topology Visualization", size=20)
    plt.axis('off')  # Turn off the axis
    plt.show()


if __name__ == '__main__':
    # Assuming the file is in the same directory as the script
    create_network_visualization('results/neighbors.ndjson')