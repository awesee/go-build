package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	checkErr(err)
	cmdName := filepath.Base(dir)
	target := map[string][]string{
		"darwin":  {"amd64"},
		"windows": {"amd64"},
		"linux":   {"amd64"},
	}
	err = os.Setenv("CGO_ENABLED", "0")
	checkErr(err)
	for goos, goarchList := range target {
		err = os.Setenv("GOOS", goos)
		checkErr(err)
		goexe := ""
		if goos == "windows" {
			goexe = ".exe"
		}
		for _, goarch := range goarchList {
			err = os.Setenv("GOARCH", goarch)
			checkErr(err)
			fileName := fmt.Sprintf("%s.%s-%s%s", cmdName, goos, goarch, goexe)
			fmt.Println("OS:", goos, "\tArch:", goarch, "\t", fileName)
			err = exec.Command("go", "build", "-ldflags", "-s -w", "-o", fileName).Run()
			checkErr(err)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
