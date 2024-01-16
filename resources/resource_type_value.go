package resources

import (
	"fmt"
	"strings"
)

type ResourceTypeValue struct {
	Value *ResourceType
}

func (r *ResourceTypeValue) String() string {
	return fmt.Sprintf("%d", *r)
}

func (r *ResourceTypeValue) Set(value string) error {
	switch strings.ToLower(value) {
	case "posts":
		*r.Value = Posts
	case "comments":
		*r.Value = Comments
	case "albums":
		*r.Value = Albums
	case "photos":
		*r.Value = Photos
	case "todos":
		*r.Value = Todos
	case "users":
		*r.Value = Users
	default:
		return fmt.Errorf("invalid resource type: %s", value)
	}
	return nil
}
