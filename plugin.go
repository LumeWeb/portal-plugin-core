// Package plugin provides the main entry point for the core plugin
package plugin

import (
	"go.lumeweb.com/portal-plugin-core/build"
	"go.lumeweb.com/portal-plugin-core/internal"
	"go.lumeweb.com/portal/core"
	portal_plugin_core "go.lumeweb.com/web/go/portal-plugin-core"
)

// init registers the plugin with the Portal framework
// This is called automatically when the plugin is loaded
func init() {
	core.RegisterPlugin(core.PluginInfo{
		ID:         internal.PLUGIN_NAME,
		Version:    build.GetInfo(),
		WebBundles: core.NewWebBundles(core.NewWebBundle(portal_plugin_core.GetFS())),
	})
}
