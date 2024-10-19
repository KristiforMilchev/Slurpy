package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getExecutablePath() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		os.Exit(1)
	}
	return filepath.Dir(exePath)
}

func FileFromExecutable(fileName *string) *string {
	path := filepath.Join(getExecutablePath(), *fileName)
	return &path
}
