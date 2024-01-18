package resources

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type ResourceType string

const (
	Posts    ResourceType = "posts"
	Comments ResourceType = "comments"
	Albums   ResourceType = "albums"
	Photos   ResourceType = "photos"
	Todos    ResourceType = "todos"
	Users    ResourceType = "users"
)

var jphUrl = "https://jsonplaceholder.typicode.com"

func GetResources(r ResourceType) []interface{} {
	res, err := http.Get(jphUrl + "/" + string(r))

	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()

	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err.Error())
	}

	var resources []interface{}

	for _, d := range data {
		resources = append(resources, d)
	}

	return resources
}

func PickRandomResource(rs []interface{}) interface{} {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return rs[randomizer.Intn(len(rs))]
}
