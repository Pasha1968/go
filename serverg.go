package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// type record struct {
// 	records []*person
// }
type record struct {
	records map[int]person
}
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type id struct {
	id int `json:"id"`
}

// func createid(c int) {
// 	c = c + 1
// }
func Create() *record {
	return &record{
		records: make(map[int]person),
	}
}
func (c *record) Update(key int, enter person) {
	c.records[key] = enter
}
func (c *record) Read(key int) person {
	return c.records[key]
}

// func (c record) Add(val *person) {
// 	c.records = append(c.records, val)
// }

var data = Create()
var unic int

func structure(w http.ResponseWriter, r *http.Request) {
	var resp person
	//var rec = Create()
	fmt.Fprintf(w, "Привет ")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "err %q\n", err, err.Error())
	} else {
		switch r.Method {
		case "POST":
			{
				err = json.Unmarshal(body, &resp)
				fmt.Fprintf(w, "Post работает \n")
				if err != nil {
					fmt.Println(w, "can't unmarshal: ", err.Error())
				}
				unic = unic + 1
				data.Update(unic, resp)
				fmt.Println(data.records)
				fmt.Fprintf(w, "id is %d", unic)
				file, _ := json.MarshalIndent(data.records, "", " ")
				_ = ioutil.WriteFile("test.json", file, 0644)
			}
		case "GET":
			{
				fmt.Println(r.URL.Query().Get("id"))
				id, err := strconv.Atoi(r.URL.Query().Get("id"))
				if err != nil {
					fmt.Println(w, "can't transfer ", err.Error())
				}
				fmt.Println(id)
				if id == 0 {
					for i := 1; i < unic; i++ {
						fmt.Fprintf(w, " id:%d \n Name:%s , Age:%d \n", i, data.records[i].Name, data.records[i].Age)
					}
					fmt.Println(data)
				} else {
					fmt.Println(data.records[id])
					fmt.Fprintf(w, "Name:%s , Age:%d", data.records[id].Name, data.records[id].Age)
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/user", structure)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
