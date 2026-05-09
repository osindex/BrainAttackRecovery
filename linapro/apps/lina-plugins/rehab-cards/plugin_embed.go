// This file embeds the rehab-cards source plugin assets.

package rehabcards

import "embed"

// EmbeddedFiles contains the plugin manifest, frontend pages, and lifecycle resources.
//
//go:embed plugin.yaml frontend manifest
var EmbeddedFiles embed.FS
