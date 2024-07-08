package main

import "strings"

func TowerBuilder(nFloors int) []string {
	const star = "*"
	e := []string{}
	for i := 0; i < nFloors; i++ {
		space := strings.Repeat(" ", nFloors-i-1)
		str := strings.Repeat(star, 2*i+1)
		fString := space + str + space
		e = append(e, fString)
	}
	return e
}

// https://www.codewars.com/kata/576757b1df89ecf5bd00073b/train/go
