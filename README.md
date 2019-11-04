## jsorter

[![CircleCI](https://circleci.com/gh/morinokami/jsorter.svg?style=svg)](https://circleci.com/gh/morinokami/jsorter)

jsorter sorts a JSON input and outputs the result.

![Screenshot](./demo.gif)

### Installation

```sh
$ go get -u github.com/morinokami/jsorter
```

### Usage

Just supply any JSON data to the standard input of `jsorter`.
Note that `jsorter` doesn't sort array values, since order matters for arrays.
Also note that each value in an array would be sorted if its type is an object.

```sh
$ cat sample.json
{
  "menu": {
    "id": "file",
    "value": "File",
    "popup": {
      "menuitem": [
        { "value": "New", "onclick": "CreateNewDoc()" },
        { "value": "Open", "onclick": "OpenDoc()" },
        { "value": "Close", "onclick": "CloseDoc()" }
      ]
    }
  }
}
$ jsorter < sample.json # Redirection
{
  "menu": {
    "id": "file",
    "popup": {
      "menuitem": [
        {
          "onclick": "CreateNewDoc()",
          "value": "New"
        },
        {
          "onclick": "OpenDoc()",
          "value": "Open"
        },
        {
          "onclick": "CloseDoc()",
          "value": "Close"
        }
      ]
    },
    "value": "File"
  }
}
$ jsorter -r < sample.json # Reverse order
{
  "menu": {
    "value": "File",
    "popup": {
      "menuitem": [
        {
          "value": "New",
          "onclick": "CreateNewDoc()"
        },
        {
          "value": "Open",
          "onclick": "OpenDoc()"
        },
        {
          "value": "Close",
          "onclick": "CloseDoc()"
        }
      ]
    },
    "id": "file"
  }
}
$ curl -s https://jsonplaceholder.typicode.com/users/1 | jsorter # Piping
{
  "address": {
    "city": "Gwenborough",
    "geo": {
      "lat": "-37.3159",
      "lng": "81.1496"
    },
    "street": "Kulas Light",
    "suite": "Apt. 556",
    "zipcode": "92998-3874"
  },
  "company": {
    "bs": "harness real-time e-markets",
    "catchPhrase": "Multi-layered client-server neural-net",
    "name": "Romaguera-Crona"
  },
  "email": "Sincere@april.biz",
  "id": 1,
  "name": "Leanne Graham",
  "phone": "1-770-736-8031 x56442",
  "username": "Bret",
  "website": "hildegard.org"
}
```
