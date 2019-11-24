package main

import (
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/dragonator/super-octo-waffle/handlers"
)

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	r.GET("/pinnedItems/:organization", handlers.FetchPinnedItemsHandler)
	r.GET("/pinnedItems/:organization/:repository", handlers.FetchRepositoryDataHandler)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
