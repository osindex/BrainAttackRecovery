// This file embeds the rehab-records source plugin assets.

package rehabrecords

import "embed"

// EmbeddedFiles contains the plugin manifest, frontend pages, and lifecycle resources.
//
//go:embed plugin.yaml frontend manifest
var EmbeddedFiles embed.FS
