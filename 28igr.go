package main

import (
	"fmt"
	"time"
)

func runningSum(nums []int) []int {
	tmp := make([]int, len(nums))
	var temp int
	copy(tmp, nums)
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		tmp[i] = temp
	}
	return tmp
}

// type person struct {
// 	age     int
// 	name    string
// 	surname string
// }

// type worker struct {
// 	man person
// 	//	jobname  string
// 	location string
// }
type person struct {
	age     int
	name    string
	surname string
}

type worker struct {
	man *person
	//	jobname  string
	location string
}
type student struct {
	creature      *worker
	course        int
	university    string
	Average_score float64
}
type bricklayer struct {
	creature            *worker
	days_without_salary int
	mine                string
}
type butcher struct {
	creature *worker
	needed   bool
	shop     string
}
type doctor struct {
	//creature worker
	cr       *worker
	type_    string
	hospital string
}

func (man *person) birthmet() int {
	currentTime := time.Now()
	year := currentTime.Year()
	return year - man.age
}
func (man *person) petya() {
	fmt.Println("Теперь ты Петя ")
	man.name = "Petya"
	man.surname = "nePetya"
}
func (chel *bricklayer) updateAge(newday int) {
	(*chel).days_without_salary = newday
}
func (receiver *person) Happy_birthday() {
	fmt.Printf("Happy birthday, you were born at %d, so u are %d years old \n", receiver.birthmet(), receiver.age)
}
func createMed(chel worker) (ex doctor) {
	ex.cr = &chel
	ex.type_ = "pediatr"
	return ex
}

func main() {
	nums := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(runningSum(nums))
	// var kamen = bricklayer{
	// 	days_without_salary: 3,
	// 	mine:                "caar",
	// 	creature: worker{
	// 		location: "Kiev",
	// 		man: person{
	// 			age:     20,
	// 			name:    "vita",
	// 			surname: "vota",
	// 		},
	// 	},
	// }
	var chelovek = person{
		age:     28,
		name:    "Kolya",
		surname: "neKolya",
	}
	var work1 worker
	var work2 student
	work1.man = &chelovek
	work1.location = "Lviv"
	work2.creature = &work1
	fmt.Println(work1)
	fmt.Println(work2)

	var chelovek1 = person{
		age:     165,
		name:    "Pasha",
		surname: "nePasha",
	}
	var work3 worker
	var work4 butcher
	work3.man = &chelovek1
	work3.location = "Dnepr"
	//work4.creature.man.Happy_birthday() panic
	work4.creature = &work3
	fmt.Println(chelovek1)
	work4.creature.man.petya()
	fmt.Println(chelovek1)
	work2.creature.man.Happy_birthday()

	chelovek.age = 45

	chelovek.Happy_birthday()
	var test1, test2, test3 worker
	var pudge butcher
	var kamenshik bricklayer
	var uk1, uk2 *butcher
	uk1 = &pudge
	uk2 = &pudge
	fmt.Println(uk1)
	uk1.needed = true
	fmt.Println(uk2)
	test1.location = "dnepr"
	pudge.creature = &test2
	kamenshik.creature = &test3
	test1med := createMed(test1)
	fmt.Println(test1med)
	fmt.Println("---")
	fmt.Println(test1med.cr)
	fmt.Println("---2")
	fmt.Println(test2)
	fmt.Println("---")
	coin := &pudge.creature
	fmt.Println(coin)
	fmt.Println("---")
	ar := []int{1, 3, 5}
	var coins [3]*int
	for i := 0; i < 3; i++ {
		coins[i] = &ar[i]
		fmt.Println("Массив")
		fmt.Println(coins[i])
	}
	// var med doctor
	// //	med := newmedic()
	// med.hospital = "oblast"
	// fmt.Println(med)
	// fmt.Println(kamen)
	// var uk *bricklayer = &kamen
	// //var chel *bricklayer = &chel
	// //fmt.Println(chel)
	// //(*chel).days_without_salary = 16
	// fmt.Println(uk.creature.man.birthmet())
	// uk.updateAge(16)
	// fmt.Println(kamen)
	// var medic doctor
	// fmt.Println(medic)

}
