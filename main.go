package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Seasons struct {
	Seasons []Season `json:"seasons"`
}

type Season struct {
	Season   string   `json:"season"`
	Episodes []string `json:"episodes"`
}

func main() {
	file, err := os.Open("tvShows/the-office.json")
	if err != nil {
		fmt.Println("error in open file")
	}

	byteValue, _ := io.ReadAll(file)

	var seasons Seasons

	json.Unmarshal(byteValue, &seasons)

	season := rand.Intn(len(seasons.Seasons))
	episode := rand.Intn(len(seasons.Seasons[season].Episodes))

	fmt.Println("season: ", seasons.Seasons[season].Season)
	fmt.Println("episode: ", episode+1, seasons.Seasons[season].Episodes[episode])

	defer file.Close()
}

func PrintAll(seasons Seasons) {
	for i := 0; i < len(seasons.Seasons); i++ {
		fmt.Println("season: ", seasons.Seasons[i].Season)
		for j := 0; j < len(seasons.Seasons[i].Episodes); j++ {
			fmt.Println(" episode ", j, " ", seasons.Seasons[i].Episodes[j])
		}
	}
}
