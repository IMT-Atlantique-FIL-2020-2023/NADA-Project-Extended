<div id="top"></div>

[![Build][build-shield]][build-url]
[![Version][version-shield]][version-url]
[![Language][language-shield]][language-url]
[![Roadmap][roadmap-shield]][roadmap-url]
[![License][license-shield]][license-url]

<br />
<div align="center">
<a href="https://nada-extended.herokuapp.com/">
<img src="https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-webapp/blob/main/src/assets/NADA.svg" alt="Logo" width="200" height="200">
</a>

<h2 align="center">N.A.D.A Extended</h2>

<p align="center">National Atmospheric Data | Capture data from airport sensor and visualize it.</p>
<p align="center">
<a href="https://nada-extended.herokuapp.com/">
<strong>Browse the deployment »</strong>
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

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

This is an IoT application example using <a href="https://github.com/eclipse/paho.mqtt.golang">Paho</a> MQTT Client.

<br />
<img src="https://raw.githubusercontent.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/develop/assets/architecture.PNG" alt="Architecture" width="50%" align="center">


<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

- [TODO](...)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started


### Prerequisites

- MQTT broker (example: <a href="https://github.com/eclipse/mosquitto">Mosquitto</a>)
- Time series database (example: <a href="https://www.influxdata.com/">InfluxDB</a>)

### Installation

> `// TODO`

<p align="right">(<a href="#top">back to top</a>)</p>

## Launching 

### Nada-sensio (MQTT Publisher)
- Simulates a sensor publishing data to the configured MQTT Broker

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

example: go run cmd/nada-sensio/main.go S0 A0 windspeed 

### nada-transform (MQTT subscriber + db insert)
- Subscribes to a configured MQTT topic  
- Inserts received data into the configured influxDB database

go run cmd/nada-transform/Main.go


<!-- CONTRIBUTING -->

## Contributing

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


NADA-SENSIO
Message Broker
NADA-SERVE
NADA-DB
NADA-TRANSFORM
