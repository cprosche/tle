# tle

[![Go Reference](https://pkg.go.dev/badge/github.com/cprosche/tle.svg)](https://pkg.go.dev/github.com/cprosche/tle)

A TLE handling package for Go

## Installation

To use the package, run:

```bash
go get github.com/cprosche/tle
```

## Features

- Provides TLE type
- Parses TLEs and 3LEs
- Validates checksums
- Supports Alpha 5 format

## Example

```go
package main

import (
    "fmt"
    "github.com/cprosche/tle"
)

func main() {
    tleStr := `ISS (ZARYA)
1 25544U 98067A   21293.61892701  .00000913  00000-0  26179-4 0  9991
2 25544  51.6443  88.9783 0004266  35.1673  33.6829 15.48976801304895`

    tle, err := tle.Parse(tleStr)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(tle)
}
```

## What is a TLE (or 2LE or 3LE)?

> **TLE or 2LE**: Two-Line Element Set

> **3LE**: Three-Line Element Set

TLEs are used to describe the orbits of Earth-orbiting objects. They are used by the United States Department of Defense to track objects in orbit.

Recent TLEs for most objects can be found at the following websites:

- [Space-Track](https://www.space-track.org/)
- [Celestrak](https://www.celestrak.com/)
- [N2YO](https://www.n2yo.com/)

## Technical References

- https://www.space-track.org/documentation#tle
- https://en.wikipedia.org/wiki/Two-line_element_set

## License

This package is licensed under the MIT license. See the LICENSE file for more information.
