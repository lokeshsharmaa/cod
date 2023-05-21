// Code generated by goa v3.11.3, DO NOT EDIT.
//
// externalservice HTTP client CLI support package
//
// Command:
// $ goa gen assignment_crossnokaye/cod/externalservice/design

package cli

import (
	externalservicec "assignment_crossnokaye/cod/externalservice/gen/http/externalservice/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `externalservice (create-character|get-character|update-character|delete-character)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` externalservice create-character --body '{
      "description": "Ad excepturi officia necessitatibus autem vero.",
      "experience": 6304895944173359943,
      "health": 7361944416981608895,
      "name": "Ipsa eum necessitatibus ratione commodi."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		externalserviceFlags = flag.NewFlagSet("externalservice", flag.ContinueOnError)

		externalserviceCreateCharacterFlags    = flag.NewFlagSet("create-character", flag.ExitOnError)
		externalserviceCreateCharacterBodyFlag = externalserviceCreateCharacterFlags.String("body", "REQUIRED", "")

		externalserviceGetCharacterFlags  = flag.NewFlagSet("get-character", flag.ExitOnError)
		externalserviceGetCharacterIDFlag = externalserviceGetCharacterFlags.String("id", "REQUIRED", "Character ID")

		externalserviceUpdateCharacterFlags    = flag.NewFlagSet("update-character", flag.ExitOnError)
		externalserviceUpdateCharacterBodyFlag = externalserviceUpdateCharacterFlags.String("body", "REQUIRED", "")
		externalserviceUpdateCharacterIDFlag   = externalserviceUpdateCharacterFlags.String("id", "REQUIRED", "Character ID")

		externalserviceDeleteCharacterFlags  = flag.NewFlagSet("delete-character", flag.ExitOnError)
		externalserviceDeleteCharacterIDFlag = externalserviceDeleteCharacterFlags.String("id", "REQUIRED", "Character ID")
	)
	externalserviceFlags.Usage = externalserviceUsage
	externalserviceCreateCharacterFlags.Usage = externalserviceCreateCharacterUsage
	externalserviceGetCharacterFlags.Usage = externalserviceGetCharacterUsage
	externalserviceUpdateCharacterFlags.Usage = externalserviceUpdateCharacterUsage
	externalserviceDeleteCharacterFlags.Usage = externalserviceDeleteCharacterUsage

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
		case "externalservice":
			svcf = externalserviceFlags
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
		case "externalservice":
			switch epn {
			case "create-character":
				epf = externalserviceCreateCharacterFlags

			case "get-character":
				epf = externalserviceGetCharacterFlags

			case "update-character":
				epf = externalserviceUpdateCharacterFlags

			case "delete-character":
				epf = externalserviceDeleteCharacterFlags

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
		case "externalservice":
			c := externalservicec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create-character":
				endpoint = c.CreateCharacter()
				data, err = externalservicec.BuildCreateCharacterPayload(*externalserviceCreateCharacterBodyFlag)
			case "get-character":
				endpoint = c.GetCharacter()
				data, err = externalservicec.BuildGetCharacterPayload(*externalserviceGetCharacterIDFlag)
			case "update-character":
				endpoint = c.UpdateCharacter()
				data, err = externalservicec.BuildUpdateCharacterPayload(*externalserviceUpdateCharacterBodyFlag, *externalserviceUpdateCharacterIDFlag)
			case "delete-character":
				endpoint = c.DeleteCharacter()
				data, err = externalservicec.BuildDeleteCharacterPayload(*externalserviceDeleteCharacterIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// externalserviceUsage displays the usage of the externalservice command and
// its subcommands.
func externalserviceUsage() {
	fmt.Fprintf(os.Stderr, `Contains API CRUD operations on characters, inventories, and items.
Usage:
    %[1]s [globalflags] externalservice COMMAND [flags]

COMMAND:
    create-character: Creating a new Character
    get-character: Fetching a Character
    update-character: Updating a Character
    delete-character: Deleting a Character

Additional help:
    %[1]s externalservice COMMAND --help
`, os.Args[0])
}
func externalserviceCreateCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] externalservice create-character -body JSON

Creating a new Character
    -body JSON: 

Example:
    %[1]s externalservice create-character --body '{
      "description": "Ad excepturi officia necessitatibus autem vero.",
      "experience": 6304895944173359943,
      "health": 7361944416981608895,
      "name": "Ipsa eum necessitatibus ratione commodi."
   }'
`, os.Args[0])
}

func externalserviceGetCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] externalservice get-character -id INT

Fetching a Character
    -id INT: Character ID

Example:
    %[1]s externalservice get-character --id 8516222454781191823
`, os.Args[0])
}

func externalserviceUpdateCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] externalservice update-character -body JSON -id INT

Updating a Character
    -body JSON: 
    -id INT: Character ID

Example:
    %[1]s externalservice update-character --body '{
      "description": "Quia quae distinctio ut.",
      "experience": 8109871214411202544,
      "health": 5826639684918267823,
      "name": "Eius dolor at molestiae iste qui."
   }' --id 2446497314123511081
`, os.Args[0])
}

func externalserviceDeleteCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] externalservice delete-character -id INT

Deleting a Character
    -id INT: Character ID

Example:
    %[1]s externalservice delete-character --id 6052710139295842249
`, os.Args[0])
}
