package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shu-go/gli/v2"
	"github.com/weidewang/go-strftime"
)

// Version is app version
var Version string

type globalCmd struct {
	Line string `cli:"line" default:"[%H:%M:%S] " help:"%Y %m %d %a %H %M %S %Z"`
}

func (c globalCmd) Run(args []string) error {
	if len(args) > 0 {
		c.Line = strings.Join(args, " ")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		t := time.Now()
		fmt.Printf("%v%v\n", strftime.Strftime(&t, c.Line), line)
	}

	return nil
}

func main() {
	app := gli.NewWith(&globalCmd{})
	app.Name = "ts"
	app.Desc = "insert timestamps into stdin lines"
	app.Version = Version
	app.Usage = `more YOURFILE | ts
more YOURFILE | ts %H:%M:%S`
	app.Copyright = "(C) 2020 Shuhei Kubota"
	app.Run(os.Args)
}
