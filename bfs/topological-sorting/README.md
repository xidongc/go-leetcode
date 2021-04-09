# Topological Sorting

## steps

1. create graph, calculate in-degree for each node 
2. get node whose in-degree == 0, and put them in queue
3. bfs, and recursively find nodes in-degree == 0 and put them back to queue 
