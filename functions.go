package main

import "strings"

func cleanInput(text string) []string {
	trimText := strings.TrimSpace(text)
	strSlice := strings.Split(strings.ToLower(trimText), " ")
	return strSlice
}
