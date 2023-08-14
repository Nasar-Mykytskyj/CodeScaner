# CodeScaner

## What it does
I have implemented small console application for scanning source code and finding vulnarebilities

## Build and Run

**Command-line Git** Clone Code Scanner
```
Clone your fork of the repository into current directory
$ git clone https://github.com/Nasar-Mykytskyj/CodeScaner.git
```

For building application ensure that GOPATH includes the repo then type in command line 
```
$ cd CodeScanner
$ go build -o <buildPath> code_scanner/cmd/codeScanner
```
To application type in command line 
```
./<buildPath> -src <path to src code directory> -format <json or txt> -out <path for results>
```

**Run options**
|<span> | |
|--|--|
| **-src**|  path to src code directory |
| **-format**| CodeScanner output format |
| **-out**| CodeScanner output path |
