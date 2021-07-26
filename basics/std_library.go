package main

import (
	"fmt"
	"strings"
)

func main() {
	var_list()
}

func var_list() {

	greeting := "hello, how are you? hello, hello"

	fmt.Println("To check if a string exist or not = ", strings.Contains(greeting, "hello"))

	fmt.Println("Replace result = ", strings.ReplaceAll(greeting, "hello", "hi"))

	fmt.Println("The original string = ", greeting)

	fmt.Println("The uppercase result = ", strings.ToUpper(greeting))

	fmt.Println("The index value of h is = ", strings.Index(greeting, "h"))

	fmt.Println("The split value is = ", strings.Split(greeting, ","))

}
