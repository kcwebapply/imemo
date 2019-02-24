# imemo
<img  align="right" src="image/imemo.png" width="100px">

[![GoDoc](https://godoc.org/github.com/kcwebapply/imemo?status.svg)](https://godoc.org/github.com/kcwebapply/imemo)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/imemo)
[](https://github.com/gin-gonic/gin/releases)
[![Release](https://img.shields.io/github/release/kcwebapply/imemo.svg?style=flat-square)](https://github.com/kcwebapply/iemo/release)

`imemo` is very simple `memo` tool for terminal.

![sample-demo](https://imgur.com/56c9MuR.gif)

## Installation

### On macOS

```
brew tap kcwebapply/imemo
brew install imemo
```

## Usage

### show memo list
listing all memo. `imemo a` also available.

```
> imemo all 
----------------------------------------------------------------------------------
| 1| playing tennis with Mike on Next  Tuesday                                   |
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
```

### save new memo
please input what you want to write on memo.

![sample-demo](https://imgur.com/whWhi0X.gif)

`imemo s` also available.

```
> imemo save "meeting at 13:30"
----------------------------------------------------------------------------------
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
memo saved!
```

## delete memo
please input memo's Id which you want to delete.

![sample-demo](https://imgur.com/ekB0A1M.gif)

```
> imemo d 2
----------------------------------------------------------------------------------
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
memo deleted!
```

## edit memo
`imemo edit` edit memo data.

![sample-demo](https://imgur.com/56AxBqS.gif)

```
> imemo e 
>> xxx  // input new memo.
memo edited!
```

## clear memo
`imemo clear` delete all memo data.

![sample-demo](https://imgur.com/NvdI1Mi.gif)

```
> imemo clear
clear 2 memo !
```



