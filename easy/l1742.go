package easy

func countBalls(lowLimit int, highLimit int) int {
	var boxes, max = make(map[int]int), 0

	for i := lowLimit; i <= highLimit; i++ {
		var box int
		var ball = i

		for ball > 0 {
			box, ball = box+(ball%10), ball/10
		}
		boxes[box]++

		if boxes[box] > max {
			max = boxes[box]
		}
	}

	return max
}
