// Package main contains API for Canaveral CLI.
// See README.md for more documentation
package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/urfave/cli/v2"
)

func main() {

	// quiet flag
	var qFlag = false

	// Set home directory path of current user
	usr, err := user.Current()
	check(err)
	usrHome = usr.HomeDir

	app := &cli.App{
		Name:  "canaveral",
		Usage: "Launch your new project effortlessly.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "quiet",
				Aliases:     []string{"q"},
				Usage:       "Quiet Mode. Silences all output when active",
				Destination: &qFlag,
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "launch",
				Aliases:     []string{"c", "add", "create"},
				Description: "Creates a new project with name of your choice.",
				Usage:       "Launch New Project",
				Action: func(c *cli.Context) error {
					projName := c.Args().Get(0)
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					addProjectHandler(projName)
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r", "del", "rem", "delete"},
				Description: `Deletes target project from workspace.
				You must provide the name of the project you want to delete.`,
				Usage: "Delete Existing Project",
				Action: func(c *cli.Context) error {
					projName := c.Args().Get(0)
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					remProjectHandler(projName)
					return nil
				},
			},
			{
				Name:    "space",
				Aliases: []string{"path", "setpath"},
				Description: `Sets path to your personal canaveral workspace.
					This path should be one that you can remember.
					It will become the home for all your projects.`,
				Usage: "Set canaveral workspace path.",
				Action: func(c *cli.Context) error {
					newWorkspace := c.Args().Get(0)
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					setWorkspaceHandler(newWorkspace)
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			if qFlag {
				fmt.Println("(okay, I'll try to be quiet.)")
			}
			showWorkspaceHandler()
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
