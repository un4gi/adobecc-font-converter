# Adobe CC Font Converter

![License](https://img.shields.io/github/license/un4gi/adobecc-font-converter?style=plastic) [![Twitter Follow](https://img.shields.io/twitter/follow/un4gi_io?label=%40un4gi_io&style=social)](https://twitter.com/un4gi_io)

Adobe CC Font Converter is a simple CLI tool designed to help convert 
licensed and activated Adobe CC fonts into proper font files.

When you activate fonts in Adobe CC, the font files are not stored in
the typical system fonts directory. Instead, Adobe stores the fonts 
in a hidden folder without file extensions, making them difficult to
locate and use with other software. This tool will make simple work
of accessing your licensed and activated Adobe CC fonts in a standard
font format of your choice. Simply run the tool as follows:

## Usage

```bash
adobecc-font-converter -f [file format] -o [output file path]
```

That's really it! Enjoy! :)

## Installation

### Using Go:

```bash
go install github.com/Un4gi/adobecc-font-converter@latest
```

## Build From Source

```bash
git clone https://github.com/Un4gi/adobecc-font-converter.git
cd ./adobecc-font-converter
go build .
```
