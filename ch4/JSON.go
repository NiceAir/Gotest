package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movie struct {
	Title string `json:"tit"`
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies=  []Movie{
	{Title:"Caseblance", Year:1942, Color:false,
		Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main()  {
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		os.Exit(0)
	}
//	fmt.Printf("%s\n", data)
	var tit []struct{Tit string}
	if err := json.Unmarshal(data, &tit); err != nil {
		os.Exit(0)
	}
	fmt.Println(tit)
}