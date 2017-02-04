package main

import "C"

import (
	"fmt"
	"log"
	"os"

	"github.com/gguillemas/plugins"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./%s pluginDir [pluginName]", os.Args[0])
		os.Exit(2)
	}

	ps, err := plugins.LoadPlugins(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 3 {
		fmt.Println("Available plugins:")
		fmt.Println("\tNAME\tVERSION\tDESCRIPTION")
		for _, p := range ps {
			fmt.Printf(
				"\t%s\t%s\t%s\n",
				p.Metadata.Name,
				p.Metadata.Version,
				p.Metadata.Description,
			)
		}
		os.Exit(0)
	}

	p := ps.ByName(os.Args[2])

	hello, err := p.Plugin.Lookup("Hello")
	if err != nil {
		log.Fatal(err)
	}

	hello.(func())()
}
