// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC client CLI support package
//
// Command:
// $ goa gen assignment_crossnokaye/cod/itemservice/design

package cli

import (
	itemservicec "assignment_crossnokaye/cod/itemservice/gen/grpc/item_service/client"
	"flag"
	"fmt"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `item-service (create|get|update|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` item-service create --message '{
      "damage": 3205420170586316006,
      "description": "Et in alias est expedita quos eligendi.",
      "healing": 1329053061876795014,
      "name": "Est velit culpa aspernatur dolorem.",
      "protection": 120370757675553720
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		itemServiceFlags = flag.NewFlagSet("item-service", flag.ContinueOnError)

		itemServiceCreateFlags       = flag.NewFlagSet("create", flag.ExitOnError)
		itemServiceCreateMessageFlag = itemServiceCreateFlags.String("message", "", "")

		itemServiceGetFlags       = flag.NewFlagSet("get", flag.ExitOnError)
		itemServiceGetMessageFlag = itemServiceGetFlags.String("message", "", "")

		itemServiceUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		itemServiceUpdateMessageFlag = itemServiceUpdateFlags.String("message", "", "")

		itemServiceDeleteFlags       = flag.NewFlagSet("delete", flag.ExitOnError)
		itemServiceDeleteMessageFlag = itemServiceDeleteFlags.String("message", "", "")
	)
	itemServiceFlags.Usage = itemServiceUsage
	itemServiceCreateFlags.Usage = itemServiceCreateUsage
	itemServiceGetFlags.Usage = itemServiceGetUsage
	itemServiceUpdateFlags.Usage = itemServiceUpdateUsage
	itemServiceDeleteFlags.Usage = itemServiceDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "item-service":
			svcf = itemServiceFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "item-service":
			switch epn {
			case "create":
				epf = itemServiceCreateFlags

			case "get":
				epf = itemServiceGetFlags

			case "update":
				epf = itemServiceUpdateFlags

			case "delete":
				epf = itemServiceDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "item-service":
			c := itemservicec.NewClient(cc, opts...)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = itemservicec.BuildCreatePayload(*itemServiceCreateMessageFlag)
			case "get":
				endpoint = c.Get()
				data, err = itemservicec.BuildGetPayload(*itemServiceGetMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = itemservicec.BuildUpdatePayload(*itemServiceUpdateMessageFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = itemservicec.BuildDeletePayload(*itemServiceDeleteMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// item-serviceUsage displays the usage of the item-service command and its
// subcommands.
func itemServiceUsage() {
	fmt.Fprintf(os.Stderr, `The ItemService service handles CRUD operations for items
Usage:
    %[1]s [globalflags] item-service COMMAND [flags]

COMMAND:
    create: Create a new item
    get: Retrieve an item by ID
    update: Update an item
    delete: Delete an item by ID

Additional help:
    %[1]s item-service COMMAND --help
`, os.Args[0])
}
func itemServiceCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service create -message JSON

Create a new item
    -message JSON: 

Example:
    %[1]s item-service create --message '{
      "damage": 3205420170586316006,
      "description": "Et in alias est expedita quos eligendi.",
      "healing": 1329053061876795014,
      "name": "Est velit culpa aspernatur dolorem.",
      "protection": 120370757675553720
   }'
`, os.Args[0])
}

func itemServiceGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service get -message JSON

Retrieve an item by ID
    -message JSON: 

Example:
    %[1]s item-service get --message '{
      "id": 2011463225270065069
   }'
`, os.Args[0])
}

func itemServiceUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service update -message JSON

Update an item
    -message JSON: 

Example:
    %[1]s item-service update --message '{
      "damage": 7849875426662081395,
      "description": "Aut facilis alias eos.",
      "healing": 1175548812336538195,
      "id": 1771631765212308878,
      "name": "Commodi facilis maiores.",
      "protection": 1137017544043901924
   }'
`, os.Args[0])
}

func itemServiceDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service delete -message JSON

Delete an item by ID
    -message JSON: 

Example:
    %[1]s item-service delete --message '{
      "id": 6263443980830511219
   }'
`, os.Args[0])
}
