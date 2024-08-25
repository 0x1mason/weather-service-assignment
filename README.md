# weather-service-assignment

## Pre-requisistes
Go 1.22 or newer

## Usage
The application runs an HTTP listener on port `8888`. You may get forecast for a set of latitude and longitude coordinates by sending a `GET` requests to `/forecast/{latitude},{longitude}`. E.g.,


```shell
$ curl localhost:8888/forecast/39.7456,-97.0892
{"characterization":"hot","shortForecast":"Sunny"}
```

## Tests
You may run the unit test with

```shell
go test ./...
```

## Shortcuts
This is a very simple implementation and I relied solely on the stdlib. I skipped any validation and there is only one "happy" test case; typically I would cover error cases as well and make use of a proper assertion package such as `https://github.com/stretchr/testify`.