/*

  Test cases:
    - 42 -> 42
    - "    -42" -> -42 (!!!)
    - "    +42" -> 42 (!!!)
    - "   test -4193 with words" -> 0
    - "4193 with words" -> 4193 (!!!)
    - "    -4193 with words" -> 4193 (!!!)
    - "999999999999999" -> math.MaxInt32 (!!!)
    - "-999999999999999" -> math.MinInt32 (!!!)
    - "" -> 0 (!!!)
    - "-" -> 0 (!!!)
    - "+" -> 0 (!!!)
    - "a" -> 0 (!!!)
  
  Testing grounds:
    "   test -4193 with words"
    
func myAtoi(s string) int {
  s = ignoreLeadingWhitespaces(s)
  hasMinus, s := findSign(s)
  s = readDigits(s)
  return convertToInt(s, hasMinus)
}

func ignoreLeadingWhitespaces(s string) string {
  if len(s) == 0 {
    return s
  }
  
  whitespacesAmount := 0
  
  for i := 0; i < len(s); i++ {
    if s[i] != ' ' {
      break
    }
    
    whitespacesAmount++
  }
  
  return s[whitespacesAmount:]
}

func findSign(s string) (bool, string) {
  if len(s) == 0 {
    return false, s
  }
  
  if s[0] == '-' {
    return true, s[1:]
  } else if s[0] == '+' {
    return false, s[1:]
  } else {
    return false, s
  }
}

func readDigits(s string) string {
  i := 0
  
  for i < len(s) {
    if !isDigit(s[i]) {
      break
    }
    
    i++
  }
  
  return s[0:i]
}

func isDigit(b byte) bool {
  return b == '0' || 
    b == '1' ||
    b == '2' ||
    b == '3' ||
    b == '4' ||
    b == '5' ||
    b == '6' ||
    b == '7' ||
    b == '8' ||
    b == '9';
}


func convertToInt(s string, hasMinus bool) int {
  if s == "" {
    return 0
  }
  
  num := 0
  multiplier := 1
  for i := len(s) - 1; i >= 0; i-- {
    diff := byteToInt(s[i]) * multiplier
    
    if hasMinus {
      if num - diff > num || num - diff <= math.MinInt32 || diff <= math.MinInt32  {
        return math.MinInt32
      }
      num -= diff
    } else {
      if num + diff < num || num + diff >= math.MaxInt32 || diff >= math.MaxInt32 {
        return math.MaxInt32
      }
      num += diff
    }
    
    multiplier *= 10
  }
  
  return num
}

func byteToInt(b byte) int {
  return int(b - '0')
}

*/

// My solution was wrong on the 1000000000000000000000000000000000000000000000000000000000000000000000534534 input, so I gave up, this task is stupid :)

func myAtoi(str string) int {
	var (
		ans   int64
		start bool
		sign  = 1
	)

	for _, v := range str {
		if '0' <= v && v <= '9' {
			start = true
			ans = ans*10 + int64(v-'0')
			if ans >= math.MaxInt32+1 { // avoid int64 overflow
				break
			}
		} else if !start && v == ' ' {
			continue
		} else if !start && v == '+' {
			sign = 1
			start = true
		} else if !start && v == '-' {
			sign = -1
			start = true
		} else {
			break
		}
	}

	ans *= int64(sign)

	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}

	if ans < math.MinInt32 {
		ans = math.MinInt32
	}

	return int(ans)
}












