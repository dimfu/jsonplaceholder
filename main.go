package main

import (
	"flag"
	"jsonplaceholder/resources"
	"log"

	"github.com/gookit/goutil/dump"
)

var resourceArg resources.ResourceType

func init() {
	flag.Var(&resources.ResourceTypeValue{Value: &resourceArg}, "res", "specify resource type")
}

func main() {
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			log.Fatal(f.Name, " flag is required")
		}
	})

	if !resources.IsValidResource(resourceArg) {
		log.Fatal("invalid resource type: ", resourceArg)
	}

	resource := resources.GetResource(resourceArg)

	dump.P(resource)
}
