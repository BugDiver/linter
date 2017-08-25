package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bugdiver/linter/upload"
	"github.com/bugdiver/linter/version"
)

const (
	GOARCH        = "GOARCH"
	GOOS          = "GOOS"
	X86           = "386"
	X86_64        = "amd64"
	darwin        = "darwin"
	pkg           = ".pkg"
	packagesBuild = "packagesbuild"
	spider        = "spider"
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
	executablePath := getGaugeExecutablePath(spider)
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
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--distro":
			createDarwinPackage()
		case "--deploy":
			if err := upload.Upload(); err != nil {
				fmt.Printf("Failed to upload: \n %s \n", err.Error())
			}
		}
	} else {
		setEnv(map[string]string{GOOS: darwin, GOARCH: X86_64})
		compileGauge()
	}
}

func createDarwinPackage() {
	runProcess("rm", "-rf", "deploy")
	runProcess(packagesBuild, "-v", darwinPackageProject)
	runProcess("mv", "build/deploy", "deploy")
	runProcess("mv", "deploy/spider.pkg", "deploy/spider"+version.GetVersion()+".pkg")
}

func getGaugeExecutablePath(file string) string {
	return filepath.Join("bin", "spider")
}
