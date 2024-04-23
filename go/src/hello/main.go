package main

import "fmt"

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// // import  "hello/test"

// // func main() {
// // 	test.Test()
// // }

// // import (
// // 	"fmt"
// // 	"hello/test"
// // )

// // func main() {
// // 	mybill := test.NewBill("mario's bill")
// // 	fmt.Println(mybill)
// // }

// // type Plants struct {
// // 	Name string
// // }

// // func (ple *Plants) Changer(newname string) {
// // 	ple.Name = newname
// // }

// // func main() {
// // 	liliac := Plants{Name: "Liliac"}
// // 	liliac.Changer("Rose")
// // 	fmt.Println(liliac.Name)
// // }

// // func main(){
// // 	var x int
// // 	fmt.Scan(&x)
// // 	fmt.Println(x)
// // }

// // vvvesit imiya fam pasport danniye i stuct polucat

// // type Person struct {
// // 	Name    string
// // 	Surname string
// // 	Age     string
// // }

// // func main() {

// // 	var user Person

// // 	fmt.Scan(&user.Name)
// // 	fmt.Scan(&user.Surname)
// // 	fmt.Scan(&user.Age)

// // 	fmt.Println(user)
// // }

// // func (ple *Plants) Changer(newname string) {
// // 	ple.Name = newname
// // }

// type Person struct {
// 	Name    string
// 	Surname string
// 	Age     string
// }
// func (ple *Person) Cahnger()
// func main() {
// 	var user Person
// 	for true {
// 		var c int
// 		fmt.Println("-----------------------")
// 		fmt.Println("1. Registration")
// 		fmt.Println("2. Edit name")
// 		fmt.Println("3. Show Info")
// 		fmt.Println("0. Exit")
// 		fmt.Println("-----------------------")
// 		fmt.Scan(&c)
// 		switch c {
// 		case 0:
// 			os.Exit(0)
// 		case 1:
// 			fmt.Println("Welcome to registration")
// 			fmt.Println("Write your data\n")
// 			fmt.Print("Name:")
// 			fmt.Scan(&user.Name)
// 			fmt.Print("Surname:")
// 			fmt.Scan(&user.Surname)
// 			fmt.Print("Age:")
// 			fmt.Scan(&user.Age)
// 			fmt.Println(user)
// 		case 2:
// 			fmt.Println("What do you want to change:")
// 		case 3:
// 			fmt.Println("This is three")
// 		default:
// 			fmt.Println("Def")
// 		}
// 	}
// 	fmt.Println("Code finished")
// 	text := fmt.Sprintf("date_%s", time.Now())
// 	fmt.Println(text)
// }

func main() {
	var arr = []int{1, 2, 3}
	arr2 := &arr[0]
	fmt.Println(&arr[0])
	arr = append(*arr2, 4)
	fmt.Println(&arr[0])
}
