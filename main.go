package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

func main() {

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"songs": &graphql.Field{
				Type: graphql.NewList(songType),
				Args: songArgs,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					s := Song{}
					b, err := json.Marshal(params.Args)
					if err != nil {
						panic(err)
					}
					err = json.Unmarshal(b, &s)
					filtered := Filter(songs, func(v Song) bool {
						return strings.EqualFold(v.Album, s.Album) || strings.EqualFold(v.Title, s.Title)
					})
					return filtered, nil
				},
			},
			"album": &graphql.Field{
				Type: albumType,
				Args: albumArgs,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(string)
					for _, album := range albums {
						if album.ID == id {
							return album, nil
						}
					}
					return nil, nil
				},
			},
		},
	})
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{"createSong": &graphql.Field{
			Type: songType,
			Args: songArgs,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var song Song
				song.ID = params.Args["id"].(string)
				song.Album = params.Args["album"].(string)
				song.Title = params.Args["title"].(string)
				song.Duration = params.Args["duration"].(string)
				songs = append(songs, song)
				return song, nil
			},
		},
		},
	})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":12345", nil)
}

func Filter(songs []Song, f func(Song) bool) []Song {
	vsf := make([]Song, 0)
	for _, v := range songs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
