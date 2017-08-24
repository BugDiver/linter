package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	GOARCH        = "GOARCH"
	GOOS          = "GOOS"
	X86           = "386"
	X86_64        = "amd64"
	darwin        = "darwin"
	pkg           = ".pkg"
	packagesBuild = "packagesbuild"
	linter        = "linter"
)

var darwinPackageProject = filepath.Join("build", "linter.pkgproj")

func runProcess(command string, arg ...string) {
	cmd := exec.Command(command, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Execute %v\n", cmd.Args)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func runCommand(command string, arg ...string) (string, error) {
	cmd := exec.Command(command, arg...)
	bytes, err := cmd.Output()
	return strings.TrimSpace(fmt.Sprintf("%s", bytes)), err
}

func compileGauge() {
	executablePath := getGaugeExecutablePath(linter)
	args := []string{
		"build",
		"-ldflags", "-s -w", "-o", executablePath,
	}
	runProcess("go", args...)
}

func setEnv(envVariables map[string]string) {
	for k, v := range envVariables {
		os.Setenv(k, v)
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--distro" {
		createDarwinPackage()
	} else {
		setEnv(map[string]string{GOOS: darwin, GOARCH: X86_64})
		compileGauge()
	}
}

func createDarwinPackage() {
	runProcess(packagesBuild, "-v", darwinPackageProject)
}

func getGaugeExecutablePath(file string) string {
	return filepath.Join(getBinDir(), getExecutableName(file))
}

func getBinDir() string {
	return filepath.Join("bin", fmt.Sprintf("%s_%s", getGOOS(), getGOARCH()))
}

func getExecutableName(file string) string {
	return file
}

func getGOARCH() string {
	goArch := os.Getenv(GOARCH)
	if goArch == "" {
		goArch = runtime.GOARCH
	}
	return goArch
}

func getGOOS() string {
	goOS := os.Getenv(GOOS)
	if goOS == "" {
		goOS = runtime.GOOS
	}
	return goOS
}