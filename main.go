package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Fprintln(os.Stderr, ``)
}

func checkVersion(command string, versionflag string, version string) error {
	cmd := exec.Command(command, versionflag)
	bt, err := cmd.Output()
	if err != nil {
		return err
	}

	firstLine := strings.Split(string(bt), "\n")[0]
	pos := strings.LastIndex(firstLine, " ")
	v := strings.TrimSpace(firstLine[pos:])
	if v != version {
		return fmt.Errorf("go version %#v of %s, required version was: %#v", v, command, version)
	}
	return nil
}

func run(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if len(os.Args) < 4 {
		usage()
		os.Exit(1)
		return
	}
	args := os.Args[1:]

	version := args[0]
	versionFlag := args[1]
	command := args[2]
	rest := []string{}
	if len(args) > 3 {
		rest = args[3:]
	}

	err := checkVersion(command, versionFlag, version)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = run(command, rest)
	if err != nil {
		os.Exit(1)
	}

}
