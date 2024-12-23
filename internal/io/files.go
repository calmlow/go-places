package io

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

const BROWSER = "/mnt/c/Windows/explorer.exe"

func WriteTmpFile(repoFullPath string) {
	content := fmt.Sprintf("TMP_GO_SELECTED_REPO=%s\n", repoFullPath)
	tmpFile := []byte(content)
	err := os.WriteFile("/tmp/selected-repo.txt.tmp", tmpFile, 0644)
	if err != nil {
		panic(err)
	}
}

func CleanRepoPath(repoFullPath string) string {
	path := repoFullPath
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path = strings.Replace(path, "~", home, 1)
	return path
}

func ExecCommand(cmd string) {
	err := exec.Command("bash", "-c", cmd).Start()

	if err != nil {
		log.Fatal(err)
	}
}

func OpenBrowser(url string) {
	cmd := fmt.Sprintf("%s %s", BROWSER, url)
	err := exec.Command("bash", "-c", cmd).Start()

	if err != nil {
		log.Fatal(err)
	}
}

func OpenFileVSCodeBrowser(filename string) {
	cmd := fmt.Sprintf("v %s", filename)
	err := exec.Command("bash", "-c", cmd).Start()

	if err != nil {
		log.Fatal(err)
	}
}

func GetReadmeFileContents(p string) ([]byte, error) {
	p = CleanRepoPath(p)
	readmeFiles := []string{"readme.md", "README.md"}

	for _, filename := range readmeFiles {
		fullPath := path.Join(p, filename)
		log.Println("Looking for fullPath: ", fullPath)
		fileInfo, err := os.Stat(fullPath)
		if err == nil && !fileInfo.IsDir() {
			dat, err := os.ReadFile(fullPath)
			if err != nil {
				return nil, fmt.Errorf("error reading %s: %w", fullPath, err)
			}
			return dat, nil
		}
	}

	return nil, fmt.Errorf("no readme.md or README.md found in %s", p)
}
