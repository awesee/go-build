package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	packageName, err := exec.Command("go", "list").Output()
	cmdName := strings.TrimSpace(string(packageName))
	if err != nil || cmdName == "" {
		cmdName, err = os.Getwd()
		checkErr(err)
	}
	cmdName = filepath.Base(cmdName)
	target := map[string][]string{
		"darwin":  {"amd64"},
		"windows": {"amd64"},
		"linux":   {"amd64"},
	}
	checkErr(os.Setenv("CGO_ENABLED", "0"))
	for goos, goarchList := range target {
		checkErr(os.Setenv("GOOS", goos))
		goexe := ""
		if goos == "windows" {
			goexe = ".exe"
		}
		for _, goarch := range goarchList {
			checkErr(os.Setenv("GOARCH", goarch))
			fileName := fmt.Sprintf("%s.%s-%s%s", cmdName, goos, goarch, goexe)
			fmt.Println("OS:", goos, "\tArch:", goarch, "\t", fileName)
			err = exec.Command("go", "build", "-ldflags", "-s -w", "-o", fileName).Run()
			checkErr(err)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
