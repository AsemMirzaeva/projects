package main

import (
	// "fmt"
	"log"
	"os"
)

// var (
// 	newFile *os.File
// 	err error
// )

// func main() {
// 	newFile, err = os.Create("test.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(newFile)
// 	newFile.Close()
// }

// func main() {

// 	err := os.Truncate("test.txt", 100)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// var (
// 	fileInfo os.FileInfo
// 	err error
// )

// func main() {

// 	fileInfo, err = os.Stat("test.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("File name:", fileInfo.Name())
// 	fmt.Println("Size in bytes:", fileInfo.Size())
// 	fmt.Println("Permissions:", fileInfo.Mode())
// }

// func main() {
// 	originalPath := "test.txt"
// 	newPath := "test2.txt"
// 	err := os.Rename(originalPath, newPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	err := os.Remove("test.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	
	file, err := os.Open("text.txt")
}