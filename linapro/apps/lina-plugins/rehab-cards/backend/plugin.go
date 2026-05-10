// Package backend wires the rehab-cards source plugin into the LinaPro host.
package backend

import (
	"context"

	"lina-core/pkg/pluginhost"
	rehabcards "lina-plugin-rehab-cards"
	cardcontroller "lina-plugin-rehab-cards/backend/internal/controller/card"
	categorycontroller "lina-plugin-rehab-cards/backend/internal/controller/category"
	crawlercontroller "lina-plugin-rehab-cards/backend/internal/controller/crawler"
	imagecontroller "lina-plugin-rehab-cards/backend/internal/controller/image"
)

// pluginID is the immutable identifier published by the embedded source plugin.
const pluginID = "rehab-cards"

// init registers the rehab-cards source plugin and its HTTP callbacks.
func init() {
	plugin := pluginhost.NewSourcePlugin(pluginID)
	plugin.Assets().UseEmbeddedFiles(rehabcards.EmbeddedFiles)
	plugin.HTTP().RegisterRoutes(
		pluginhost.ExtensionPointHTTPRouteRegister,
		pluginhost.CallbackExecutionModeBlocking,
		registerRoutes,
	)
	pluginhost.RegisterSourcePlugin(plugin)
}

// registerRoutes binds card, category, and crawler routes through host middleware.
func registerRoutes(ctx context.Context, registrar pluginhost.HTTPRegistrar) error {
	routes := registrar.Routes()
	middlewares := routes.Middlewares()
	routes.Group("/api/v1", func(group pluginhost.RouteGroup) {
		group.Middleware(
			middlewares.NeverDoneCtx(),
			middlewares.HandlerResponse(),
			middlewares.CORS(),
			middlewares.RequestBodyLimit(),
			middlewares.Ctx(),
		)
		group.Group("/", func(group pluginhost.RouteGroup) {
			group.Middleware(
				middlewares.Auth(),
				middlewares.Permission(),
			)
			group.Bind(categorycontroller.NewV1())
			group.Bind(cardcontroller.NewV1())
			group.Bind(crawlercontroller.NewV1())
			group.Bind(imagecontroller.NewV1())
		})
	})
	return nil
}
