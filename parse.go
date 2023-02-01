package dockerfileparser

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const (
	AddString         = "ADD"
	ArgString         = "ARG"
	CmdString         = "CMD"
	CopyString        = "COPY"
	EntrypointString  = "ENTRYPOINT"
	EnvString         = "ENV"
	ExposeString      = "EXPOSE"
	FromString        = "FROME"
	HealthcheckString = "HEALTHCHECK"
	LabelString       = "LABEL"
	MaintainerString  = "MAINTAINER"
	OnbuildString     = "ONBUILD"
	RunString         = "RUN"
	ShellString       = "SHELL"
	StopSignalString  = "STOPSIGNAL"
	UserString        = "USER"
	VolumeString      = "VOLUME"
	WorkdirString     = "WORKDIR"
)

type DockerFileContent struct {
	All         [][]string
	Add         []string
	Arg         []string
	Cmd         []string
	Copy        []string
	Entrypoint  []string
	Env         []string
	Expose      []string
	From        []string
	Healthcheck []string
	Label       []string
	Maintainer  []string
	Onbuild     []string
	Run         []string
	Shell       []string
	StopSignal  []string
	User        []string
	Volume      []string
	Workdir     []string
}

func Parse(f io.Reader) (c DockerFileContent, err error) {
	scanner := bufio.NewScanner(f)
	c = DockerFileContent{}
	for scanner.Scan() {
		cmd, args := parseLine(scanner.Text())
		switch strings.ToUpper(cmd) {
		case AddString:
			c.Add = append(c.Add, args)
			c.All = append(c.All, []string{AddString, args})
		case ArgString:
			c.Arg = append(c.Arg, args)
			c.All = append(c.All, []string{ArgString, args})
		case CmdString:
			c.Cmd = append(c.Cmd, args)
			c.All = append(c.All, []string{CmdString, args})
		case CopyString:
			c.Copy = append(c.Copy, args)
			c.All = append(c.All, []string{CopyString, args})
		case EntrypointString:
			c.Entrypoint = append(c.Entrypoint, args)
			c.All = append(c.All, []string{EntrypointString, args})
		case EnvString:
			c.Env = append(c.Env, args)
			c.All = append(c.All, []string{EnvString, args})
		case ExposeString:
			c.Expose = append(c.Expose, args)
			c.All = append(c.All, []string{ExposeString, args})
		case FromString:
			c.From = append(c.From, args)
			c.All = append(c.All, []string{FromString, args})
		case HealthcheckString:
			c.Healthcheck = append(c.Healthcheck, args)
			c.All = append(c.All, []string{HealthcheckString, args})
		case LabelString:
			c.Label = append(c.Label, args)
			c.All = append(c.All, []string{LabelString, args})
		case MaintainerString:
			c.Maintainer = append(c.Maintainer, args)
			c.All = append(c.All, []string{MaintainerString, args})
		case OnbuildString:
			c.Onbuild = append(c.Onbuild, args)
			c.All = append(c.All, []string{OnbuildString, args})
		case RunString:
			c.Run = append(c.Run, args)
			c.All = append(c.All, []string{RunString, args})
		case ShellString:
			c.Shell = append(c.Shell, args)
			c.All = append(c.All, []string{ShellString, args})
		case StopSignalString:
			c.StopSignal = append(c.StopSignal, args)
			c.All = append(c.All, []string{StopSignalString, args})
		case UserString:
			c.User = append(c.User, args)
			c.All = append(c.All, []string{UserString, args})
		case VolumeString:
			c.Volume = append(c.Volume, args)
			c.All = append(c.All, []string{VolumeString, args})
		case WorkdirString:
			c.Workdir = append(c.Workdir, args)
			c.All = append(c.All, []string{WorkdirString, args})
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
