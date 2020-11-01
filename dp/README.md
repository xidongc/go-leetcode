# dynamic programming 

## dp **_vs_** memorized dfs 
turn dp into a dfs memorization process eg: 44 wildcard matching 

```go
func isMatch(s string, p string) bool { 
    return dfs_helper(len(s), len(p), s, p, map[[2]int]bool{})
}

func dfs_helper(i, j int, s string, p string, memo map[[2]int]bool) bool {
    if value, exist := memo[[2]int{i, j}]; exist {
        return value
    }
    if j == 0 && i == 0 {
        return true 
    } else if j == 0 {
        return false 
    } else if i == 0 {
        result := dfs_helper(i, j-1, s, p,memo) && string(p[j-1]) == "*"
        memo[[2]int{i, j}] = result
        return result
    }
    
    if s[i-1] == p[j-1] {
        result := dfs_helper(i-1, j-1, s, p,memo)
        memo[[2]int{i, j}] = result
        return result
    } else if string(p[j-1]) == "?" {
        result := dfs_helper(i-1, j-1, s, p,memo)
        memo[[2]int{i, j}] = result
        return result
    } else if string(p[j-1]) == "*" {
        result := dfs_helper(i-1, j, s, p,memo) || dfs_helper(i-1, j-1, s, p, memo) || dfs_helper(i, j-1, s, p,memo)
        memo[[2]int{i, j}] = result
        return result
    }
    return false
} 
```

## range 
eg: palindrome pre-calculate
```go
dp[i][i]
dp[i][i+1]
for length := 2; length < size; i ++ {
    for start := 0; start + length < size; start ++ {
    	dp[i][j] = ...
    }
}
```
