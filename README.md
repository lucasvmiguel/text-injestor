# text-injestor
[![Build Status](https://travis-ci.org/lucasvmiguel/text-injestor.svg?branch=master)](https://travis-ci.org/lucasvmiguel/text-injestor)
[![GoDoc](https://godoc.org/github.com/lucasvmiguel/text-injestor?status.svg)](https://godoc.org/github.com/lucasvmiguel/text-injestor)

## Overview

An application that can ingest text, calculate statistics on the data
& report back those stats.

## Architeture Explanation

[Here](https://medium.com/@LucasVieiraDev/dependencies-in-golang-projects-f46a11fef832)

## Installation

Make sure you have a working Go environment. [See
the install instructions for Go](http://golang.org/doc/install.html).

To install text-injestor, simply run:
```
git clone https://github.com/lucasvmiguel/text-injestor.git
```

## Configuration

```
[http]
port=":8080"
[http.handlers]
stats="/stats"
```

## Usage Without Docker

command:
```
$ make run
```

## Usage With Docker

command:
```
$ make docker-build
$ make docker-run
```

## Documentation

[Here](https://godoc.org/github.com/lucasvmiguel/text-injestor)