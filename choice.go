func Choice[T comparable](items map[T]float64) T {
	elements := make([]T, 0, len(items))
	cumWeights := make([]float64, 0, len(items)) // haha cum
	var total float64

	for item, weight := range items {
		elements = append(elements, item)
		total += weight
		cumWeights = append(cumWeights, total)
	}

	r := rand.Float64() * total
	idx := sort.Search(len(cumWeights), func(i int) bool {
		return cumWeights[i] >= r
	})

	return elements[idx]
}
