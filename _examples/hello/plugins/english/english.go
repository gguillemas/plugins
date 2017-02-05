package main

import "C"

import (
	"fmt"

	"github.com/gguillemas/plugins"
)

var PluginMetadata = plugins.PluginMetadata{
	Name:        "english",
	Version:     "0.0.1",
	Description: "This plugin says \"hello\" in English.",
}

func Hello() {
	fmt.Println("Hello!")
}
