package plugins

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-plugin"
	// other imports
)

// PluginConfig holds configuration for a single plugin
type PluginConfig struct {
	Name   string
	Path   string                 // Path to the executable plugin
	Config map[string]interface{} // Additional plugin-specific configuration
}

// HandshakeConfig is used to configure the handshake between this plugin
// and the host.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "EXAMPLE_PLUGIN",
	MagicCookieValue: "hello",
}

// Define a map for plugins if needed
var pluginMap = map[string]plugin.Plugin{
	// "example": &ExamplePlugin{},
	// Add your plugin mappings here
}

// discoverPlugins should search a directory for plugin executables and return their configurations
func discoverPlugins(directory string) ([]PluginConfig, error) {
	var configs []PluginConfig
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Can't access path, ignore it
		}
		if !info.IsDir() {
			configs = append(configs, PluginConfig{
				Name:   info.Name(),
				Path:   path,
				Config: make(map[string]interface{}), // Placeholder for additional config
			})
		}
		return nil
	})
	return configs, err
}

// loadPlugins loads and initializes plugins based on the given configurations
func loadPlugins(pluginConfigs []PluginConfig) []*plugin.Client {
	var clients []*plugin.Client
	for _, config := range pluginConfigs {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: HandshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command(config.Path), // Start the plugin process
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolNetRPC, plugin.ProtocolGRPC,
			},
		})
		clients = append(clients, client)
	}
	return clients
}

// cleanupPlugins terminates all plugin processes and cleans up resources
func cleanupPlugins(clients []*plugin.Client) {
	for _, client := range clients {
		client.Kill()
	}
}

func main() {
	// Define the directory where your plugins are located
	pluginDir := "./plugins" // This is an example path

	// Discover plugins in the specified directory
	pluginConfigs, err := discoverPlugins(pluginDir)
	if err != nil {
		log.Fatalf("Failed to discover plugins: %v", err)
	}

	// Load and initialize the discovered plugins
	clients := loadPlugins(pluginConfigs)
	// Ensure that we clean up all plugins upon program termination
	defer cleanupPlugins(clients)

	// Here, you could interact with the plugins as needed, for example:
	// 1. Establish RPC connections to the plugins
	// 2. Execute plugin-specific functions or methods
	// 3. Collect and process data from plugins

	// Example: Print the names of loaded plugins
	for _, config := range pluginConfigs {
		log.Printf("Loaded plugin: %s\n", config.Name)
	}

	// Add your own logic to use the plugins as required by your application

	// Finally, cleanup all plugin clients before exiting
	// (this will be handled by the defer statement)
}
