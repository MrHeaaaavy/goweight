package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/jondot/goweight/pkg"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	jsonOutput      = flag.Bool("json", false, "Output json")
	jsonOutputShort = flag.Bool("j", false, "Output json")

	buildTags = flag.String("tags", "", "Build tags")
	packages  = flag.String("packages", "", "Packages to build")

	showVersion = flag.Bool("version", false, "Shows version")

	existsWork = flag.String("work", "", "Exists work path")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("%s (%s)\n", version, commit)
		return
	}

	weight := pkg.NewGoWeight()
	if *buildTags != "" {
		weight.BuildCmd = append(weight.BuildCmd, "-tags", *buildTags)
	}
	if *packages != "" {
		weight.BuildCmd = append(weight.BuildCmd, *packages)
	}

	var work string
	if len(*existsWork) > 0 {
		work = *existsWork
	} else {
		work = weight.BuildCurrent()
	}

	log.Printf("Working on path %s\n", work)

	modules := weight.Process(work)

	if *jsonOutput || *jsonOutputShort {
		m, _ := json.Marshal(modules)
		fmt.Print(string(m))
	} else {
		for _, module := range modules {
			fmt.Printf("%8s %s\n", module.SizeHuman, module.Name)
		}
	}
}
