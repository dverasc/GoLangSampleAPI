//RESTFUL API to show childish gambino album catalogue in JSON

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//if I had a MongoDB for this
// var client*mongo.Client

//Album is model for albums
//ReleaseDate, Label, Formats, USSales, USChartPeak
type Album struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releasedate"`
	Label       string `json:"label"`
	Formats     string `json:"formats"`
	USSales     string `json:"ussales"`
	USChartPeak string `json:"uschartpeak"`
}

//initiate discinfo var as a slice

var albums []Album

//Index

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

//Show

//if i had MongoDB, then this would be endpoint to getAlbum
//response.Header().Set("content-type", "application/json")
//id, _ := primitive.ObjectIDFromHex(params["id"])
//var album Album
//collection := client.Database("ChildishGambino").Collection("Albums")
//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
//err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
//if err != nil {
//response.WriteHeader(http.StatusInternalServerError)
//response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
//return
//}
//json.NewEncoder(response).Encode(person)
//}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//Create

//if I had MongoDB instead of mock data
//	_ = json.NewDecoder(request.Body).Decode(&person)
//	collection := client.Database("ChildishGambino").Collection("Albums")
//	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//	result, _ := collection.InsertOne(ctx, person)
//	json.NewEncoder(response).Encode(result)

func createAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAlbum Album
	json.NewDecoder(r.Body).Decode(&newAlbum)
	newAlbum.ID = strconv.Itoa(len(albums) + 1)
	albums = append(albums, newAlbum)
	json.NewEncoder(w).Encode(newAlbum)
}

//Update

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	params := mux.Vars(r)
	for i, item := range albums {
		if item.ID == params["id"] {
			albums = append(albums[:i], albums[i+1:]...)
			var newAlbum Album
			json.NewDecoder(r.Body).Decode(&newAlbum)
			newAlbum.ID = params["id"]
			albums = append(albums, newAlbum)
			json.NewEncoder(w).Encode(newAlbum)
			return
		}
	}
}

//Delete

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range albums {
		if item.ID == params["id"] {
			albums = append(albums[:i], albums[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(albums)
}

func main() {

	//if instead of mock data, I had MongoDB for this
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:8080")
	//client, _ mongo.Connect(ctx, clientOptions)

	//firstalbum of mock data
	albums = append(albums,
		Album{ID: "1", Name: "Camp", ReleaseDate: "November 15, 2011", Label: "Universal", Formats: "CD, LP, Digital", USSales: "242,000", USChartPeak: "11"},
		Album{ID: "2", Name: "Because the Internet", ReleaseDate: "December 10, 2013", Label: "Universal", Formats: "CD, LP, Digital", USSales: "796,000", USChartPeak: "7"},
		Album{ID: "3", Name: "Awaken My Love!", ReleaseDate: "December 2, 2016", Label: "Universal", Formats: "CD, LP, Digital", USSales: "1,320,000", USChartPeak: "5"})

	//initialize router
	router := mux.NewRouter()

	//endpoints for CRUD blah blah
	router.HandleFunc("/album", getAlbum).Methods("GET")
	router.HandleFunc("/album/{id}", getAlbum).Methods("GET")
	router.HandleFunc("/album", createAlbum).Methods("POST")
	router.HandleFunc("/album/{id}", updateAlbum).Methods("POST")
	router.HandleFunc("/album/{id}", deleteAlbum).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
	//http.ListenAndServe(":8080", router)

}
