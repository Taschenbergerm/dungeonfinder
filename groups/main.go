package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/taschenbergerm/dungeonfinger/groups/docs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"time"
)

const Database = "groups"
const Collection = "groups"

type Groups []Group

type Group struct {
	Id           string   `json:"id", bson:"id"`
	Name         string   `json:"name" bson:"name"`
	Capacity     int      `json:"capacity" bson:"capacity"`
	Participants []string `json:"participants" bson:"participants"`
	Dm           string   `json:"dm" bson:"dm"`
}

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
	v1.GET("/greet/:name", Hello)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":5000")
}

// HelloService godoc
// @Summary Say Hello to a User
// @Schemas
// @Tags greetings
// @Accept json
// @Produce json
// @Success 200 {string} Hello <name>
// @Param name path string true "Name to Greet"
// @Router /greet/{name} [get]
func Hello(g *gin.Context) {
	name := g.Param("name")
	g.JSON(http.StatusOK, fmt.Sprintf("Hello %s", name))
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
	payload := Groups{}
	groups := Groups{
		{"1", "Marvelous Group", 1, []string{"Theo"}, "Marvin"},
		{"2", "Marvelous Group2", 2, []string{"Theo"}, "Marvin"}}
	isOpen := g.Query("open")
	log.Printf("isOpen = %s", isOpen)
	if isOpen == "true" {
		log.Print("Filter down Groups to only open Groups")
		for _, group := range groups {
			if group.Capacity > len(group.Participants) {
				log.Printf("Found %s to be open for new player", group.Name)
				payload = append(payload, group)
			}
		}
	} else {
		log.Print("Do not filter done groups")
		payload = groups
	}
	log.Info().Str("path", "/groups").Int("Length", len(payload))
	g.JSON(200, payload)
}

func queryUsers(uri string) (Groups, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return Groups{}, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return Groups{}, nil
	}
	defer client.Disconnect(ctx)
	db := client.Database(Database)
	collection := db.Collection(Collection)
	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return Groups{}, nil
	}
	var groupsFiltered Groups

	if err = cursor.All(ctx, &groupsFiltered); err != nil {
		return Groups{}, nil
	}
	return groupsFiltered, nil
}
