# bisect 

1. **sorted** array
2. figure out which type of problem (find any/first/last position) 

## Time Complexity

T(n) = T(n/2) + O(1) = O(log n)

## template

1. start + 1 < end
2. mid := start + (end - start) / 2
3. start = mid (last position); 
   end = mid (first position) 
4. first return start (first position); 
   first return end (last position); 
   
```go
mid := start + (end - start) / 2
if nums[mid] == target {
    end = mid // first position for sample
} else if nums[mid] < target {
    start = mid 
} else {
    end = mid
}
if nums[start] == target {
    return start // first position for sample
}
if numsp[end] == target {
    return end
}
return -1 // not found
```