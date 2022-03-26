package assets

import (
	"embed"
	"io/fs"
)

var (
	FileSystem fs.FS
)

//go:embed static/*
var content embed.FS

func init() {
	Register(content)
}

func Register(content fs.FS) {
	subFs, err := fs.Sub(content, "static")
	if err == nil {
		FileSystem = subFs
	}
}
