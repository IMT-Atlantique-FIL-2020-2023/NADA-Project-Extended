<div id="top"></div>

[![Build][build-shield]][build-url]
[![Version][version-shield]][version-url]
[![Language][language-shield]][language-url]
[![Roadmap][roadmap-shield]][roadmap-url]
[![License][license-shield]][license-url]

<br />
<div align="center">
<a href="https://nada-extended.herokuapp.com/">
<img src="./docs/NADA.svg" alt="Logo" width="200" height="200">
</a>

<h2 align="center">N.A.D.A Extended</h2>

<p align="center">National Atmospheric Data | Capture data from airport sensor and visualize it.</p>
<p align="center">
<a href="https://nada-extended.herokuapp.com/">
<strong>Browse the NADA-serve deployment »</strong>
</a>
<br />
<br />
<a href="https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-webapp/">View Front</a>
·
<a href="https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/issues">Report Bug</a>
·
<a href="https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/issues">Request Feature</a>
</p>
</div>

<!-- TABLE OF CONTENTS -->
## Table of Contents

- [Table of Contents](#table-of-contents)
- [About The Project](#about-the-project)
  - [General purpose](#general-purpose)
  - [General architecture](#general-architecture)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Nada-sensio (MQTT Publisher)](#nada-sensio-mqtt-publisher)
- [nada-transform (MQTT subscriber + database insert)](#nada-transform-mqtt-subscriber--database-insert)
- [nada-serve (GraphQL API for NADA-Webapp)](#nada-serve-graphql-api-for-nada-webapp)
  - [Source code](#source-code)
  - [Launching](#launching)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## About The Project

### General purpose

The NADA project is a student project.
The team is composed of @creeperdeking, @RaphaelPainter, @maxerbox & @mlhoutel. They are IT students at IMT Atlantique, France.
It allows an user to view on a web page (NADA-webapp) aggregated sensor data captured from sensors (NADA-sensio) served by a GraphQL Api (NADA-serve)

### General architecture

This is an IoT application using [Paho](https://github.com/eclipse/paho.mqtt.golang) MQTT Client.

<img src="https://raw.githubusercontent.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/develop/assets/architecture.PNG" alt="Architecture" width="50%" align="center">

Frontend made in VueJS 3 + Vite : [NADA-Webapp](https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-webapp)

## Getting Started

### Prerequisites

- MQTT broker : [Mosquitto](https://github.com/eclipse/mosquitto)
- [Influxdb v2](https://www.influxdata.com/)

### Installation

- git clone <repository>

We do provide an .devcontainer file ! This mean that if you have the vscode extension installed, project will be operational in a minute. Just click on "Reopen in container popup". Every .env.example file is configured to be operational with the container by default, so you just have to rename those files !

__More infos :__

Name: Remote - Containers  
Id: ms-vscode-remote.remote-containers  
Description: Open any folder or repository inside a Docker container and take advantage of Visual Studio Code's full feature set.  
VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers

### Configuration

- create and configure [NADA-Sensio](.nada-sensio.env.example), [NADA-Transform](.nada-transform.env.example), and [NADA-Serve](.nada-serve.env.example) .env files (follow links to see examples)

## Nada-sensio (MQTT Publisher)

- Simulates a sensor, publishing data to the configured MQTT Broker

go run cmd/nada-sensio/main.go  [sensorID] [airportID] [measureType]

Accepted measureType:

- temperature
- altitude
- pressure
- latitude
- longitude
- windspeed
- winddirx
- winddiry

**example**:
> go run cmd/nada-sensio/main.go S0 A0 windspeed 

## nada-transform (MQTT subscriber + database insert)

- Subscribes to a configured MQTT topic  
- Inserts received data from NADA-Sensio into the configured influxDB database

**example**:
> go run cmd/nada-transform/Main.go

## nada-serve (GraphQL API for NADA-Webapp)

- Query influxdb database and read data inserted by nada-transform
- Serve data through a graphql api

GraphQL Schema [available here](./api/schema.graphql)  
Influxdb example queries made by nada-webapp [available here](./scripts/queries.flux)  
We are providing a heroku instance with a graphql playground at [https://nada-extended.herokuapp.com/](https://nada-extended.herokuapp.com/)  

### Source code

- [main.go](./cmd/nada-serve)
- [internal code](./internal/app/nada-serve)

### Launching

Be sure to have a file named `.nada-serve.env` in your current working directory. You can take a look at an [example](./.nada-serve.env) to create it

`go run cmd/nada-serve/main.go`

## Contributing

You can use the .devcontainer to quick bootstrap the project

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

- [Img Shields](https://shields.io)

<p align="right">(<a href="#top">back to top</a>)</p>

[build-shield]: https://img.shields.io/github/workflow/status/IMT-Atlantique-FIL-2020-2023/NADA-extended/Deploy%20to%20heroku./main?style=flat-square
[build-url]: https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/blob/main/.github/workflows/serve.yml
[version-shield]: https://img.shields.io/github/go-mod/go-version/IMT-Atlantique-FIL-2020-2023/NADA-extended?style=flat-square&color=orange
[version-url]: https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/blob/main/go.mod
[language-shield]: https://img.shields.io/github/languages/top/IMT-Atlantique-FIL-2020-2023/NADA-extended?style=flat-square
[language-url]: https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/search?l=go
[roadmap-shield]: https://img.shields.io/badge/roadmap-available-brightgreen?style=flat-square
[roadmap-url]: https://github.com/orgs/IMT-Atlantique-FIL-2020-2023/projects/1
[license-shield]: https://img.shields.io/github/license/IMT-Atlantique-FIL-2020-2023/NADA-extended?style=flat-square
[license-url]: https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/blob/main/LICENSE/
[logo]: src/assets/logo.png
