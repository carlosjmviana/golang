# Depth First Search

Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures. The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph) and explores as far as possible along each branch before backtracking [1].

The graph below shows the order in which nodes are visited in DFS:

[![](https://mermaid.ink/img/pako:eNpV0MkKwjAQBuBXCXP6Cxbcl97U1rrd9JhLMNEWukhND1L67qahFTK5DB__MEwaepRSUUCvSrwTdg95wUxtgYnnMd_32Q6Yel7PVvbAypUQWA-ysxIBM1cOwHKQ0EoMbFw5mrX_ZZGlEzB35QwsBomtXMzY2KWroUlH3WOMRpSrKhepNKc2nXDSicoVp8C0Uj1FnWlOvGhNtH5LoVUkU11WFDxF9lEjErUub9_iQYGuajWEwlSYn8v7VPsDS5BTNw)](https://mermaid.live/edit#pako:eNpV0MkKwjAQBuBXCXP6Cxbcl97U1rrd9JhLMNEWukhND1L67qahFTK5DB__MEwaepRSUUCvSrwTdg95wUxtgYnnMd_32Q6Yel7PVvbAypUQWA-ysxIBM1cOwHKQ0EoMbFw5mrX_ZZGlEzB35QwsBomtXMzY2KWroUlH3WOMRpSrKhepNKc2nXDSicoVp8C0Uj1FnWlOvGhNtH5LoVUkU11WFDxF9lEjErUub9_iQYGuajWEwlSYn8v7VPsDS5BTNw)

## Pseudocode

### Recursive
```
procedure DFS(G, v) is
    label v as discovered
    for all directed edges from v to w that are in G.adjacentEdges(v) do
        if vertex w is not labeled as discovered then
            recursively call DFS(G, w)
```

### Iterative
```
procedure DFS_iterative(G, v) is
    let S be a stack
    S.push(v)
    while S is not empty do
        v = S.pop()
        if v is not labeled as discovered then
            label v as discovered
            for all edges from v to w in G.adjacentEdges(v) do 
                S.push(w)
```

## Implementation
For simplicity, will use an undirected and unweighted graph as the following example:

```
  edges = [[1,2],[1,3], [2,4]], where edges[i] = [x,y] there is and edged between nodes x and y.
```

We will represent the graph as an adjacency list


### Iterative
[Iterative Version](./README.md)

### Recursive
* [Recursive Version](./dfs_recursive.go)
* [Recursive Version with cyclic check](./dfs_recursive_with_cycle_check.go)

# References
* [1] - [Wikipedia.org - Depth-First Search](https://en.wikipedia.org/wiki/Depth-first_search)
* [2] - [Techie Delight - Depth First Search](https://www.techiedelight.com/depth-first-search/)