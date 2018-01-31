package commands

import (
	"log"
	"os/exec"
	"os"
	"fmt"
	"bufio"
)

func TestCommand(workingDirectory string){
	fmt.Println("Running goconvey ... ")

	if (os.Getenv("library_app_root") == ""){
		fmt.Println("No 'library_app_root' environment variable was found. Setting 'library_app_root' to :",workingDirectory)
		os.Setenv("library_app_root", workingDirectory)
	}
	cmd := exec.Command("goconvey","")

	//get a pipe so that we can scan the output and display it to users
	cmdReader, err := cmd.StdoutPipe()
	if (err != nil){
		log.Fatal(err)
	}

	//scan and relay the information back to the user
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("goconvey | %s\n", scanner.Text())
		}
	}()

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}