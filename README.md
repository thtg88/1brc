# 1 Billion Row Challenge (1BRC) - Golang

**Please note: this repository is a work-in-progress. More details will be added to this README once the code is ready.**

This repository provides a few possible solutions to the 1BRC problem, as originally detailed at https://github.com/gunnarmorling/1brc:

> A fun exploration of how quickly 1B rows from a text file can be aggregated

But here they are implemented in Golang (instead of Java).

## Requirements

Go 1.22. You can install and use it using [gvm](https://github.com/moovweb/gvm) with:

```bash
gvm install go1.22
gvm use go1.22
```

## Sequential Solution

This is contained under the [cmd/1brc-sequential/main.go](cmd/1brc-sequential/main.go) command.

Build it with:

```bash
go build -o 1brc-sequential cmd/1brc-sequential/main.go
```

Run it with:

```bash
./1brc-sequential
```

This solution starts a goroutine, where a producer reads from the CSV row by row, and publishes a `TemperatureReading` struct on a data channel.

The data channel is read by a consumer in a separate goroutine, which aggregates each reading one by one in a `CityStats` struct.

Once both the producer and consumer goroutines have completed, the slice of `CityStats`, sorted by city, is returned, and written back to disk in CSV format.

This solution processes 1B rows in ~7m30s.
