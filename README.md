## jsorter

[![CircleCI](https://circleci.com/gh/morinokami/jsorter.svg?style=svg)](https://circleci.com/gh/morinokami/jsorter)

jsorter sorts a JSON input and outputs the result.

![Screenshot](./demo.gif)

### Installation

```sh
$ go get -u github.com/morinokami/jsorter
```

### Usage

```sh
$ cat sample.json
{
  xxx
}
$ jsorter < sample.json
{
  xxx
}
$ jsorter -r < sample.json
{
  xxx
}
```
