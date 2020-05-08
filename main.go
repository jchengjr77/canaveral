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

	// Flags
	var qFlag = false
	var projType = "default"

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
					addProjectHandler(projName, projType)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "type",
						Aliases:     []string{"t"},
						Usage:       "Specify the type of project you create.",
						Destination: &projType,
					},
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
			{
				Name:    "add git",
				Aliases: []string{"git", "addgit"},
				Description: `Allows canaveral to use your git credentials for repo management.
					Username and password are required.
					Username and password are stored in native storage for security.`,
				Usage: "Add git info to canaveral.",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return gitAddWrapper()
				},
			},
			{
				Name:    "remove git",
				Aliases: []string{"rgit", "remgit", "removegit"},
				Description: `Deletes your git credentials from native storage.
					Canaveral will no longer have any way to reference your git credentials.`,
				Usage: "Remove git info from canaveral",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return remGitCredsHandler()
				},
			},
			{
				Name:        "print git",
				Aliases:     []string{"pgit", "printgit"},
				Description: `Prints the git username currntly stored`,
				Usage:       "Print git info to command line",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					printGitUser()
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
