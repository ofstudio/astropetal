package embeded

import (
	"embed"
	"io/fs"
)

//go:embed "client-cert"
var clientCertFS embed.FS
var ClientCertFS, _ = fs.Sub(clientCertFS, "client-cert")
