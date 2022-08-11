# invaps (inverter value Prometheus service)

Invaps is a tool which reads inverter processdata values from a MariaDB database and acts as endpoint which can be scraped from Prometheus monitoring solution.

## Description / Overview

Invaps is one of several building blocks for generating a Grafana dashboard for Kostal Plenticore inverters. Invafetch reads the Processdata values at regular intervals from the Inverter API and stores the results in JSON format in a MariaDB table. The Invaps tool uses these values, i.e. reads them and makes them available to Prometheus on request. Grafana, in turn, uses Prometheus as a data source to create a dashboard for the Kostal Plenticore inverter. Here, a modular concept was implemented so that one application, as small as possible, is responsible for a single task at a time. The MariaDB database serves as the interface for the Invafetch and Invaps tools and thus as a buffer for the Processdata values. For a description of Invafetch see [https://github.com/geschke/invafetch](https://github.com/geschke/invafetch), a complete example including definition of the Grafana dashboard and a Docker compose file to start all components in a Docker environment can be found at [https://github.com/geschke/grkopv-dashboard](https://github.com/geschke/grkopv-dashboard).


## Installation

The recommended installation method is to use the Docker image, a fully commented example can be found at [https://github.com/geschke/grkopv-dashboard](https://github.com/geschke/grkopv-dashboard). Besides that, invaps can be installed from source like any other program written in Go:

```text
$ git clone https://github.com/geschke/invaps
$ cd invaps/
$ go build
$ go install
```

This command builds the invaps command, producing an executable binary. It then installs that binary as $HOME/go/bin/invaps (or, under Windows, %USERPROFILE%\go\bin\invaps.exe).
Thus invaps can be started simply in the command line.


## Configuration

A .env file is used for configuration, which must be located either in the current directory, in a ./config or /config directory. Invaps requires access to the database used by invafetch. As with invafetch, invaps requires read and write access, as the stored values are currently deleted after two days to reduce space requirements.

In addition to the database connection, the web server port can also be configured, with this set to 8080 by default.

Overview of the configuration options:

|Name of environment variable|Defaults|Example|Hint|
|----------------------------|--------|-------|----|
|DBHOST|(empty)|"MARIADB DATABASE SERVER"|database server|
|DBUSER|(empty)|"DATABASE USERNAME"|database username|
|DBNAME|(empty)|"DATABASE NAME"|name of database|
|DBPASSWORD|(empty)|"DATABASE PASSWORD"|password of database user|
|DBPORT|"3306"|"3306"|MariaDB port (optional)|
|PORT|"8080"|"8080"|Webserver port|

## Quick Start

Invaps is built on top of the [Gin](https://gin-gonic.com/) HTTP web framework. Gin uses the environment variable GIN_MODE to set up debug mode, which contains additional output not required for operation. If GIN_MODE is not set, debug mode is enabled; for operation and to disable debug mode, set GIN_MODE=release.

When invaps is started, the configuration file .env is read. If this is missing or no connection can be established, invaps is terminated.

A successful start of invaps looks like the following:

```text
$ ./invaps
2022/08/11 18:10:56 invaps starting on port 8080...
2022/08/11 18:10:56 in recordCurrentValues again!
2022/08/11 18:10:56 in recordcurrentValues again with last values!
[...]
```

Invaps then makes the inverter metrics available at the URL http://[server][:port]/metrics.

## License

Invaps is licensed under the MIT License. You can read the full terms here: [LICENSE](LICENSE).