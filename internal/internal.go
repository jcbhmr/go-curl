package internal

import (
	_ "embed"
	"strings"
)

//go:generate go run ./gen.go
//go:embed VERSION
var versionRaw string
var Version = strings.TrimRight(versionRaw, "\r\n")
