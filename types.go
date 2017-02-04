package plugins

import "plugin"

type Plugins []Plugin

type Plugin struct {
	Path     string
	Plugin   *plugin.Plugin
	Metadata PluginMetadata
}

type PluginMetadata struct {
	Name        string
	Version     string
	Description string
}
