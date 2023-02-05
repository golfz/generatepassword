# GetPass
Password Generator

## Installation
```bash
$ go install github.com/golfz/getpass
```

## Help
```bash
$ getpass -h

Usage: getpass [-l <length>] [-no-upper] [-no-lower] [-no-number] [-no-special] [-h]
    -l <int>      length of password (default 16)
    -no-upper     exclude upper case letters
    -no-lower     exclude lower case letters
    -no-number    exclude numbers
    -no-special   exclude special characters
    -h            show help
```

## Example
```bash
$ getpass -l 32
cVD-B+m2RiAaQRFU!g]XcQaWMw50!Ky!
```

