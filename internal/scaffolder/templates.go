package scaffolder

import "embed"

//go:embed templates/docker/*.tmpl
var dockerTemplates embed.FS
