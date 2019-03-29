package main

import (
	"fmt"
	"os"

	"github.com/adrianosela/godaddy"
	"github.com/urfave/cli"
)

var dryRun = false

func main() {
	app := cli.NewApp()
	app.Name = "dns"
	app.EnableBashCompletion = true
	app.Usage = "CLI tool for managing DNS in GoDaddy"

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Println("[ERROR] The command provided is not supported: ", command)
		c.App.Run([]string{"help"})
	}

	app.Version = "1.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "dry-run",
			Usage:       "run the command in dry-run mode",
			Destination: &dryRun,
		},
	}

	record := cli.Command{
		Name:  "record",
		Usage: "dns record operations",
		Subcommands: []cli.Command{
			// {
			// 	Name:   "create",
			// 	Usage:  "creates a new DNS record",
			// 	Action: createRecord,
			// 	Flags: []cli.Flag{
			// 		cli.StringFlag{
			// 			Name:  "type",
			// 			Usage: "record type",
			// 			Value: "",
			// 		},
			// 		cli.StringFlag{
			// 			Name:  "name",
			// 			Usage: "Name of the record: <name>.yourdomain.com",
			// 			Value: "",
			// 		},
			// 		cli.StringFlag{
			// 			Name:  "points-to",
			// 			Usage: "IP or Hostname for the record to point to",
			// 			Value: "",
			// 		},
			// 	},
			// },
			{
				Name:   "get",
				Usage:  "get all records for a given domain",
				Action: getRecords,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "domain",
						Usage: "owned domain for which to retrieve records",
						Value: "",
					},
					cli.StringFlag{
						Name:  "type",
						Usage: "record type",
						Value: "",
					},
					cli.StringFlag{
						Name:  "name",
						Usage: "name of the record: <name>.yourdomain.com",
						Value: "",
					},
				},
			},
		},
	}

	app.Commands = []cli.Command{
		record,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("\nError: %s \n", err)
		os.Exit(1)
	}
}

// GetRecords gets all records for a given domain
func getRecords(c *cli.Context) error {
	gd := godaddy.NewClient(os.Getenv("GD_API_KEY"), os.Getenv("GD_API_SECRET"), godaddy.HostProd)
	return gd.GetRecords(c.String("domain"), c.String("type"), c.String("name"))
}
