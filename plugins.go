package plugins

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"plugin"
)

func (ps Plugins) ByName(name string) Plugin {
	for _, p := range ps {
		if p.Metadata.Name == name {
			return p
		}
	}
	return Plugin{}
}

func LoadPlugins(pluginsPath string) (Plugins, error) {
	var ps Plugins
	pluginFiles, err := ioutil.ReadDir(pluginsPath)
	if err != nil {
		return Plugins{}, err
	}
	for _, file := range pluginFiles {
		if filepath.Ext(file.Name()) != ".so" {
			continue
		}
		data, err := LoadPlugin(path.Join(pluginsPath, file.Name()))
		if err != nil {
			return Plugins{}, err
		}
		ps = append(ps, data)
	}
	return ps, nil
}

func LoadPlugin(pluginPath string) (Plugin, error) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return Plugin{}, err
	}
	pm, err := p.Lookup("PluginMetadata")
	if err != nil {
		return Plugin{}, err
	}
	metadata := *pm.(*PluginMetadata)
	return Plugin{
		Path:     pluginPath,
		Plugin:   p,
		Metadata: metadata,
	}, nil
}
