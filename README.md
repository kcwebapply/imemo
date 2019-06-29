# imemo
<img  align="right" src="image/imemo.png" width="100px">

[![GoDoc](https://godoc.org/github.com/kcwebapply/imemo?status.svg)](https://godoc.org/github.com/kcwebapply/imemo)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/imemo)
[](https://github.com/gin-gonic/gin/releases)
[![Release](https://img.shields.io/github/release/kcwebapply/imemo.svg?style=flat-square)](https://github.com/kcwebapply/iemo/release)

`imemo` is very simple `memo` tool for terminal.

![sample-demo](https://imgur.com/La20wdF.gif)

## Installation

### On macOS

```
brew tap kcwebapply/imemo
brew install imemo
```

## Usage

### Ls
`imemo ls` list all memo.  

![sample-demo](https://imgur.com/UkavgEa.gif)

```
> imemo ls 
----------------------------------------------------------------------------------
| 1| playing tennis with Mike on Next  Tuesday                                   |
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
```

### add new memo
`imemo add` add what you want to write.

![sample-demo](https://imgur.com/WlfzW7Z.gif)


```
> imemo add "meeting at 13:30"
----------------------------------------------------------------------------------
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
memo saved!
```

## delete memo
please input memo's Id which you want to delete.

![sample-demo](https://imgur.com/9y2Osxq.gif)

```
> imemo rm 2
----------------------------------------------------------------------------------
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
memo deleted!
```


## clear memo
`imemo clear` delete all memo data.

![sample-demo](https://imgur.com/6c6dJPt.gif)

```
> imemo clear
clear 2 memo !
```



