package main

import (
	"fmt"
	"os"
)

func defer_xmpl_1() {

    defer fmt.Println(", World")
    fmt.Print("Hello")
}


func defer_xmpl_2() {

    file := createFile("Test.txt")
    defer closeFile(file)
    // Might be a lot of code here...
    writeFile(file, "Closing this File was deferred!")
}

func createFile(filename string) *os.File {
    
    fmt.Println("Creating file:", filename)
    file, err := os.Create(filename)
    if err != nil {
    
        panic(err)
    }
    
    return file
}

func writeFile(file *os.File, data string) {

    fmt.Println("Writing to file")
    fmt.Fprintln(file, data)
}

func closeFile(file *os.File) {

    fmt.Println("Closing file")
    file.Close()
}