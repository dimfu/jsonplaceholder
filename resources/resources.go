package resources

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

var (
	jphUrl         = "https://jsonplaceholder.typicode.com"
	validResources = map[ResourceType]struct{}{
		Posts:    {},
		Comments: {},
		Albums:   {},
		Photos:   {},
		Todos:    {},
		Users:    {},
	}
)

func IsValidResource(r ResourceType) bool {
	_, valid := validResources[r]
	return valid
}

func GetResource(r ResourceType) interface{} {
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
