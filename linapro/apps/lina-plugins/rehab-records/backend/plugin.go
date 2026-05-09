// Package backend wires the rehab-records source plugin into the LinaPro host.
package backend

import (
	"context"

	"lina-core/pkg/pluginhost"
	rehabrecords "lina-plugin-rehab-records"
	recordcontroller "lina-plugin-rehab-records/backend/internal/controller/record"
	summarycontroller "lina-plugin-rehab-records/backend/internal/controller/summary"
)

// pluginID is the immutable identifier published by the embedded source plugin.
const pluginID = "rehab-records"

// init registers the rehab-records source plugin and its HTTP callbacks.
func init() {
	plugin := pluginhost.NewSourcePlugin(pluginID)
	plugin.Assets().UseEmbeddedFiles(rehabrecords.EmbeddedFiles)
	plugin.HTTP().RegisterRoutes(pluginhost.ExtensionPointHTTPRouteRegister, pluginhost.CallbackExecutionModeBlocking, registerRoutes)
	pluginhost.RegisterSourcePlugin(plugin)
}

// registerRoutes binds record and summary routes through host middleware.
func registerRoutes(ctx context.Context, registrar pluginhost.HTTPRegistrar) error {
	routes := registrar.Routes()
	middlewares := routes.Middlewares()
	routes.Group("/api/v1", func(group pluginhost.RouteGroup) {
		group.Middleware(middlewares.NeverDoneCtx(), middlewares.HandlerResponse(), middlewares.CORS(), middlewares.RequestBodyLimit(), middlewares.Ctx())
		group.Group("/", func(group pluginhost.RouteGroup) {
			group.Middleware(middlewares.Auth(), middlewares.Permission())
			group.Bind(recordcontroller.NewV1())
			group.Bind(summarycontroller.NewV1())
		})
	})
	return nil
}
