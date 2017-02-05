package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	outputDir string
	goTool    string
)

func init() {
	flag.StringVar(&outputDir, "o", "", "output directory")
	flag.StringVar(&goTool, "g", "go", "alternative go tool")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}

	pluginsDir := args[0]

	if outputDir != "" {
		err := os.MkdirAll(outputDir, 0774)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := filepath.Walk(pluginsDir, func(filePath string, file os.FileInfo, err error) error {
		if filepath.Ext(file.Name()) != ".go" {
			return nil
		}
		buildErr := buildPlugin(goTool, filepath.Dir(filePath), outputDir)
		if buildErr != nil {
			return buildErr
		}
		return filepath.SkipDir
	})
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: plugins-build [-g goTool] [-o outputDir] pluginsDir")
	flag.PrintDefaults()
}

func buildPlugin(goTool, pluginDir, outputDir string) error {
	outputFile := filepath.Join(outputDir, filepath.Base(pluginDir)+".so")
	cmd := exec.Command(goTool, "build", "-o", outputFile, "-buildmode=plugin", "./"+pluginDir)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %s", pluginDir, stderr.String())
	}
	return nil
}
