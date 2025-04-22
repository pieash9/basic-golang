package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.Mode())
	// fmt.Println(fileInfo.ModTime())

	// read file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 12)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := range buf {
	// 	println("data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(data))

	//  read folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("Hi Go!")
	// f.WriteString("Nice language!")

	// bytes := []byte("Hello go language!")
	// f.Write(bytes)

	// read and write in another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("Written to new file successfully!")

	// delete a  file
	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// delete file
	err := os.Remove("example2.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted!")
}
