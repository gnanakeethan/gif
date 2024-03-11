package main

import (
	"github.com/hashicorp/go-plugin"
)

// Path: plugin_spec.go
type BoaviztaPlugin struct {
	Method       string         `json:"method"`
	Path         string         `json:"path"`
	GlobalConfig map[string]any `json:"global-config"`
}

func (p *BoaviztaPlugin) Init() error {
	return nil
}
func (p *BoaviztaPlugin) Execute() error {
	return nil
}

// HandshakeConfig is used to configure the handshake between this plugin
// and the host. Importantly, the ProtocolVersion and the MagicCookieKey
// and Value should match what the host expects, or the host will reject the
// plugin connection.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"example": &ExamplePlugin{},
}

type ExamplePlugin struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written below
	Impl PluginInterface
}

func main() {
	// We're a host. Start by launching the plugin process.
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"example": &ExamplePlugin{Impl: &BoaviztaPlugin{}},
		},
	})
}
