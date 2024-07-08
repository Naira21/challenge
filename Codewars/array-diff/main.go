package main

func ArrayDiff(a, b []int) []int {
	bMap := make(map[int]bool)

	for _, v := range b {
		bMap[v] = true
	}

	var t []int
	for _, w := range a {
		if !bMap[w] {
			t = append(t, w)
		}
	}

	return t
}

// https://www.codewars.com/kata/523f5d21c841566fde000009/train/go
