# sliding window

## solution

consider using hash heap solution, refer `utils/treemap` for data structure implementation
the reason why using hash heap is its remove(index) TP is O(log n), beside if using golang 
1. make sure using heap.Push / heap.Pop instead of &s.Push() / &s.Pop(), the latter one does 
not guarantee ordering
2. s[len(s) - 1] is like s[0] in python, however, before pop, it is not the smallest item 
in heap, make sure use s.Front() to get top ele in heap 
3. use *s to inherit Push and Pop, better Len, Less and Swap

median uses two heap and balance after insert and remove 
max/min uses one heap

hash heap data structure can also be used in line sweep problem eg: skyline 
