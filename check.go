package main

import "strings"

// Delete last form, recursivly
func DeleteLastForm(tab, form []string) []string {
	letter := GetLetter(form)
	for i := range tab {
		tab[i] = strings.ReplaceAll(tab[i], letter, ".")
	}
	return tab
}

// Get letters of the actual forms

func GetLetter(tab []string) string {
	for _, i := range tab {
		for _, v := range i {
			if v != '.' {
				return string(v)
			}
		}
	}
	return ""
}

// Count all letters
func CountNbrPoint(tab []string, s rune) int {
	count := 0
	for _, i := range tab {
		for _, y := range i {
			if y == s {
				count++
			}
		}
	}

	return count
}

// This function verify tab have I quantity of S

func TabContain(tab []string, s string, i int) bool {
	n := 0
	for _, y := range tab {
		for _, v := range y {
			if string(v) == s {
				n++
			}
		}
	}
	
	return n == i
}
