# doauto

A simple command line utility for automating the creation of DigitalOcean Droplets

## Description

`doauto` helps in automatically creating [DigitalOcean](https://www.digitalocean.com) Droplets based on the config file within the same directory as the doauto executable.

## Getting Started

### Dependencies

* [Viper](https://github.com/spf13/viper) - Configuration library for Golang
* [godo](https://github.com/digitalocean/godo) - Go client library for DigitalOcean's API

### Installing

```bash
~$ go get https://github.com/Pensu/doauto.git
~$ cd doauto
doauto$ go run main.go
```

### Executing program

* Add your [DigitalOcean API](https://docs.digitalocean.com/reference/api/api-reference/) key to the `config.yaml` file
* Edit the `dropletName`, `region`, `size`, and `image` fields to fit your needs. Examples are provided within the `config.yaml` file.
* Once the configuration is set, run `go run main.go`

### Help

If you run into any errors, please feel free to open an issue on GitHub. 

## Authors

* [Pensu](https://github.com/Pensu)

## License

This project is licensed under the Apache License 2.0 License - see the LICENSE file for details

