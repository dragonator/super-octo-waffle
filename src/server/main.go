package main

import (
	"net/http"
	"os"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"

	"github.com/dragonator/super-octo-waffle/src/server/handlers"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	authorized := r.Group("/")
	authorized.Use(authRequired())
	authorized.GET("/api/pinnedRepos/:organization", handlers.FetchPinnedItemsHandler)
	authorized.GET("/api/repo/:organization/:repository", handlers.FetchRepositoryDataHandler)
	authorized.GET("/api/commit/:organization/:repository/:sha", handlers.DownloadCommitPatchHandler)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}

// ValidateRequest will verify that a token received from an http request
// is valid and signyed by Auth0
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		audience := os.Getenv("AUTH0_API_IDENTIFIER")
		domain := os.Getenv("AUTH0_DOMAIN")

		var auth0Domain = "https://" + domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(c.Request)

		if err != nil {
			terminateWithError(http.StatusUnauthorized, "token is not valid", c)
			return
		}
		c.Next()
	}
}

func terminateWithError(statusCode int, message string, c *gin.Context) {
	c.JSON(statusCode, gin.H{"error": message})
	c.Abort()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
