package static

//go:generate go-bindata -pkg static -o bindata.go files/...
//go:generate go fmt bindata.go
//go:generate sed -i.bak "s/Css/CSS/" bindata.go
//go:generate rm bindata.go.bak
//go:generate sed -i.bak "s/Html/HTML/" bindata.go
//go:generate rm bindata.go.bak

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

// FileSystem is an HTTP filesystem handle for static files.
func FileSystem() http.FileSystem {
	fs := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "files"}
	return &binaryFileSystem{
		fs,
	}
}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name[1:])
}
