package main

import "os/exec"
import "os"
import "fmt"
import "strings"

func main() {
	accessLogPath := "C:\\Path\\to\\Improper Quotes Monitor\\access-log.txt"

	myFlags := os.Args[1:]
	myFlagsString := strings.Join(myFlags, " ")
	fmt.Println(myFlags)

	//Alert box
	command := exec.Command("C:\\Windows\\System32\\wscript.exe", "C:\\Path\\to\\Improper Quotes Monitor\\alertbox.vbs")
	err := command.Run()
	if err != nil {
		panic(err)
	}


	//Open access log
	f, err := os.OpenFile(accessLogPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//Write to the file
	if _, err = f.WriteString(myFlagsString + "\r\n"); err != nil {
		panic(err)
	}
}