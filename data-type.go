// data type number in go is 2 type : integer and floating point ("Float")
// data boolean is same : like true and false (here we call bool as boolean) and the representation is same true : false
// data string in go is same : but here we only can use the string with ("") the tag, string here is from 0 - infinity

// the behavior of this string representation we can check the string length and take the char in the index range we choose
package main

import "fmt"

func main() {
	fmt.Println("Integer : ")
	fmt.Println("Return satu = ", 1, "\n")
	fmt.Println("Return dua = ", 4, "\n")
	fmt.Println("Return (9,80) ", 9.80, "\n")

	fmt.Println("\n")

	fmt.Println("Boolean: ")
	fmt.Println("Benar ", true)
	fmt.Println("Salah ", false)

	fmt.Println("\n")

	fmt.Println("String: ")
	fmt.Println("Rama")
	fmt.Println("Ganteng")
	fmt.Println("Fullstack Developer")

	fmt.Println("\n")

	fmt.Println("String Power: ")
	fmt.Println(len("Rama Dita"))
	fmt.Println("Fullstack Developer"[3]) // it's return the ASCI or The Byte

}
