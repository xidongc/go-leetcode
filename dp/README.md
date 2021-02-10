# dynamic programming 

1. coordinate dp eg: maximal square 
2. sequence dp: f[i] means pre i ele, eg: word break II 
3. double sequence dp:d[i][j] means rel pre i ele in str1 and pre j ele in str2, eg: 44 wildcard match
4. range dp eg: palindrome pre-calculate
5. game theory dp: dp[i] means status when i ele left eg: coins in a line

```gotemplate
for length := 2; length < size; length ++ {
    for start := 0; start + length < size; start ++ {
    	dp[start][start+length] = ...
    }
}
```

## dp **_vs_** memorized dfs 
turn dp into a dfs memorization process eg: 44 wildcard matching 

```gotemplate
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
memorized dfs can't use space complexity optimization below

## optimize space complexity 
1. if f[i] only has rel with f[i-1] then replace with f[i%2], f[(i-1)%2]
