package realize

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// The Project struct defines the informations about a project
type Project struct {
	reload  time.Time
	base    string
	Name    string  `yaml:"app_name,omitempty"`
	Path    string  `yaml:"app_path,omitempty"`
	Run     bool    `yaml:"app_run,omitempty"`
	Bin     bool    `yaml:"app_bin,omitempty"`
	Build   bool    `yaml:"app_build,omitempty"`
	Fmt     bool    `yaml:"app_fmt,omitempty"`
	Watcher Watcher `yaml:"app_watcher,omitempty"`
}

// GoRun  is an implementation of the bin execution
func (p *Project) GoRun(channel chan bool, runner chan bool, wr *sync.WaitGroup) error {
	name := strings.Split(p.Path, "/")
	stop := make(chan bool, 1)
	var run string

	if name[len(name)-1] == "" {
		name := strings.Split(slash(p.base), "/")
		run = name[len(name)-1]
	} else {
		run = name[len(name)-1]
	}

	build := exec.Command(slash(os.Getenv("GOBIN")) + slash(run))
	build.Dir = p.base
	defer func() {
		if err := build.Process.Kill(); err != nil {
			log.Fatal(Red("failed to stop: "), Red(err))
		}
		log.Println(pname(p.Name, 2), ":", RedS("Stopped"))
		wr.Done()
	}()

	stdout, err := build.StdoutPipe()
	if err != nil {
		log.Println(Red(err.Error()))
		return err
	}
	if err := build.Start(); err != nil {
		log.Println(Red(err.Error()))
		return err
	}
	close(runner)

	in := bufio.NewScanner(stdout)
	go func() {
		for in.Scan() {
			select {
			default:
				log.Println(pname(p.Name, 3), ":", BlueS(in.Text()))
			}
		}
		close(stop)
	}()

	for {
		select {
		case <-channel:
			return nil
		case <-stop:
			return nil
		}
	}
}

// GoBuild is an implementation of the "go build"
func (p *Project) GoBuild() error {
	var out bytes.Buffer

	build := exec.Command("go", "build")
	build.Dir = p.base
	build.Stdout = &out
	if err := build.Run(); err != nil {
		return err
	}
	return nil
}

// GoInstall is an implementation of the "go install"
func (p *Project) GoInstall() error {
	var out bytes.Buffer
	base, _ := os.Getwd()
	path := base + p.Path

	err := os.Setenv("GOBIN", slash(os.Getenv("GOPATH"))+slash("bin"))
	if err != nil {
		return err
	}

	build := exec.Command("go", "install")
	build.Dir = path
	build.Stdout = &out
	if err := build.Run(); err != nil {
		return err
	}
	return nil
}

// GoFmt is an implementation of the gofmt
func (p *Project) GoFmt() (io.Writer, error) {
	var out bytes.Buffer
	build := exec.Command("gofmt", "-s", "-w", "-e", ".")
	build.Dir = p.base
	build.Stdout = &out
	build.Stderr = &out
	if err := build.Run(); err != nil {
		return build.Stderr, err
	}
	return nil, nil
}
