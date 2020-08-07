package statistic

func Average(xs []float64) float64 {

	return SumSlice(xs) / float64(len(xs))
}

func SumSlice(xs []float64) float64 {
	var total float64 = 0
	if len(xs) == 0 {
		return total
	}
	for _, x := range xs {
		total += x
	}
	return total
}
