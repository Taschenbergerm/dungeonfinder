package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func queryGroupById(id string) (*Group, error) {
	groups, err := queryGroups(false)

	if err != nil {
		return &Group{}, err
	}
	for _, group := range *groups {
		if group.Id == id {
			return &group, nil
		}
	}
	return &Group{}, nil
}

func queryGroups(isOpen bool) (*Groups, error) {
	payload := Groups{}
	groups := Groups{
		{"1", "Marvelous Group", 1, []string{"Liam"}, "Mathew"},
		{"2", "Marvelous Group2", 2, []string{"Liam"}, "Mathe"}}
	log.Printf("isOpen = %v", isOpen)
	if isOpen {
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
	return &payload, nil
}

func retrieveGroups(uri string) (Groups, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return Groups{}, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return Groups{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := client.Database(Database)
	collection := db.Collection(Collection)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return Groups{}, nil
	}
	var groups Groups

	if err = cursor.All(ctx, &groups); err != nil {
		return Groups{}, err
	}
	return groups, nil
}
