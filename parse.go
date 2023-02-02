package dockerfileparser

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var (
	cmdString = map[string]int{ // 0: invalid, 1: recommended, 2: discard
		"ADD":         1,
		"ARG":         1,
		"CMD":         1,
		"COPY":        1,
		"ENTRYPOINT":  1,
		"ENV":         1,
		"EXPOSE":      1,
		"FROME":       1,
		"HEALTHCHECK": 1,
		"LABEL":       1,
		"MAINTAINER":  1,
		"ONBUILD":     1,
		"RUN":         1,
		"SHELL":       1,
		"STOPSIGNAL":  1,
		"USER":        1,
		"VOLUME":      1,
		"WORKDIR":     1,
	}
)

type command struct {
	Name string
	Args []string
}

func (c *command) add(arg string) {
	c.Args = append(c.Args, c.Name)
}

type DockerFileContent struct {
	all      [][]string
	commands map[string]*command
}

func (d *DockerFileContent) Add(cmd, arg string) {
	if d.commands[cmd] == nil {
		d.commands[cmd] = &command{}
	}
	d.commands[cmd].add(arg)
	d.all = append(d.all, []string{cmd, arg})
}

func isValidCmd(cmd string) bool {
	if cmdString[cmd] == 0 {
		return false
	}
	return true
}

func Parse(f io.Reader) (c DockerFileContent, err error) {
	scanner := bufio.NewScanner(f)
	c = DockerFileContent{commands: make(map[string]*command)}
	for scanner.Scan() {
		cmd, args := parseLine(scanner.Text())
		if isValidCmd(cmd) {
			c.Add(strings.ToUpper(cmd), args)
		}
	}
	err = scanner.Err()
	return
}

func ParseFile(file string) (c DockerFileContent, err error) {
	f, err := os.Open(file)
	if err != nil {
		return DockerFileContent{}, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	return Parse(f)
}
