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

// registerRoutes binds rehab-cards routes through host middleware. The image
// proxy is exposed publicly so the patient H5 can render <img> tags without
// attaching a JWT, while CRUD and crawler routes still require authentication
// and per-permission authorization.
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
		group.Bind(imagecontroller.NewV1())
		group.Group("/", func(group pluginhost.RouteGroup) {
			group.Middleware(
				middlewares.Auth(),
				middlewares.Permission(),
			)
			group.Bind(categorycontroller.NewV1())
			group.Bind(cardcontroller.NewV1())
			group.Bind(crawlercontroller.NewV1())
		})
	})
	return nil
}
