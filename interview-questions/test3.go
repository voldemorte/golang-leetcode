package interviewquestions

// # Day 2 — Median of a Running Data Stream (Two Heaps)
//
// ## Problem:
// ### Design a data structure that supports:
//
// - AddNum(x int): add an integer to the stream
//
// - FindMedian() float64: return the median of all elements seen so far
//
// ### Behavior
//
// - Median for odd count = middle element.
//
// - Median for even count = average of the two middle elements (as float).
//
// ### Examples
//
// ```
// Stream: [1]        -> median = 1.0
// Stream: [1, 2]     -> median = (1+2)/2 = 1.5
// Stream: [1, 2, 3]  -> median = 2.0
// Stream: [2, 3, 4, 5] -> median = (3+4)/2 = 3.5
// ```
// ### Constraints
//
// - Up to ~10^5 operations.
//
// - Values fit in 32-bit signed int.
//
// - Aim for AddNum in O(log n) and FindMedian in O(1).
//
// ### Goal (Go)
// Implement:
// ```
// type MedianFinder struct { /* your fields */ }
//
// func Constructor() MedianFinder
// func (m *MedianFinder) AddNum(num int)
// func (m *MedianFinder) FindMedian() float64
// ```
// ### Hinted Approach (don’t have to follow):
//
// - Use two heaps:
//
//   - low = max-heap for the lower half
//
//   - high = min-heap for the upper half
//
// - Invariant:
//
//   - All elements in low ≤ all elements in high
//
//   - Size balance: len(low) is either equal to len(high) or exactly 1 more
//
// - Add flow (common pattern):
//
//   - Push to low
//
//   - Move top(low) to high
//
//   - Rebalance if len(high) > len(low) by moving top(high) back to low
//
// - Median:
//
//   - If sizes equal → (top(low)+top(high))/2.0
//
//   - Else → float64(top(low))
//
// ### Edge Cases
//
// - Repeated numbers
//
// - Negative numbers
//
// - Large counts (efficiency)
//
// - Alternating small/large insertions (rebalance still holds)
//
// ### Follow-ups (optional)
//
// 1. Support RemoveNum(x int) in near O(log n) (harder; needs lazy deletion or indexed heaps).
//
// 2. If all numbers are known in advance, how to get all medians in O(n log n) vs O(n log k)?
//
// 3. Streaming from stdin: process one number per line and print current median.
