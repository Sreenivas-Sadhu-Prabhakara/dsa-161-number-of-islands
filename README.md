# Number of Islands

> Striver SDE Sheet 161 · Day 23 Graph — scan every cell; each unsunk land cell is a new island, and DFS sinks the rest of it so it's never recounted. Go, stdlib only, tested.

## Problem

Count the islands — maximal groups of 4-connected `1` cells — in a binary grid. (Diagonal adjacency doesn't connect; the test pins the 5-island checkerboard.)

## Approach

Connected-component counting with the grid as its own visited set: the outer scan increments the count each time it steps on still-standing land, then DFS "sinks" the entire island by overwriting its cells. Every cell is touched a constant number of times. The same count-and-flood skeleton answers provinces, closed islands and enclave problems; the union-find alternative matters only when cells flip dynamically. Note the API honesty: the grid is consumed — documented rather than silently copied, since the copy is the caller's decision.

**Complexity:** O(R·C) time; recursion up to the island size (BFS swaps that for a queue).

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
