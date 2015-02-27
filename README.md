icns
====
![](/images/usage.gif)

## Description
icns provides to generate and set your own custom folder icons.

```bash
$ icns generate /Users/nekova/Picture/awesome.png
#=> Generate .icns file in current directory
```

## Usage
```bash
'icns' generate <path/to/image-file>
'icns' set (<path/to/image-file> | <path/to/icns-file>)
'icns' reset
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/nekova/icns
```

## Contribution

1. Fork ([https://github.com/nekova/icns/fork](https://github.com/nekova/icns/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[nekova](https://github.com/nekova)
