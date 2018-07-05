# text-injestor
[![Build Status](https://travis-ci.org/lucasvmiguel/text-injestor.svg?branch=master)](https://travis-ci.org/lucasvmiguel/text-injestor)
[![GoDoc](https://godoc.org/github.com/lucasvmiguel/text-injestor?status.svg)](https://godoc.org/github.com/lucasvmiguel/text-injestor)

## Overview

An application that can ingest text, calculate statistics on the data
and report back those stats.

Usually, I prefer a code without a lot of comments, but I think it is the best way to explain my logic in this case.

## Architeture Explanation

There are 3 main packages:
* api - Responsable for run the http api
* handlers - The http handlers
* textanalyzer - the package responsable to index and show some analytics about texts
  * I wouldn't need to store the words total, characters total, etc, but I think that processing the text once and store all kind of information is better than processing for each kind of information


More about the architecture chosen:
* [dependencies](https://medium.com/@LucasVieiraDev/dependencies-in-golang-projects-f46a11fef832)
* [configuration](https://medium.com/@LucasVieiraDev/configuration-in-golang-packages-6bcda76d011b)

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

## Testing

### Unit

command:
```
$ make test-unit
```

### E2E

command:
```
$ make test-e2e
```

## Documentation

[Here](https://godoc.org/github.com/lucasvmiguel/text-injestor)