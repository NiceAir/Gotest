package main

import (
	"fmt"
	"log"
	"os"

	"Gotest/package/github"
)

func main()  {
	var input = os.Args[1:]
	//input = "repo:golang/go is:open json decoder"
	result, err := github.SearchIssues(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items{
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}



