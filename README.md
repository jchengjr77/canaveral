# Canaveral

![Go](https://github.com/jchengjr77/canaveral/workflows/Go/badge.svg?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Contents

-   [Introduction](#introduction)
-   [The Name](#the-name)
-   [Getting Started](#getting-started)
    -   [Dependencies](#dependencies)
    -   [Installation](#installation)
-   [Basic usage](#basic-usage)
-   [Project Types](#project-types)
-   [Troubleshooting](#troubleshooting)
-   [Contributing](#contributing)
-   [License](#license)
-   [Credits](#credits)

## Introduction

Launch your new projects seamlessly. Canaveral automates all your project setup away.

**We wanted to have a project manager that was fast, intuitive, and actually useful. So we made Canaveral.**

Canaveral is a Command Line Interface (CLI) tool that can add, remove, and view your projects. It is a tool built by developers for developers. Don't worry about where your projects are located, how they are organized, what their names are, or how to set them up. Canaveral does that for you.

## The Name

Cape Canaveral Air Force Station is one of two main launch sites for U.S. space missions. It is used primarily for launching spacecraft into equatorial orbits (as opposed to the Vanderberg Air Force Station in California, for polar orbits). This tool is designed to help you launch all your projects without having to worry about the slow, mundane early stages of setting up. Hence we chose the spaceship-launch naming theme.

While both sites are important, Canaveral just sounded cooler than Vanderberg.

## Getting Started

### Dependencies

#### npm

Canaveral uses `npm` to install useful dependencies for you. To ensure that Canaveral works properly, please make sure you have `npm` installed: [Node.js Download](https://nodejs.org/en/)

#### Github

If you want to use Canaveral's Github features, you will need a github account, as well as a personal access token. Find out how to get a personal access token [here](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line).

#### VSCode

Canaveral has support for opening projects in [VSCode](https://code.visualstudio.com/). To enable this, you will need to make sure you can run the `code` command from the command line. Mac OS users will need to install this through VSCode if they have not already. Instructions can be found [here](https://code.visualstudio.com/docs/editor/command-line#_launching-from-command-line).

### Installation

There are **three ways** to install Canaveral: Homebrew, Go get/install, or download the executable.

#### Method 1: Homebrew (MacOS)

This method might only work for MacOS users, since it requires `brew`.

1. Make sure you have [Homebrew](https://brew.sh/) installed and updated (`brew update`).

2. Run:

   ```bash
   $ brew tap jchengjr77/homebrew-private https://github.com/jchengjr77/homebrew-private.git
   $ brew install canaveral
   ```

3. If there are no errors, you should be done! Check that canaveral is properly installed by running:

   ```bash
   $ which canaveral
   ```

   A path to the executable `canaveral` should be printed.

#### Method 2: Go build

1. First, you need to have Go installed. If you don't, follow the instructions [here](https://golang.org/doc/install).

   Go should have automatically configured your GOPATH to be `~/go`, or another reasonable default. If you are new to Go, you may need to add `$GOPATH/bin` to your PATH, so you can execute Go programs from the command line. _NOTE: Windows users must add to their PATH a different way. See link to instructions below._

   To the end of your `.zshrc` or `.bashrc` (whatever your shell's config file is), you can add: 

   `export PATH=$PATH:$(go env GOPATH)/bin` 

   or

   `PATH=$PATH:$(go env GOPATH)/bin`

   **For Windows users:** [Add to PATH Windows 10](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/)

   To check your current gopath, run `go env gopath`. For more information, run `go help gopath`.

2. Then, `go get` Canaveral and install it:

```bash
$ go get github.com/jchengjr77/canaveral
$ go install github.com/jchengjr77/canaveral
```

NOTE: `go get` can take a while to run (30-60s). Allow it to run for a couple minutes if necessary.

This should have put an executable named `canaveral` into the folder `$GOPATH/bin`. If you set your PATH 	to include `$GOPATH/bin`, then you should be good to go.

3. Check that Canaveral is properly installed by running:

```bash
$ which canaveral
```

And the path that is returned should be `$GOPATH/bin`, where `$GOPATH` is replaced with your actual gopath. For instance, JJ's canaveral will install in `~/go/bin`, because `~/go` is the `$GOPATH`.

#### Method 3: Download Executable

In the [Canaveral Releases](https://github.com/jchengjr77/canaveral/releases) section, you will find all current releases of Canaveral.
We suggest you grab the latest one: [v0.6.0](https://github.com/jchengjr77/canaveral/releases/tag/v0.6.0)

Select the correct package for your computer system and download it. See release notes for guidance.
After you download and unzip Canaveral, you will be left with a single executable. **You need to add it to your \$PATH.**

Add the following line to your shell's config file (.bashrc, .zshrc for bash and zsh respectively).
_NOTE: Windows users must add to their PATH a different way. See link to instructions below._

```bash
$ export PATH=$PATH:$HOME/path/to/canaveral
```

Here are some helpful links related to adding to \$PATH:

[For Mac](https://apple.stackexchange.com/questions/41542/adding-a-new-executable-to-the-path-environment-variable)

[For Linux](https://askubuntu.com/questions/322772/how-do-i-add-an-executable-to-my-search-path)

[For Windows 10 (Important)](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/)

## Basic Usage

##### For a full list of commands, please refer to `canaveral -h`.

##### For help on a specific command, please refer to `canaveral [command name] -h`.

The most basic function of canaveral is **viewing all your projects**:

```bash
$ canaveral
```

This command may tell you to "specify a workspace." To **specify a workspace**:

```bash
$ canaveral space /path/to/your/workspace
```

##### What is a workspace?

Your canaveral workspace is the place where you will put all your projects. A workspace is a path to a folder of your choosing. Every time you use canaveral, you will be interacting with projects in your workspace. You can specify your workspace to be anywhere you want. For instance, JJ's workspace is `/Users/[omitted]/Documents/Personal/projects`. Note that workspace paths should start from root (`/`).

NOTE: Though it is not necessary, we recommend you add a script or alias to help you quickly navigate to your selected workspace. While developing canaveral, this simple little tool save us a lot of headache from navigation. An alias would look something like this: (more information on zsh aliases [here](https://blog.lftechnology.com/command-line-productivity-with-zsh-aliases-28b7cebfdff9))

```c
// inside your .zshrc or .bashrc
alias gotoprojs="cd ~/path/to/canaveral/workspace"
```

After you specify your workspace, running `canaveral` should show a list of projects in your workspace. At this point, it may be the case that your workspace is empty. Lets **launch a new project**!

```bash
$ canaveral launch coolproject
```

Congratulations! You have launched your first canaveral project, named `coolproject`. However, `coolproject` has nothing in it.

For most people, a blank project like `coolproject` is not very useful. How about we **remove a project**, in this case `coolproject`, and **create a new [React](https://reactjs.org/) project**!

```bash
$ canaveral remove coolproject
... (confirm project deletion) ...
$ canaveral launch -type react coolreactproject
```

Now we have a new react project called `coolreactproject`! If you did not previously have a tool called `create-react-app` installed, Canaveral took care of that installation for you. Let's take a look at what our new project contains by **opening our project in [VSCode](https://code.visualstudio.com/)**. (Note: this requires you to have [VSCode](https://code.visualstudio.com/) as well as the [VSCode Command](https://code.visualstudio.com/docs/editor/command-line#_launching-from-command-line) installed).

```bash
$ canaveral code coolreactproject
```

This should open a new VSCode window showing your new project `coolreactproject`.

Additionally, Canaveral can make a new Github repo for you! First, make sure you **add your git credentials**:

```bash
$ canaveral addgithub
... (prompt for username and personal auth token) ...
```

Again, for guidance on how to get a personal auth token from github, follow [this link](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line).

To **view your git credentials**:

```bash
$ canaveral printgithub
```

At any point in time, if you wish to **remove your git credentials**, you can do so using:

```bash
$ canaveral removegithub
```

Lets try to get canaveral to **initialize a new project and git repo**! This time, let's create a [node](https://nodejs.org/en/) project:

```bash
$ canaveral launch --gitinit --type node coolproject2
```

Now, your project `coolproject2` is git enabled! You start by adding your own remote repository, staging changes and pushing commits, and all that other good git stuff.

Suppose we hadn't initialized `coolproject2` as a git repo, we could do 

```bash
$ canaveral gitinit -p coolproject2
```

or navigate to the project directory and simply do 

```bash
$ canaveral gitinit
```

instead. Now suppose that we've made some changes to `coolproject2` that we want to add to git. We can stage all of our changes with

```bash
$ canaveral gitadd .
```

or only specific files with

```bash
$ canaveral gitadd file1 file2 file3 ...
```

and then we can commit these changes. Much like the git cli, we can either do

```bash
$ canaveral gitcommit
```

to open the vim editor, or 

```bash
$ canaveral gitcommit -m "My commit message."
```

to commit our changes with the message "My commit message." One new git feature introduced by Canaveral is the concept of commit reminders. Suppose that we've added some print statements to the file `MyFile.go` and we want to remove them before we commit our changes. Rather than relying on our memory, we can do

```bash
$ canaveral gitremind MyFile.go "Remove the print statements on lines 13, 29, and 81."
```

then when we do

```bash
$ canaveral gitcommit MyFile.go
```

we will be prompted with a reminder informing us that we have some reminders setup for `MyFile.go`. At the time of commit we can then choose whether or not to display them. If we do choose to see the reminders, we're also given the option to cancel the commit if we no longer want to go through with it.

**Note: Currently `canaveral gitinit` and `canaveral gitadd` are equivalent to simply doing `git init` and `git add`. However, these functions exist as they may interface with the `canaveral gitremind` command in the future. However, `git commit` is not the same as `canaveral gitcommit` in that `git commit` will not print reminders.**

Suppose our project has hit a nasty bug and we've decided to configure the handy-dandy debugger in VSCode. We can quickly solve our problem, but configuring this tool creates a pesky `.vscode/` folder that git wants to track every time we commit. Rather than manually ignoring the folder, we can do

```bash
$ canaveral gitignore .vscode/
```

and `.vscode/` will be appended to the `.gitignore` file. If you don't have a `.gitignore` yet, it will be created.

## Project Types

Canaveral should be usable for all developers. Creating projects, looking at all your projects, and removing projects are universal features that anybody can use. These are fundamental features to Canaveral.

Furthermore, if you use one of the following technologies, you're in luck. More features for you!
(This list will most likely expand, so keep an eye out).

-   [React.js](https://reactjs.org/)
-   [React Native](https://reactnative.dev/)
-   [Node.js](https://nodejs.org/en/)
-   [Python (miniconda)](https://docs.conda.io/en/latest/miniconda.html)
-   [C](https://www.cprogramming.com/)

## Troubleshooting

For fixing common problems, please refer to this list of steps:

-   Make sure all dependencies are installed properly. See [dependencies](#dependencies) section
-   Refer to canaveral's built in help for reference. Type `canaveral --help`
-   For help on a specific command, such as `launch`, type `canaveral launch -help`
-   If there are any other bugs, please contact either [JJ Cheng](mailto:jonathanchengjr77@gmail.com) or [Sean Prendi](sean.prendi@gmail.com). You may be able to help identify and patch an important aspect of Canaveral!

## Contributing

Since we want Canaveral to be for all developers, contributions are super welcome! Whether you fix a bug, provide user feedback, create a big feature PR, or make our documentation prettier, the help is deeply appreciated.

Contributors are especially welcome to add support for new technologies. The more diversity Canaveral receives, the more useful the tool is for everybody. If you are comfortable with a certain technology that isn't on the current list, please reach out!

For all contribution inquiries, please email [JJ Cheng](mailto:jonathanchengjr77@gmail.com).

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

## Credits

### Creators:

JJ Cheng - [github](https://github.com/jchengjr77) - [jjcheng.me](https://jjcheng.me)

Sean Prendi - [github](https://github.com/SeanPrendi) - [seanprendi.me](https://www.seanprendi.me)
