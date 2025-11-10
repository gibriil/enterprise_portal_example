package html

import "embed"

//go:embed header.tmpl footer.tmpl
var WrapperTemplates embed.FS
