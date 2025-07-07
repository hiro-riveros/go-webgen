package scaffolder

import "embed"

//go:embed templates/docker/*.tmpl
var dockerTemplates embed.FS

//go:embed templates/gin/*.tmpl
//go:embed templates/gin/api/**/*.tmpl
//go:embed templates/gin/config/*.tmpl
var ginTemplates embed.FS
