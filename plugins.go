package plugins

import (
	"io/ioutil"
	"os"
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

func LoadPluginsWalk(pluginsRoot string) (Plugins, error) {
	var ps Plugins
	err := filepath.Walk(pluginsRoot, func(filePath string, file os.FileInfo, err error) error {
		if filepath.Ext(file.Name()) != ".so" {
			return nil
		}
		p, err := LoadPlugin(filePath)
		if err != nil {
			return err
		}
		ps = append(ps, p)
		return nil
	})
	if err != nil {
		return Plugins{}, err
	}
	return ps, nil
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
		p, err := LoadPlugin(path.Join(pluginsPath, file.Name()))
		if err != nil {
			return Plugins{}, err
		}
		ps = append(ps, p)
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
