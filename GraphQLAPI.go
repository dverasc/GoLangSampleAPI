package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/graphql-go/graphql"
)

type Album struct {
	ID int 
    Artist Artist
    Name string 
	ReleaseDate string 
	Label string 
	Formats string 
	USSales string 
	USChartPeak string 
}

type Artist struct {
    Name string
    Albums []int
}

//Album{ID: "1", Name: "Camp", ReleaseDate: "November 15, 2011", Label: "Universal", Formats: "CD, LP, Digital", USSales: "242,000", USChartPeak: "11"},
//Album{ID: "2", Name: "Because the Internet", ReleaseDate: "December 10, 2013", Label: "Universal", Formats: "CD, LP, Digital", USSales: "796,000", USChartPeak: "7"},
//Album{ID: "3", Name: "Awaken My Love!", ReleaseDate: "December 2, 2016", Label: "Universal", Formats: "CD, LP, Digital", USSales: "1,320,000", USChartPeak: "5"})


func populate() []Album {

    artist := &Artist{Name: "Childish Gambino", Albums: []int{1}}
    album := Album {

        ID: 1,
        Artist: *artist,
        Name: "Camp", 
        ReleaseDate: "November 15, 2011", 
        Label: "Universal", 
        Formats: "CD, LP, Digital", 
        USSales: "242,000", 
        USChartPeak: "11"}


    album2 := Album {

        ID: 2,
        Artist: *artist,
        Name: "Because the Internet", 
        ReleaseDate: "December 10, 2013", 
        Label: "Universal", 
        Formats: "CD, LP, Digital", 
        USSales: "796,000", 
        USChartPeak: "7" }

     album3 := Album {
        ID: 3, 
        Artist: *artist,
        Name: "Awaken My Love!", 
        ReleaseDate: "December 2, 2016", 
        Label: "Universal", 
        Formats: "CD, LP, Digital", 
        USSales: "1,320,000", 
        USChartPeak: "5"}



var albums []Album
albums = append (albums, album)
albums = append (albums, album2)
albums = append (albums, album3)
return albums

}

var artistType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Artist",
        Fields: graphql.Fields{
            "Name": &graphql.Field{
                Type: graphql.String,
            },
            "Albums": &graphql.Field{
                Type: graphql.NewList(graphql.Int),
            },
        },
    },
)


var albumType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Album",
        Fields: graphql.Fields{
            "ID": &graphql.Field{
                Type: graphql.Int,
            },
            "Artist": &graphql.Field{
                Type: artistType,
            },
            "Name": &graphql.Field{
                Type: graphql.String,
            },
            "ReleaseDate": &graphql.Field{
                Type: graphql.String,
            },
            "Label": &graphql.Field{
                Type: graphql.String,
            },
            "Formats": &graphql.Field{
                Type: graphql.String,
            },
            "USSales": &graphql.Field{
                Type: graphql.String,
            },
            "USChartPeak": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

func main() {
   
    albums := populate()
    
     // Schema
     fields := graphql.Fields{
         "album": &graphql.Field{
             Type:        albumType,
             //add a description to the field
             Description: "Get Tutorial By ID",
            
             //defining argument to choose specific album by ID
             Args: graphql.FieldConfigArgument{
                 "id": &graphql.ArgumentConfig{
                     Type: graphql.Int,
                 },
             },
             Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                 id, ok := p.Args["id"].(int)
                 if ok {
                     // Find album by parsing the array
                     for _, album := range albums {
                         if int(album.ID) == id {
                             return album, nil
                         }
                     }
                 }
                 return nil, nil
             },
         },
         "list": &graphql.Field{
             Type:        graphql.NewList(albumType),
             Description: "Get Album List",
             Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                 return albums, nil
             },
         },
     }
     rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
     schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
     schema, err := graphql.NewSchema(schemaConfig)
     if err != nil {
         log.Fatalf("failed to create new schema, error: %v", err)
     }
 
     // Query
     query := `
         {
             album(id:2) {
                 Name 
                 Artist {
                     Name
                     Albums
                 }
             }
         }
     `
     params := graphql.Params{Schema: schema, RequestString: query}
     r := graphql.Do(params)
     if len(r.Errors) > 0 {
         log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
     }
     rJSON, _ := json.Marshal(r)
     fmt.Printf("%s \n", rJSON)
 }
