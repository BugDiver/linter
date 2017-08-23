package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/bugdiver/linter/templates"
)

type config struct {
	Code       string
	FileName   string
	LintConfig string
}

func createCommand(cmdName string, args ...string) *exec.Cmd {
	cmd := exec.Command(cmdName, args...)
	p, _ := os.Getwd()
	cmd.Dir = p
	return cmd
}

func attachToConsole(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func checkNodeInstallation() {
	cmd := createCommand("node", "-v")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Node is not installed\n %s \n", err.Error())
	}
	cmd = createCommand("npm", "-v")
	if err := cmd.Run(); err != nil {
		fmt.Printf("npm not found in PATH\n %s", err.Error())
	}
}

func installIfESLintNotInstalled() {
	cmd := createCommand("node", "-e", "require(\"eslint\")")
	if err := cmd.Run(); err != nil {
		cmd := createCommand("npm", "install", "eslint", "--prefix", os.Getenv("HOME"))
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to install some packages\n %s \n", err.Error())
		}
	}
}

func readFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("failed to read the file:\n %s", file)
		os.Exit(1)
	}
	return template.JSEscapeString(string(content))
}

func traverseDir(dir string) {
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && filepath.Ext(path) == ".js" {
			code := readFile(path)
			check(code, path)
		}
		return nil
	})
}

func check(code, fileName string) {
	scriptTemplate, _ := template.New("test").Parse(templates.Script)
	codeConfig := config{
		Code: code, FileName: fileName, LintConfig: templates.Config,
	}
	codeString := bytes.NewBufferString("")
	err := scriptTemplate.Execute(codeString, codeConfig)
	if err != nil {
		fmt.Printf("something went wrong\n %s", err.Error())
	}
	cmd := createCommand("node", "-e", codeString.String())
	attachToConsole(cmd)
	cmd.Run()
}

func main() {
	checkNodeInstallation()
	installIfESLintNotInstalled()
	dir, _ := os.Getwd()
	traverseDir(dir)
}
