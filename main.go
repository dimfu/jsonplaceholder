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

	// handle single resource
	if firstArgWithDash != 2 {
		rs := resources.GetResources(resourceArg)
		r := resources.PickRandomResource(rs)

		dump.P(r)
		os.Exit(1)
	}

	switch command {
	case "resources": // get all resources
		rs := resources.GetResources(resourceArg)
		dump.P(rs)
	default:
		log.Fatal("no command found for this")
	}
}
