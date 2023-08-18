package main

import (
	"fmt"
	"os"
	"os/exec"
)

const sep string = string(os.PathSeparator)

func main(){
	path := os.Getenv("FERNFLOWER_PATH")
	if path == "" {
		fmt.Println("Envitoment variable FERNFLOWER_PATH is not set")
		return
	}
	args := os.Args[1:]
	pArgs := append([]string{"-jar", path}, args...)
	
	proc := exec.Command("java", pArgs...)
	wd, e := os.Getwd()
	if e != nil{
		fmt.Println("Failed to get working dir")
		return
	}
	proc.Dir = wd
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stdout
	proc.Stdin = os.Stdin
	err := proc.Run()
	if err != nil{
		fmt.Println("Failed to run:", err)
	}
	proc.Wait()
}