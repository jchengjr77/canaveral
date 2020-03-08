# Canaveral

Launch your new projects seamlessly. Canaveral automates all your project setup away.

##### I wanted to have a project manager that was fast, intuitive, and actually useful. So I made Canaveral.

Canaveral is a Command Line Interface (CLI) tool that can add, remove, and view your projects. It is a tool built by developers for developers. Don't worry about where your projects are located, how they are organized, what their names are, or how to set them up. Canaveral does that for you.

## The Name

Cape Canaveral Air Force Station is one of two main launch sites for U.S. space missions. It is used primarily for launching spacecraft into equatorial orbits (as opposed to the Vanderberg Air Force Station in California, for polar orbits). This tool is designed to help you launch all your projects without having to worry about the slow, mundane early stages of setting up. Hence we chose the spaceship-launch naming theme.

While both sites are important, Canaveral just sounded cooler than Vanderberg.

## Getting Started

WARNING: The CLI is still in development mode, and has yet to be packaged properly. If you want to jump in early, follow the instructions below. Otherwise, stay posted for the next stable release.

First, you need to have Go installed. Follow the instructions [here](https://golang.org/doc/install).

Be sure that your GOPATH is configured correctly, so you are able to execute go binaries:

```bash
$ export PATH=$PATH:$GOPATH/bin
```

We use [cli by urfave](https://github.com/urfave/cli) to create Canaveral. Install it using:

```bash
$ go get github.com/urfave/cli
```

Finally, clone this repo and install Canaveral:

```bash
$ git clone https://github.com/jchengjr77/canaveral.git
$ go install canaveral
```

## Features

Canaveral should be usable for all developers. Creating projects, looking at all your projects, and removing projects are universal features that any developer can use. These are fundamental features to Canaveral.

Additionally, Canaveral can make a new Github repo for you! Once you create a project, it will push a standard initial commit for you, and provide you with the link to your repo. Just give Canaveral your git credentials and it will do the rest.

However, if you use one of the following technologies, you're in luck. More features for you!

#### React.js

Project creation includes `create-react-app` functionality, as well as an automatically initialized README, personalized with your project's details.

This means you will get a `app.js`, your basic `package.json`, `app.json`, as well as your basic `public` , `node_modules`, and `src` folders, among other useful knick-knacks.

#### React Native

Project creation includes `expo` functionality, as well as an automatically initialized README, personalized with your project's details.

This means you get all the `expo` generated resources, and your standard `App.js`, `app.json`, `package.json`, `node_modules`, etc.

#### Node.js

Project creation includes `npm` functionality, as well as an automatically initialized README, personalized with your project's details.

We will walk you through a lean setup process for your initial `package.json`, and you will be all set up for your node project.

#### Python

Project creation automatically creates `src` and `tests` folders for your code organization. We also create a `setup.py` file as well as `requirements.txt`. Canaveral also automatically initializes a README, personalized with your project's details.

#### C

Project creation automatically creates subdirectories `/src` and `/include` for your modules and headers. Canaveral also gives you a basic `Makefile` that you can customize. Canaveral also automatically initializes a README, personalized with your project's details.

## Contributing

Since we want Canaveral to be for all developers, contributions are super welcome! Whether you fix a bug, provide user feedback, create a big feature PR, or make our documentation prettier, the help is appreciated.

Contribution is especially useful for adding support for new technologies. The more diversity Canaveral receives, the more useful the tool is for everybody. If you know a certain technology that isn't on the current list, please reach out!

For all contribution inquiries, please email [JJ Cheng](mailto:jonathanchengjr77@gmail.com)

## The Creators

Jonathan Cheng - [github](https://github.com/jchengjr77) - [jjcheng.me](https://jjcheng.me)

Sean Prendi - [github](https://github.com/SeanPrendi)
