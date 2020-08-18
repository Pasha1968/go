package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	//	"strconv"
	//"strings"
)

func getuser(id string) (getuse string) {
	getuse = fmt.Sprintf("http://localhost:8080/user?id=%s", id)
	fmt.Println(getuse)
	return getuse
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	//text = strings.Replace(text, "\n", "", -1)
	fmt.Println(text)
	if len(text) > 5 {
		resp, err := http.Get("http://localhost:8080/user")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok")
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	} else {
		data := []byte(`{"name": "heeeeeeee","age": 24}`)
		r := bytes.NewReader(data)
		resp, err := http.Post("http://localhost:8080/user", "application/json", r)
		fmt.Printf("%v %v", err, resp)
	}
	// data := []byte(`{"name": "hell","age": 24}`)
	// r := bytes.NewReader(data)
	// resp, err := http.Post("http://localhost:8080/user", "application/json", r)
	// fmt.Printf("%v %v", err, resp)
	//id, err := strconv.Atoi(text)
	// teee := fmt.Sprintf("http://localhost:8080/user?id=%s", text)
	// fmt.Println(teee)
	// resp, err := http.Get(teee)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
}
