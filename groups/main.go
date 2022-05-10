package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/taschenbergerm/dungeonfinger/groups/docs"
	"os"
)

// @BasePath /api/v1
// @version alpha
// @title GroupManagementService
// @description API to the Group Management Service
func main() {
	os.Getenv("MONGO_URI")
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	v1.GET("/groups", getGroups)
	v1.GET("/groups/:id", getGroupById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":5000")
}

// getGroups godoc
// @summary Get all available Groups
// @schemas
// @tags groups
// @accept json
// @produces json
// @success 200 {struct} Groups
// @param open  query bool false "false"
// @Router /groups [get]
func getGroups(g *gin.Context) {
	isOpen := g.Query("open") == "true"
	payload, err := queryGroups(isOpen)

	if err != nil {
		g.JSON(500, gin.H{"error": err})
	}
	g.JSON(200, payload)
}

// getGroupById godoc
// @summary Get Group by Id
// @schemas
// @tags groups
// @accept json
// @produces json
// @success 200 {struct} Group
// @param id  query string true
// @Router /groups [get]
func getGroupById(g *gin.Context) {
	groupId := g.Param("id")
	payload, err := queryGroupById(groupId)
	if err != nil {
		g.JSON(500, gin.H{"error": err})
	}
	g.JSON(200, payload)
}
