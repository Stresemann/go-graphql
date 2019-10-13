package main

import "github.com/graphql-go/graphql"

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

var songType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Song",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"album": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"duration": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var songArgs = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
	},
	"album": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
	},
	"title": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
	},
	"duration": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
	},
}
