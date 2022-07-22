package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Upload(fout string, s [][]string) int {

	err := os.WriteFile(fout, []byte(""), 0644) // Stergere fisier
	check(err)

	f, err := os.OpenFile(fout, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	sep := ""

	for _, lin := range s {
		for i, col := range lin {
			if i == len(lin)-1 {
				sep = ""
			} else {
				sep = "#"
			}

			d1 := fmt.Sprintf("%s%s", col, sep)
			_, err = f.WriteString(d1)
			check(err)
		}
		_, err = f.WriteString("\n")
		check(err)
	}

	return 1
}

func Unload(fin string) [][]string {
	var s [][]string

	// dat, err := os.ReadFile(fin)
	// check(err)
	// fmt.Print(string(dat))

	readFile, err := os.Open(fin)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for i := 0; fileScanner.Scan(); i++ {
		s = append(s, strings.Split(fileScanner.Text(), "#"))
	}
	return s
}
