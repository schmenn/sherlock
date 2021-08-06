# Sherlock

basically [sherlock](https://github.com/sherlock-project/sherlock), but in go

## Installation

### before go 1.17
`go get -u github.com/Schmenn/sherlock`

### go 1.17 and above
`go install github.com/Schmenn/sherlock@latest`

## Usage

`sherlock <username>`

### Docker

Build the image with:

`docker build -t sherlock .`

And run it with:

`docker run -t sherlock someUsername`
