// Package backend wires the rehab-cards source plugin into the LinaPro host.
package backend

import (
	"context"

	"lina-core/pkg/logger"
	"lina-core/pkg/pluginhost"
	rehabcards "lina-plugin-rehab-cards"
	cardcontroller "lina-plugin-rehab-cards/backend/internal/controller/card"
	categorycontroller "lina-plugin-rehab-cards/backend/internal/controller/category"
	crawlercontroller "lina-plugin-rehab-cards/backend/internal/controller/crawler"
	seedsvc "lina-plugin-rehab-cards/backend/internal/service/seed"
)

// pluginID is the immutable identifier published by the embedded source plugin.
const pluginID = "rehab-cards"

// rehabCardCatalogSyncName identifies the daily catalog refresh cron.
const rehabCardCatalogSyncName = "rehab-card-catalog-sync"

// rehabCardCatalogSyncDisplayName is the English title for the catalog refresh cron.
const rehabCardCatalogSyncDisplayName = "Rehab Card Catalog Sync"

// rehabCardCatalogSyncDescription is the English description for the catalog refresh cron.
const rehabCardCatalogSyncDescription = "Inserts any newly defined rehab picture cards from the canonical catalog without overwriting manual edits."

// init registers the rehab-cards source plugin and its HTTP, cron, and hook callbacks.
func init() {
	plugin := pluginhost.NewSourcePlugin(pluginID)
	plugin.Assets().UseEmbeddedFiles(rehabcards.EmbeddedFiles)
	plugin.HTTP().RegisterRoutes(
		pluginhost.ExtensionPointHTTPRouteRegister,
		pluginhost.CallbackExecutionModeBlocking,
		registerRoutes,
	)
	plugin.Cron().RegisterCron(
		pluginhost.ExtensionPointCronRegister,
		pluginhost.CallbackExecutionModeBlocking,
		registerCatalogCron,
	)
	plugin.Hooks().RegisterHook(
		pluginhost.ExtensionPointSystemStarted,
		pluginhost.CallbackExecutionModeAsync,
		syncCatalogOnSystemStarted,
	)
	pluginhost.RegisterSourcePlugin(plugin)
}

// registerRoutes binds rehab-cards routes through host middleware. The patient
// H5 reads cards via the public listing endpoint and renders images directly
// from upstream URLs without going through any backend proxy.
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
		group.Bind(cardcontroller.NewPublicV1())
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

// registerCatalogCron schedules the daily rehab-card catalog refresh. The cron
// only runs on the primary cluster node to avoid duplicate insert attempts.
func registerCatalogCron(ctx context.Context, registrar pluginhost.CronRegistrar) error {
	return registrar.AddWithMetadata(
		ctx,
		"0 30 3 * * *",
		rehabCardCatalogSyncName,
		rehabCardCatalogSyncDisplayName,
		rehabCardCatalogSyncDescription,
		func(jobCtx context.Context) error {
			if registrar != nil && !registrar.IsPrimaryNode() {
				return nil
			}
			inserted, err := seedsvc.New().Sync(jobCtx)
			if err != nil {
				return err
			}
			logger.Infof(jobCtx, "rehab-card catalog cron sync finished inserted=%d", inserted)
			return nil
		},
	)
}

// syncCatalogOnSystemStarted runs one eager catalog sync after host startup so
// the patient H5 has playable cards even before the daily cron fires.
func syncCatalogOnSystemStarted(ctx context.Context, payload pluginhost.HookPayload) error {
	inserted, err := seedsvc.New().Sync(ctx)
	if err != nil {
		logger.Warningf(ctx, "rehab-card catalog startup sync failed err=%v", err)
		return err
	}
	logger.Infof(ctx, "rehab-card catalog startup sync finished inserted=%d", inserted)
	return nil
}
