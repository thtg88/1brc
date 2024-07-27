# 1 Billion Row Challenge (1BRC) - Golang

**Please note: this repository is a work-in-progress. More details will be added to this README once the code is ready.**

This repository provides a few possible solutions to the 1BRC problem, as originally detailed at https://github.com/gunnarmorling/1brc:

> A fun exploration of how quickly 1B rows from a text file can be aggregated

But here they are implemented in Golang (instead of Java).

## Equipment

Tests are performed locally on a MacBook Pro (15-inch, 2016) with a 2.6 GHz Quad-Core Intel Core i7 and 16 GB 2133 MHz LPDDR3.

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

## Buffered Sequential Solution

This is contained under the [cmd/1brc-buffered-sequential/main.go](cmd/1brc-buffered-sequential/main.go) command.

Build it with:

```bash
go build -o 1brc-buffered-sequential cmd/1brc-buffered-sequential/main.go
```

Run it with:

```bash
./1brc-buffered-sequential
```

This solution works similarly to the [Sequential Solution one](#sequential-solution), with the exception that temperature readings are produced and consumed on a buffered channel.

This solution processes 1B rows in ~5m30s.

## Not Calculating The Average Temperature On Every Reading

In an effort to reduce floating point calculations during consumption of temperature readings, I have attempted to skip those and only leaving that to the end when the final stats are needed. This is regulated by a config variable `CalculateAverageForEachReading bool`. This however does not seem to have any impact on the overall time spent by the 2 solutions above.

## Raw File Read Producer Solution

This is contained under the [cmd/1brc-raw-file-read/main.go](cmd/1brc-raw-file-read/main.go) command.

Build it with:

```bash
go build -o 1brc-raw-file-read cmd/1brc-raw-file-read/main.go
```

Run it with:

```bash
./1brc-raw-file-read
```

This solution changes the main producer to not use native CSV line reading, but to read a set of characters and building the rows there by splitting by new-line character and comma separator.

This solution processes 1B rows in ~3m45s.
