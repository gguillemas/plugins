package main

import "C"

import (
	"fmt"

	"github.com/gguillemas/plugins"
)

var PluginMetadata = plugins.PluginMetadata{
	Name:        "spanish",
	Version:     "0.0.1",
	Description: "This plugin says \"hello\" in Spanish.",
}

func Hello() {
	fmt.Println("Hola!")
}
