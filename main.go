// Package main contains API for Canaveral CLI.
// See README.md for more documentation
package main

import (
	gh "canaveral/gh"
	"canaveral/git"
	"canaveral/lib"
	"canaveral/vscodesupport"
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
	var initRepo = false
	var commitMessage = ""

	// Set home directory path of current user
	usr, err := user.Current()
	lib.Check(err)
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
				Aliases:     []string{"c", "create"},
				Description: "Creates a new project, specify name and type.",
				Usage:       "Launch New Project",
				Action: func(c *cli.Context) error {
					projName := c.Args().Get(0)
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					addProjectHandler(projName, projType, initRepo)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "type",
						Aliases:     []string{"t"},
						Usage:       "Specify the type of project you create.",
						Destination: &projType,
					},
					&cli.BoolFlag{
						Name:        "gitinit",
						Aliases:     []string{"g"},
						Usage:       "Initialize a git repo",
						Destination: &initRepo,
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
				Name:    "add github credentials",
				Aliases: []string{"gh", "github", "addgh", "addgithub"},
				Description: `Allows canaveral to use your github credentials for repo management.
					Username and password are required.
					Username and password are stored in native storage for security.`,
				Usage: "Add github info to canaveral.",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return gh.GHAddWrapper()
				},
			},
			{
				Name:        "remove github",
				Aliases:     []string{"rgh", "remgh", "rgithub", "remgithub", "removegithub"},
				Description: `Deletes your github credentials from native storage. Canaveral will no longer have any way to reference your githubcredentials.`,
				Usage:       "Remove github info from canaveral",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return gh.RemGHCredsHandler()
				},
			},
			{
				Name:        "print github",
				Aliases:     []string{"pgh", "pgithub", "printgithub"},
				Description: `Prints the github username currntly stored`,
				Usage:       "Print github info to command line",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					gh.PrintGHUser()
					return nil
				},
			},
			{
				Name:        "git status",
				Aliases:     []string{"status"},
				Description: `Prints current git status in a git directory`,
				Usage:       "Print git status",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					git.Status()
					return nil
				},
			},
			{
				Name:        "git add",
				Aliases:     []string{"add"},
				Description: `Adds all files to next git commit`,
				Usage:       "Add git files. Specify filenames as commandline arguments, use '.' to add all files",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					if c.Args().Len() == 0 {
						fmt.Println("Files to add must be specified. Use '.' for all files")
					}
					git.Add(c.Args().Slice())
					return nil
				},
			},
			{
				Name:        "git commit",
				Aliases:     []string{"commit"},
				Description: `Commits currently added files`,
				Usage:       "Commit changed files",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					git.Commit(commitMessage)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "commit message",
						Aliases:     []string{"m"},
						Usage:       "Add commit message from commandline",
						Destination: &commitMessage,
					},
				},
			},
			{
				Name:        "git ignore",
				Aliases:     []string{"ignore", "ign"},
				Description: `Add specified file to .gitignore`,
				Usage:       "Ignore specified files",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					git.Ignore(c.Args().Slice())
					return nil
				},
			},
			{
				Name:        "code",
				Aliases:     []string{"vscode"},
				Description: "Opens selected project in vscode",
				Usage:       "Opens selected project in vscode",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					projName := c.Args().Get(0)
					vscodesupport.OpenCode(projName, usrHome+confDir+wsFName)
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
