# Breadth-first and Depth-first Traversal

This is an example of breadth-first and depth-first traversal in Go. A tree is described in info.json. main.go performs a breath-first and depth-first traversal.

## Run
1. Install [Go](https://golang.org).
2. Run this from the root of the repository:
```
go run main.go
```

Your terminal should output this.

Breadth First

|Data|Distance|
|---|---|---|
|1     |0|
|   2     |1|
|   3     |1|
|   4     |2|
|   5     |2|
|   6     |2|
|   7     |2|

Depth First

|Data|Distance|
|---|---|---|
|   1         |0|
|   3         |1|
|   7         |2|
|   6         |2|
|   2         |1|
|   5         |2|
|   4         |2|

## Discussion

This output represents the breadth-first and depth-first traversals of the tree described in info.json. Each node in the tree has a Data property and each node is some Distance from the root node of the tree. The order in which the nodes are processed by the algorithm is the order in which they are printed.

info.json is an array of objects corresponding to the nodes in the tree. Each node has a "data" property (a unique integer which identifies it) and a "children" property (an array containing the "data" values of its children). The root node must always have a "data" property of 1—the program identifies the root node as having a "data" value of 1.

## Example

This is a very simple tree.

```
[
  {
    "data": 1,
    "children": [2]
  },
  {
    "data": 2,
    "children": []
  }
]
```

It looks like this.
```
    [1]
    /
  [2]
```

`go run main.go` outputs this.

Breadth First

Data|Distance
---|---
   1        | 0
   2        | 1

Depth First

Data|Distance
---|---
   1 |        0
   2 |        1

## Next Steps

Modify the tree in info.json (or make your own and call it info.json) and look at the traversals for more insight into how breadth-first and depth-first patterns differ.
