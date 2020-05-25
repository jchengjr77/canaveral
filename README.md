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

WARNING: The CLI is still in development, and has yet to be packaged properly. If you want to jump in early, follow the instructions below. Otherwise, stay posted for the next stable release.

There are two ways to install Canaveral.

#### Method 1: Go build

First, you need to have Go installed. If you don't, follow the instructions [here](https://golang.org/doc/install).

Be sure that your GOPATH is configured correctly, so you are able to execute go binaries.

Then, `go get` Canaveral and install it:

```bash
$ go get github.com/jchengjr77/canaveral
$ go install github.com/jchengjr77/canaveral
```

This should have put an executable named `canaveral` into the folder `$GOPATH/bin`. If Canaveral isn't working, check that folder to see if the executable really exists.

#### Method 2: Download Executable

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

Your canaveral workspace is the place where you will put all your projects. A workspace is a path to a folder of your choosing. Every time you use canaveral, you will be interacting with projects in your workspace. You can specify your workspace to be anywhere you want. For instance, JJ's workspace is `/Users/[omitted]/Documents/Personal/projects`. Note that workspace paths should from root (`/`).

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

Lets try to get canaveral to **initialize a new project and repo**! This time, let's create a [node](https://nodejs.org/en/) project:

```bash
$ canaveral launch --gitinit --type node coolproject2
```

Now, your project `coolproject2` is git enabled! You start by adding your own remote repository, staging changes and pushing commits, and all that other good git stuff.

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
