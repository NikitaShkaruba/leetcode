/*

  Test cases:
    - 3 : III
    - 58 : L V III
    - 1994: M CM XC IV
         ^
      
      1994 : M
      994 : M CM
      94 : M CM XC
      4 : M CM XC IV

      58 : 
      8 : L
      3 : LVII

*/
func intToRoman(num int) string {
	if num <= 0 {
		return ""
	}

	roman := ""
	div := getInitialDiv(num)

	for num != 0 {
		left := int(math.Floor(float64(num/div))) * div

		if left == 0 {
			div /= 10
			continue
		}

		closestRomanInt, closestRomanLetter := getClosestRoman(left)

		num -= closestRomanInt
		roman += closestRomanLetter

		if left == closestRomanInt {
			div /= 10
		}
	}

	return roman
}

func getInitialDiv(num int) int {
	div := 10

	for num > div {
		div *= 10
	}

	return div / 10
}

func getClosestRoman(num int) (int, string) {
	romanInts := []int{1000, 500, 100, 50, 10, 5, 1}
	romanMap := map[int]string{
		1000: "M",
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}

	switch num {
	case 4:
		return 4, "IV"
	case 9:
		return 9, "IX"
	case 40:
		return 40, "XL"
	case 90:
		return 90, "XC"
	case 400:
		return 400, "CD"
	case 900:
		return 900, "CM"
	default:
		for i := 0; i < len(romanInts); i++ {
			if romanInts[i] > num {
				continue
			}

			return romanInts[i], romanMap[romanInts[i]]
		}
	}

	// Todo: log error
	return 0, ""
}















