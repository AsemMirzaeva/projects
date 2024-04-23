package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// func main() {
// 	mybill := test.NewBill("mario's bill")
// 	fmt.Println(mybill)
// }

// type Plants struct {
// 	Name string
// }

// func (ple *Plants) Changer(newname string) {
// 	ple.Name = newname
// }

// func main() {
// 	liliac := Plants{Name: "Liliac"}
// 	liliac.Changer("Rose")
// 	fmt.Println(liliac.Name)
// }

// func main(){
// 	var x int
// 	fmt.Scan(&x)
// 	fmt.Println(x)
// }

// vvvesit imiya fam pasport danniye i stuct polucat

// type Person struct {
// 	Name    string
// 	Surname string
// 	Age     string
// }

// func main() {

// 	var user Person

// 	fmt.Scan(&user.Name)
// 	fmt.Scan(&user.Surname)
// 	fmt.Scan(&user.Age)

// 	fmt.Println(user)
// }

// func (ple *Plants) Changer(newname string) {
// 	ple.Name = newname
// }

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func (ple *Person) Cahnger() {}
func main() {
	file, err := os.OpenFile("file.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var data []Person

	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	for _, datum := range data {
		fmt.Println("Key1:", )
	}
	var user Person

	// if err = dec.Decode(&user); err == io.EOF {
	// 	break
	// } else if err != nil {
	// 	log.Fatal(err)
	// }

	for true {
		var c int
		fmt.Println("-----------------------")
		fmt.Println("1. Registration")
		fmt.Println("2. Edit name")
		fmt.Println("3. Show Info")
		fmt.Println("0. Exit")
		fmt.Println("-----------------------")
		fmt.Scan(&c)
		switch c {
		case 0:
			os.Exit(0)
		case 1:
			fmt.Println("Welcome to registration")
			fmt.Print("Write your data\n")
			fmt.Print("Name:")
			fmt.Scan(&user.Name)
			fmt.Print("Surname:")
			fmt.Scan(&user.Surname)
			fmt.Print("Age:")
			fmt.Scan(&user.Age)

			err = json.NewEncoder(file).Encode(user)
			if err != nil {
				fmt.Println(err)
				return
			}

			// fmt.Println(user)
		case 2:
			fmt.Println(`What do you want to change?
Press:
1.To change the name
2.To change your surname
3.To change age
			`)
			var b int
			fmt.Scan(&b)
			switch b {
			case 1:
				fmt.Println("Write a new name")
				var name string
				fmt.Scan(&name)
				user.ChangerName(name)

			case 2:
				fmt.Println("Write a new surname")
				var surname string
				fmt.Scan(&surname)
				user.ChangersSurname(surname)

			case 3:
				fmt.Println("Write a new age")
				var age int
				fmt.Scan(&age)
				user.ChangerAge(age)
			}
		case 3:
			fmt.Println(user)
		default:
			fmt.Println("Def")
		}
	}
	fmt.Println("Code finished")
}

func (ple *Person) ChangerName(newname string) {
	ple.Name = newname
}

func (ple *Person) ChangersSurname(newSurname string) {
	ple.Surname = newSurname
}

func (ple *Person) ChangerAge(newAge int) {
	ple.Age = newAge
}
