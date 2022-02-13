func threeSum(ints []int) [][]int {
	if ints == nil || len(ints) < 3 {
		return nil
	}

	sort.Ints(ints)

	triplets := make([][]int, 0)

	for i := 0; i < len(ints)-2; i++ {
		if i != 0 && ints[i] == ints[i-1] {
			continue
		}

		j := i + 1
		k := len(ints) - 1

    for j < k && k < len(ints) {
			if j-1 != i && ints[j] == ints[j-1] {
				j += 1
				continue
			}
      if k+1 < len(ints) && ints[k] == ints[k+1] {
				k -= 1
				continue
			}

			sum := ints[i] + ints[j] + ints[k]

			if sum == 0 {
				triplets = append(triplets, []int{ints[i], ints[j], ints[k]})
				j += 1
				continue
			} else if sum < 0 {
				j += 1
			} else {
				k -= 1
			}
		}
	}

	return triplets
}