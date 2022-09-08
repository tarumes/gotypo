package gotypo

func levensteinDistance(word, compare string) int {
	r1 := []rune(word)
	r2 := []rune(compare)
	var in1len int = len(r1)
	var in2len int = len(r2)
	var cols []int = make([]int, in1len+1)
	for v := 1; v <= in1len; v++ {
		cols[v] = v
	}
	for x := 1; x <= in2len; x++ {
		cols[0] = x
		lastkey := x - 1
		for y := 1; y <= in1len; y++ {
			oldkey := cols[y]
			var incr int
			if r1[y-1] != r2[x-1] {
				incr = 1
			}
			a := cols[y] + 1        // insert
			b := cols[y-1] + 1      // delete
			c := lastkey + incr + 0 // substitution
			cols[y] = c
			if a < b && a < c {
				cols[y] = a
			}
			if b < a && b < c {
				cols[y] = b
			}
			lastkey = oldkey
		}
	}
	return cols[in1len]
}

func rank(input string, word string, result *Correction) {
	var score int = levensteinDistance(input, word)
	if score < result.Rank {
		result.Best = word
		result.Rank = score
	}
	if score == result.Rank && result.Rank <= 1 && result.Best != word {
		result.Same = append(result.Same, word)
	}
	if score == 0 {
		result.Best = word
		result.Rank = score
	}
}
