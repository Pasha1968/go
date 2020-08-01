package main

import (
	"fmt"
)

//Given a non-empty array of integers, every element appears twice except for one. Find that single one.
func singleNumber(nums []int) (ex int) { //8ms 5.6 mp
	m := make(map[int]int, len(nums)/2+1)
	for i := 0; i < len(nums); i++ {
		_, found := m[nums[i]]
		if found == true {
			m[nums[i]] = m[nums[i]] + 1
			// if(m[nums[i]]==2){
			//              delete(m, nums[i])
			//          }
		} else {
			m[nums[i]] = 1
		}
	}
	for key, value := range m {
		if value == 1 {
			ex = key
		}
		// for key := range m {
		//       ex = key
		//   }
	}
	return ex
}

func singleNumber2(nums []int) (ex int) { //8ms 5.6 mp
	for i := 0; i < len(nums); i++ {
		ex ^= nums[i]
	}
	return ex
}
func missingNumber(nums []int) int {
	var ex int
	m := make(map[int]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		m[i] = 0
	}
	fmt.Println(m)
	for i := 0; i < len(nums); i++ {
		_, found := m[nums[i]]
		if found == true {
			m[nums[i]] = 1
		}
	}
	for key, value := range m {
		if value == 0 {
			ex = key
		}
	}
	fmt.Println(m)
	return ex
}

func missingNumbererr(nums []int) (ex int) {
	m := make(map[int]int, len(nums)) //Прокает еррор иногда
	for i := 0; i < len(nums); i++ {
		m[i] = 0
	}
	fmt.Println(m)
	j := 0
	for key, value := range m {
		m[nums[j]] = 1
		if value == 0 {
			ex = key
		}
		j++
	}
	fmt.Println(m)

	return ex
}
func missingNumber2(nums []int) int {
	sum, numsSum := len(nums), 0
	for i := range nums {
		sum += i
		fmt.Println(sum)
		numsSum += nums[i]
		fmt.Println(numsSum)
	}
	return sum - numsSum
}
func missingNumber3(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res = res ^ i ^ nums[i]
	}
	return res ^ len(nums)
}

// func fizzBuzz(n int) []string {
// 	ex := make([]string, n)
// 	for i := 1; i < n+1; i++ {
// 		if i%15 == 0 {
// 			ex[i-1] = "FizzBuzz"
// 		} else if i%5 == 0 {
// 			ex[i-1] = "Buzz"
// 		} else if i%3 == 0 {
// 			ex[i-1] = "Fizz"
// 		} else {
// 			ex[i-1] = strconv.Itoa(i)
// 		}
// 	}
// 	return ex
// }
func main() {
	m := []int{3, 0, 1, 2, 4, 7, 5}
	// o := singleNumber(m)
	// fmt.Println(o)
	fmt.Println(missingNumbererr(m))
}
