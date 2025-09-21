import pandas as pd
import networkx as nx
import matplotlib.pyplot as plt
import json
import os

def find_neighbors_file(directory):
    try:
        for filename in os.listdir(directory):
            if filename.endswith("neighbors.ndjson"):
                return os.path.join(directory, filename)
    except FileNotFoundError:
        print(f"Error: The directory '{directory}' was not found.")
        return None
    return None

def create_network_visualization(file_path):
    if not file_path:
        print("No matching neighbors file found.")
        return

    data = []
    print(f"Reading network data from '{file_path}'")
    try:
        with open(file_path, 'r') as f:
            for line in f:
                data.append(json.loads(line))
    except FileNotFoundError:
        print(f"Error: The file '{file_path}' was not found.")
        return

    df = pd.DataFrame(data)

    G = nx.DiGraph()

    for index, row in df.iterrows():
        peer_id = row.get('PeerID')
        if peer_id is None:
            continue

        if not G.has_node(peer_id):
            G.add_node(peer_id)

        neighbor_ids = row.get('NeighborIDs', [])
        for neighbor_id in neighbor_ids:
            if not G.has_node(neighbor_id):
                G.add_node(neighbor_id)
            G.add_edge(peer_id, neighbor_id)

    if not G.nodes():
        print("Graph is empty. No visualization will be generated.")
        return

    print("Calculating graph layout...")
    plt.figure(figsize=(18, 18))

    pos = nx.spring_layout(G, k=0.15, iterations=50)

    nx.draw_networkx_nodes(G, pos, node_size=70, node_color='skyblue')
    nx.draw_networkx_edges(G, pos, edgelist=G.edges(), edge_color='gray', alpha=0.6, arrows=True)

    plt.title("Network Topology Visualization", size=20)
    plt.axis('off')
    plt.show()


if __name__ == '__main__':
    search_directory = 'results/'

    neighbors_file = find_neighbors_file(search_directory)

    create_network_visualization(neighbors_file)