<a name="unreleased"></a>
## [Unreleased]


<a name="0.0.4"></a>
## [0.0.4] - 2021-11-04
### Docs
- update changelog
- update changelog

### Feat
- **nada-sension:** add an interval env var, to set time measure interval

### Fix
- **nada-sensio:** send date in right format
- **nada-transform:** change to info if date format is not right

### BREAKING CHANGE

New env var NADA_SENSIO_MEASURE_INTERVAL_MS


<a name="0.0.3"></a>
## [0.0.3] - 2021-11-02
### Docs
- update changelog

### Fix
- **nada-serve:** prevent user from querying invalid dates
- **nada-serve:** remove console print
- **nada-serve:** critical bug with range data selection


<a name="0.0.2"></a>
## [0.0.2] - 2021-11-02
### Build
- add bazel build files

### Chore
- add nada fullfiller
- update .env.example to reflect devcontainer config
- tidy up project
- update changelog
- **nada-serve:** update changelog file

### Docs
- update changelog
- update readme
- update readme
- update changelog
- **nada-serve:** :memo: update graphQL query example
- **nada-serve:** add initial changelog

### Fix
- **nada-serve:** rename nada-serve.env file to .nada-serve.env
- **nada-serve:** do an empty request to check authorization with influxdb
- **nada-serve:** prevent launching if auth params from influxdb are invalid

### BREAKING CHANGE

rename your .env file for nada-serve


<a name="0.0.1-alpha"></a>
## 0.0.1-alpha - 2021-10-31
### Build
- remove unused vscode extension
- add vscode build config
- add mosquitto clients to dev container
- remove mosquitto client installation in docker file
- install mosquitto client in dev container
- **nada-transform:** :bug: mend main + heroku mqtt

### Chore
- move to pkg folder
- add code workspace file
- initial project structure
- rename workspace
- update vendor deps
- update deps
- add new extension scope to workspace
- :construction: add devcontainer definition file
- **nada-serve:** update vendor
- **nada-serve:** update deps
- **nada-serve:** update generated model graphql definition files
- **nada-serve:** fix typo in yaml config file
- **nada-serve:** add gqlgen config
- **nada-serve:** update deps
- **nada-serve:** update deps
- **nada-serve:** add the csv airport set
- **nada-serve:** add a new scope to vscode settings
- **nada-transform:** ignore settings.json

### Doc
- **nada-serve:** add flux scripts

### Docs
- :memo: architecture
- fix to a supported valid nano seconds date format by js
- update graphql schema for nada-serve
- **nada-sensio:** :memo: instructions: global README -> local README
- **nada-serve:** update env example
- **nada-serve:** update api gql schema
- **nada-transform:** :art: presentation
- **nada-transform:** :art: application's purposes
- **nada-transform:** :art: readme
- **nada-transform:** :fire: readMe
- **nana-serve:** update query flux doc

### Feat
- :sparkles: Mosquitto subscriber and db client
- **nada-sension:** allow to end sensio with enter keyboard key
- **nada-serve:** initial dataloaders implementation
- **nada-serve:** server better airport info
- **nada-serve:** add a better config parser
- **nada-serve:** initial config loader implementation
- **nada-serve:** implement cors configuration
- **nada-serve:** initial database implementation
- **nada-serve:** handle properly heroku port
- **nada-serve:** rename .env config file, read nada-serve yml file
- **nada-serve:** implement getMeanMeasures
- **nada-serve:** add sanitize function for time and duration
- **nada-serve:** implement config loading for sensor measurement name definition
- **nada-serve:** implement dataloaders
- **nada-serve:** add a new setting for airportfile
- **nada-serve:** generate initial resolvers
- **nada-serve:** initial main implementation
- **nada-serve:** add default SI units
- **nada-transform:** :construction: paho pub sub
- **nada-transform:** :sparkles: hive mqqt connexion
- **nada-transform:** :sparkles: auto_reconnect
- **nada-transform:** :sparkles: log mod
- **nada-transform:** :sparkles: pub -> mqtt -> sub -> influxDB
- **nada-transform:** :construction: database import
- **nada-transform:** :construction: database import
- **nada-transform:** :construction: database client and mosquitto subscriber
- **nada-transform:** :truck: changed the host of the Mosquitto MQTT Broker
- **nada_serve:** implement array parsing from env vars

### Fix
- **nada-sensio:** :bug: link with nada-transform
- **nada-sension:** send date in RFC3339 nano
- **nada-serve:** return airport ID in case of error with csv file
- **nada-serve:** prevent using empty arrays in getSubsetOfSensor
- **nada-serve:** add graphql-faker schema definitions, force resolvers qlgen
- **nada-serve:** check if database is ready instead of health
- **nada-serve:** unique ids for meanmeasures
- **nada-transform:** :bug: .env access
- **nada-transform:** :bug: incorrect database model
- **nada-transform:** rename env variable NADA_TRANSFORM_INFLUXDB_USERMAIL
- **nada-transform:** :bug: path was raw in env init
- **nada-transform:** :construction: paho
- **nada-transform:** :bug: importing database internal package

### Refactor
- **nada-transform:** split database connection, improve db schema
- **nada-transform:** :building_construction: deleted external folder
- **nada-transform:** :building_construction: deleted external folder
- **nada-transform:** :art: _value
- **nada-transform:** :fire: duplicate database file
- **nada-transform:** :fire: vendor
- **nada-transform:** :fire: vendor
- **nada-transform:** :recycle: env variables
- **nada-transform:** :loud_sound: added logging

### BREAKING CHANGE

empty array arg will not return anymore every sensors when using getSubsetOfSensor

rename your env var to NADA_TRANSFORM_INFLUXDB_ORG


[Unreleased]: https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/compare/0.0.4...HEAD
[0.0.4]: https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/compare/0.0.3...0.0.4
[0.0.3]: https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/compare/0.0.2...0.0.3
[0.0.2]: https://github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/compare/0.0.1-alpha...0.0.2
