# bfs

1. find connected graph via a node (floodfill algorithm)
2. topological sort 
3. shorted path in a simple graph 

## Complexity
time complexity: O(n)
space complexity: O(n)

## template 

1. tree or graph? (need remove dup or not?)
2. level order traversal or not? (add one more for loop)
3. using queue 

```go

import (
    "container/list"
    "fmt"
)

type void struct{}
type member struct{}

func bfs(ele *Node) {
	dedup := make(map[Node]void) // dedup if graph

    queue := list.New()
    queue.PushBack(ele)
    dedup[ele] = member
    
    while queue.Len() > 0 {
        level := make([]interface{}, 0)
        size := queue.Len()
        for i := 0; i < size; i ++ {
            head := queue.Front()
            level := append(level, head)
            for _, node := range head.children{
                if node != interface{} && _, exists := set[node]; exists{
                	queue.PushBack(node)
                    dedup[node] = member
                }    
            }
        }
        fmt.Println(level)  // bfs level traversal
    }
}

```

### Note

As golang don't have hashset data structure, using mapset

```go

type void struct{}
var member void

set := make(map[string]void) // New empty set
set["Foo"] = member          // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")      // Delete
size := len(set)        // Size
_, exists := set["Foo"] // Membership

// https://github.com/deckarep/golang-set

```