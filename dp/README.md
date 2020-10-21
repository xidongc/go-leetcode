# dynamic programming 

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
