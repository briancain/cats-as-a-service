package static

import "embed"

// Bundle up all ui components into an embedded filesystem, and export it
//go:embed *.css *.jpeg *.map
var StaticsFS embed.FS
