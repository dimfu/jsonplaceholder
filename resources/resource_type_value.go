package resources

import (
	"fmt"
	"strings"
)

type ResourceTypeValue struct {
	Value *ResourceType
}

func (r *ResourceTypeValue) String() string {
	if r.Value == nil {
		return ""
	}
	return string(*r.Value)
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
		return fmt.Errorf("invalid resource types: %s", value)
	}

	if string(*r.Value) != value {
		return fmt.Errorf("provided resource type (%s) does not match expected resource type", value)
	}

	return nil
}
