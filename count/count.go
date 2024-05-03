package count

func characters(input []rune) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range input {
		counts[char]++
	}

	return counts
}
