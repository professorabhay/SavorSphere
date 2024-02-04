package main

import(
	routes "github.com/professorabhay/savorsphere/routes"
	"os"
	"github.com/gin-gonic/gin"
)

func main(){
	port= os.Getenv("PORT")

	if port==""{
		port="8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){ //similar to w r in mux
		c.JSON(200, gin.H{"success":"Access API"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api-2"})
	})

	router.Run(":" + port) 
}