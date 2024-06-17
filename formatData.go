package main

import (
	"errors"
	"strings"
)

// This function Split the datas in an [][]string and look if the datas is correctly formatted
func GetData(listTetris [][]string, listTemp []string) ([][]string, error) {
	for i := 0; i < len(listTemp); {
		if (i+5 < len(listTemp) && listTemp[i+5] == "") || len(listTemp[i]) > 4 {
			return nil, errors.New("error in the text file")
		}

		listTetris = append(listTetris, listTemp[i:i+4])
		i += 5
	}
	return listTetris, nil
}

// This function changes all the # in the [][]string by a letter
func Rename(listTetris [][]string) {
	for i, l := 0, 'A'; i < len(listTetris); i, l = i+1, l+1 {
		for y := range listTetris[i] {
			listTetris[i][y] = strings.ReplaceAll(listTetris[i][y], "#", string(l))
		}
	}
}

// This function look if the points (in the [][]string) can be removed and if it's possible he remove the points
func DeleteUselessPoint(listTetris [][]string) {
restart:
	for i := range listTetris {
		for y, v := range listTetris[i] {
			if ContientOnlyPointHorizontal(v) {
				if y < len(listTetris[i])-1 {
					listTetris[i] = append(listTetris[i][:y], listTetris[i][y+1:]...)
				} else {
					listTetris[i] = listTetris[i][:y]
				}

				goto restart
			}

			for x := range listTetris[i][y] {
				if ContientOnlyPointVertical(listTetris[i], x) {
					DeleteAllEntry(listTetris[i], x)
					goto restart
				}
			}
		}
	}
}

// This fontion look if the string is only made of point
func ContientOnlyPointHorizontal(s string) bool {
	for _, v := range s {
		if v != '.' {
			return false
		}
	}

	return true
}

// this function look if at the same i position in all the string of the []string it's only points
func ContientOnlyPointVertical(tab []string, i int) bool {
	for _, v := range tab {
		if len(v)-1 < i {
			return false
		}

		if v[i] != '.' {
			return false
		}
	}

	return true
}

// This function delete an y entry of a string
func DeleteAllEntry(tab []string, y int) {
	for i := range tab {
		if len(tab[i]) > y+1 {
			tab[i] = tab[i][:y] + tab[i][y+1:]
		} else {
			tab[i] = tab[i][:y]
		}
	}
}

// Fill tab with points
func FillTabByPoints(n int) []string {
	tab := make([]string, n)

	for i := range tab {
		for y := 0; y < n; y++ {
			tab[i] += "."
		}
	}
	return tab
}
