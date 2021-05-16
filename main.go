package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func readbuf(inp *io.WriteCloser)  {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("scanner work : ",scanner.Text())

		io.WriteString(*inp,scanner.Text() + " custom input \n")

		if scanner.Text() == "q" {
			os.Exit(0)
		}

	}

	if scanner.Err() != nil {
		// handle error.
	}
}

func main() {

	cmd := exec.Command("sh","test_go_exec.sh")

	//change input
	inp,err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("input pipe failed :",err)
		return
	}
	defer inp.Close()
	go readbuf(&inp)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err!=nil {
		log.Fatal("cmd failed ",err)
		return
	}

	fmt.Println("cmd finished")

}

//this function shows how it works basically
func baseUsage(){
	cmd := exec.Command("sh","/Users/jacky/Documents/WorkDir/scripts/test_go_exec.sh")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err!=nil {
		log.Fatal("cmd failed ",err)
	}

	fmt.Println("cmd exit")
}