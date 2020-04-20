convertMan [![CI](https://github.com/aiziyuer/convertMan/workflows/CI/badge.svg)](https://github.com/aiziyuer/convertMan/actions?query=workflow%3ACI) [![Release](https://github.com/aiziyuer/convertMan/workflows/Release/badge.svg)](https://github.com/aiziyuer/convertMan/releases/latest)
[![GitHub Releases](https://img.shields.io/github/downloads/aiziyuer/convertMan/latest/total)](https://github.com/aiziyuer/convertMan/releases/latest)
---

`convertMan` is a tool help you convert between `yaml`/`xml`/`json`/`toml`independently.


## ⚙ Installation

``` bash
CGO_ENABLED=0 \
GOBIN=/usr/bin \
go get -u -v github.com/aiziyuer/convertMan
```

## ⚡ Quickstart

``` bash
# default output format is json
➜  convertMan <<EOF
teenagers:
  - body
  - girl
EOF
{"teenagers":["body","girl"]}

# options output format: yaml,json,toml,ini,xml
➜  convertMan -o toml <<EOF
teenagers:
  - body
  - girl
EOF
teenagers = ["body", "girl"]

➜  convertMan -h
Usage:
  convertMan [input file] [flags]

Flags:
      --config string    location of config files like $CONVERT_MAN_CONFIG  (default "/root/.convertMan")
  -h, --help             help for convertMan
  -i, --input string     input format: auto(default for file), yaml, json, ini, xml, toml (default "yaml")
  -o, --output string    output format: yaml, json, ini, xml, toml (default "json")
  -v, --verbose string   Log level (trace, debug, info, warn, error, fatal, panic)  (default "warning")
```

## 🤖 More

combine multiple yaml config with default.
``` bash
➜  ~ cat<<EOF>default.yaml
data:
  a: 1
  b: 2
  c: 3
EOF

➜  ~ cat<<EOF>new.yaml
data:
  a: 22222
EOF

➜  ~ convertMan default.yaml new.yaml
data:
    a: 22222
    b: 2
    c: 3
```

## 🎯 Features

## ⭐ FAQ
