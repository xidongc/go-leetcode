# linked list

# base problem 
1. reverse linked list
2. delete duplicates
3. merge two linked list

## solutions
1. dummy node 
2. fast/slow pointer **(fast=head.Next)**  

```go 
func midNode(head *ListNode) {
    slow. fast := head, head.Next
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next 
        slow = slow.Next    
    }
    return slow
}
```
 