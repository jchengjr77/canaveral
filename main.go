// Package main contains API for Canaveral CLI.
// See README.md for more documentation
package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/jchengjr77/canaveral/finder"
	gh "github.com/jchengjr77/canaveral/gh"
	"github.com/jchengjr77/canaveral/git"
	"github.com/jchengjr77/canaveral/lib"
	"github.com/jchengjr77/canaveral/vscodesupport"

	"github.com/urfave/cli/v2"
)

func main() {

	// Flags
	var qFlag = false
	var projType = "default"
	var initRepo = false
	var commitMessage = ""
	var projPath = ""

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
				Description: "Creates a new project, specify name, type, and initializing a git repo.",
				Usage:       "Launch New Project",
				Action: func(c *cli.Context) error {
					projName := c.Args().Get(0)
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return addProjectHandler(projName, projType, initRepo)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "type",
						Aliases: []string{"t"},
						Usage: `
	Specify the type of project you create. Supported types: 
	react, reactnative, node, python, c
	`,
						Destination: &projType,
					},
					&cli.BoolFlag{
						Name:        "gitinit",
						Aliases:     []string{"g"},
						Usage:       "Initialize a git repo for new project",
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
					return remProjectHandler(projName)
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
					return setWorkspaceHandler(newWorkspace)
				},
			},
			{
				Name:    "addgh",
				Aliases: []string{"agh", "agithub", "addgithub"},
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
				Name:        "remgh",
				Aliases:     []string{"rgh", "rgithub", "remgithub", "removegithub"},
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
				Name:        "printgh",
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
				Name:        "gitstatus",
				Aliases:     []string{"status"},
				Description: `Prints current git status in a git directory`,
				Usage:       "Print git status",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return git.Status(usrHome+confDir+wsFName, projPath)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "project",
						Aliases:     []string{"p"},
						Usage:       "Specify a project to commit",
						Destination: &projPath,
					},
				},
			},
			{
				Name:        "gitadd",
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
					return git.Add(
						c.Args().Slice(), usrHome+confDir+wsFName, projPath)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "project",
						Aliases:     []string{"p"},
						Usage:       "Specify a project to commit",
						Destination: &projPath,
					},
				},
			},
			{
				Name:        "gitcommit",
				Aliases:     []string{"commit"},
				Description: `Commits currently added files`,
				Usage:       "Commit changed files",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return git.Commit(
						commitMessage, usrHome+confDir+wsFName, projPath)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "message",
						Aliases:     []string{"m"},
						Usage:       "Add commit message from commandline",
						Destination: &commitMessage,
					},
					&cli.StringFlag{
						Name:        "project",
						Aliases:     []string{"p"},
						Usage:       "Specify a project to commit",
						Destination: &projPath,
					},
				},
			},
			{
				Name:        "gitignore",
				Aliases:     []string{"ignore", "ign"},
				Description: `Add specified file to .gitignore`,
				Usage:       "Ignore specified files",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return git.Ignore(
						c.Args().Slice(), usrHome+confDir+wsFName, projPath)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "project",
						Aliases:     []string{"p"},
						Usage:       "Specify a project to commit",
						Destination: &projPath,
					},
				},
			},
			{
				Name:        "gitinit",
				Description: `Initialize git repo`,
				Usage:       "Initialize git repo",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					return git.InitRepo(usrHome+confDir+wsFName, projPath)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "project",
						Aliases:     []string{"p"},
						Usage:       "Specify a project to commit",
						Destination: &projPath,
					},
				},
			},
			{
				Name:        "gitremind",
				Aliases:     []string{"remind"},
				Description: `Create a reminder for commits`,
				Usage:       "Get reminded of important details before comitting changes",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					if c.Args().Len() < 2 {
						fmt.Println("Too few arguments for reminder. You must specify a file to add a reminder for and a message (enclosed in quotes) to remind with. Example: canaveral remind test \"test message\"")
						return nil
					}
					if c.Args().Len() >= 3 {
						fmt.Println("Too many arguments for reminder. You must only specify a file to add a reminder for and a message (enclosed in quotes) to remind with. Example: canaveral remind test \"test message\"")
						return nil
					}
					err := git.AddReminder(c.Args().Get(0), c.Args().Get(1))
					return err
				},
			},
			{
				Name:        "printreminders",
				Aliases:     []string{"printrem", "prem"},
				Description: `Create a reminder for commits`,
				Usage:       "Get reminded of important details before comitting changes",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					err := git.CheckReminders(c.Args().First())
					return err
				},
			},
			{
				Name:        "delreminder",
				Aliases:     []string{"deleterem", "delrem", "drem"},
				Description: `delete a commit reminder for a file`,
				Usage:       "Delete a stored reminder",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					if c.Args().Len() < 1 {
						fmt.Println("You must specify a file to delete reminders for and (optionally) a reminder to delete")
						return nil
					}
					if c.Args().Len() == 1 {
						err := git.DelReminder(c.Args().First(), "")
						return err
					}
					err := git.DelReminder(c.Args().Get(0), c.Args().Get(1))
					return err
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
					err := vscodesupport.OpenCode(
						projName, usrHome+confDir+wsFName)
					return err
				},
			},
			{
				Name:    "explore",
				Aliases: []string{"open", "find"},
				Description: `
				Opens selected project in a file explorer window.
				For MacOS users, this will be the Finder.
				Argument should be a project name.`,
				Usage: "Opens selected project in a file explorer window",
				Action: func(c *cli.Context) error {
					if qFlag {
						fmt.Println("(okay, I'll try to be quiet.)")
					}
					projName := c.Args().Get(0)
					err := finder.OpenFinder(
						projName, usrHome+confDir+wsFName)
					return err
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
