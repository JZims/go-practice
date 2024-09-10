package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from Go!"

	file, err := os.Create("./practice/offmain/textFile.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)

	// fmt.Println(os.ReadFile("./textfile.txt"))
	fmt.Printf("Wrote a file with %v characters.\n", length)
	defer file.Close()
	defer readFile("./practice/offmain/textFile.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}

/*
	1.) Create checkError function to call for every creation step
	2.) Create new file with os.Create
	3.) Populate new file with io.Write_
	4.) Defer a function .Close to ensure the file is no longer usable by I/O
	5.) Create a function for reading the file
	6.) Defer a function to call the read file function until after the file is closed

*/
