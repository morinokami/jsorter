## jsorter

jsorter sorts a JSON input and outputs the result.

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
