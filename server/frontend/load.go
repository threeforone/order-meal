package frontend

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed static/dist/assets
	staticFS embed.FS

	//go:embed static/dist/index.html
	indexPage []byte

	//go:embed static/dist/vite.svg
	iconFS embed.FS
)

func Init() {

}

func LoadStaticR(g *gin.Engine) {
	// use /assets as root
	staticFiles, error := fs.Sub(staticFS, "static/dist/assets")
	if error != nil {
		panic(error)
	}
	g.StaticFS("/assets/", http.FS(staticFiles))

	//load icon
	g.StaticFileFS("/vite.svg", "static/dist/vite.svg", http.FS(iconFS))

	//index.html
	g.GET("/", func(c *gin.Context) {
		c.Header("content-type", "text/html;charset=utf-8")
		c.String(200, string(indexPage))
	})
}
