package medium

func interchangeableRectangles(rectangles [][]int) int64 {
	var cnt int64
	ratios := make(map[float64]int)

	for _, rect := range rectangles {
		ratio := float64(rect[0]) / float64(rect[1])
		ratios[ratio]++
	}

	for _, sum := range ratios {
		cnt += int64(((sum - 1) * sum) / 2)
	}

	return cnt
}
