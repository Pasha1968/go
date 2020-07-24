package main

import (
	"fmt"
)

type person struct {
	age     int
	name    string
	surname string
}

type worker struct {
	man      person
	jobname  string
	location string
}
type actress struct {
	creature     worker
	performances int
	theater      string
}
type artist struct {
	creature         worker
	performances     int
	genre            []string
	famouse_pictures []string
}
type student struct {
	creature      worker
	course        int
	university    string
	Average_score float64
}
type bricklayer struct {
	creature            worker
	days_without_salary int
	mine                string
}
type butcher struct {
	creature worker
	needed   bool
	shop     string
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var artist artist
	fmt.Println(artist.famouse_pictures)
	artist.famouse_pictures = append(artist.famouse_pictures, "name1", "name2")
	fmt.Println(artist.famouse_pictures)
	s := make([]string, 3)
	fmt.Println(s)
	s = append(s, "d")
	fmt.Println(s)
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	x = x[:0]
	fmt.Println(x)
	printSlice(x)

	x = x[:3]
	fmt.Println(x)
	printSlice(x)

	x = []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	printSlice(x)
}
