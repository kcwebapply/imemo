# imemo
<img  align="right" src="https://imgur.com/qxEBvis.png" width="150px">

[![GoDoc](https://godoc.org/github.com/kcwebapply/imemo?status.svg)](https://godoc.org/github.com/kcwebapply/imemo)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/imemo)
[](https://github.com/gin-gonic/gin/releases)
[![Release](https://img.shields.io/github/release/kcwebapply/imemo.svg?style=flat-square)](https://github.com/kcwebapply/iemo/release)

`imemo` is very simple `memo` tool for terminal.


- [ls (show memo list)](#ls)
- [add (add memo)](#add)
- [rm (delete memo)](#rm)
- [clear (delete all memo)](#clear)

<! --<img src="https://imgur.com/La20wdF.gif" width="500px" height="300px"> -->

## Installation

### On macOS

```
brew tap kcwebapply/imemo
brew install imemo
```

## Usage

### ls
`imemo ls` list all memo.  


<img src="https://imgur.com/UkavgEa.gif" width="500px">


```
> imemo ls 
----------------------------------------------------------------------------------
| 1| playing tennis with Mike on Next  Tuesday                                   |
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
```

### add 
`imemo add` add what you want to write.

<img src="https://imgur.com/WlfzW7Z.gif" width="500px">


```
> imemo add "meeting at 13:30"
----------------------------------------------------------------------------------
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
memo saved!
```

## rm
please input memo's Id which you want to delete.

<img src="https://imgur.com/9y2Osxq.gif" width="500px">

```
> imemo rm 2
----------------------------------------------------------------------------------
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
memo deleted!
```


## clear
`imemo clear` delete all memo data.

<img src="https://imgur.com/6c6dJPt.gif" width="500px">

```
> imemo clear
clear 2 memo !
```



