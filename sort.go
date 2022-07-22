package main

import (
	"sort"
	"strconv"
)

func SortFile(f string) error {
	s := Unload(f)
	sort.SliceStable(s, func(i, j int) bool {
		x, _ := strconv.Atoi(s[i][den.DurationValue])
		y, _ := strconv.Atoi(s[j][den.DurationValue])
		return x < y
	})
	Upload(f, s)
	return nil
}
