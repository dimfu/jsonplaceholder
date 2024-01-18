package main

import (
	"flag"
	"jsonplaceholder/resources"
	"log"
	"os"

	"github.com/gookit/goutil/dump"
)

var (
	resourceArg resources.ResourceType
	filterArg   string
)

func init() {
	flag.Var(&resources.ResourceTypeValue{Value: &resourceArg}, "res", "specify resource type")
	flag.StringVar(&filterArg, "f", "", "filter resources properties, can use multiple values")
}

func main() {
	firstArgWithDash := 1

	for i := 0; i < len(os.Args); i++ {
		firstArgWithDash = i

		if len(os.Args[i]) > 0 && os.Args[i][0] == '-' {
			break
		}
	}
	flag.CommandLine.Parse(os.Args[firstArgWithDash:])
	command := os.Args[1]

	if firstArgWithDash != 2 {
		log.Fatal("usage: jsonplaceholder <command>")
	}

	switch command {
	case "resources":
		rs := resources.GetResource(resourceArg)
		filteredResources := resources.FilterResourceProperties(rs)

		dump.P(filteredResources)
	default:
		log.Fatal("no command found for this")
	}
}
