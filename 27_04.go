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
	theater      string `example:"Grand_Theatre"`
}
type artist struct {
	creature         worker
	performances     int
	genre            []string `example:"genre?"`
	famouse_pictures []string `example:"picture_name"`
}
type student struct {
	creature      worker
	course        int
	university    string `example:"DNU"`
	Average_score float64
}
type bricklayer struct {
	creature            worker
	days_without_salary int `example:"3"`
	mine                string `example:"mine_name"`
}
type butcher struct {
	creature worker
	needed   bool   `example:"true"`
	shop     string `example:"shopname"`
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

	x = []int{2, 3, 4, 5, 6, 7, 8, 99}
	fmt.Println(x)
	printSlice(x)
}
