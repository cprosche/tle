# tle

A TLE handling (parsing, etc.) package for Go

## Installation

To install the package, run

```bash
go get github.com/cprosche/tle
```

## What is a TLE?

A TLE (Two-Line Element Set) is a data format used to convey sets of orbital elements that describe the orbits of Earth-orbiting satellites. A TLE set consists of two lines of data. The first line is called the "title line" and the second line is called the "data line". The title line contains the name of the satellite and a timestamp. The data line contains the orbital elements.

## Source

Format description: https://en.wikipedia.org/wiki/Two-line_element_set

## License

This package is licensed under the MIT license. See the LICENSE file for more information.
