package tests

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	root, err := findProjectRoot()
	if err != nil {
		log.Fatal("failed to find root dir")
	}

	envFile := os.Getenv("ENV_PATH")
	if envFile == "" {
		envFile = ".env"
	}

	envPath := filepath.Join(root, envFile)
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
}
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir || strings.HasSuffix(parentDir, "/") {
			return "", os.ErrNotExist
		}
		dir = parentDir
	}
}
