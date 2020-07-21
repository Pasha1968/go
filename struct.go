package main

import (
	"fmt"
	"time"
)

type person struct {
	age  int
	name string
}

type worker struct {
	man      person
	jobname  string
	location string
}

func birth(man person) {
	currentTime := time.Now()
	year := currentTime.Year()
	fmt.Println(year - man.age)
}
func (man *person) birthmet() int {
	currentTime := time.Now()
	year := currentTime.Year()
	return year - man.age
}
func main() {
	var me = worker{
		jobname:  "student",
		location: "Dnepr",
		man: person{
			name: "Pasha",
			age:  19,
		},
	}
	fmt.Println(me)

	var he worker
	fmt.Println(he)
	he.man.age = 50
	fmt.Println(he)

	it := worker{jobname: "pet", location: "Kiev"}
	it.man.name = "Pes"
	fmt.Println("it's name:", it.man.name)

	itt := worker{me.man, "tep", "Dnepr"}
	fmt.Println(itt)

	she := new(worker)
	fmt.Println(she.man.age)
	birth(me.man)
	fmt.Println("method:", me.man.birthmet())
}
