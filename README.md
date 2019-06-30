# imemo
<img  align="right" src="https://imgur.com/qxEBvis.png" width="120px">

[![GoDoc](https://godoc.org/github.com/kcwebapply/imemo?status.svg)](https://godoc.org/github.com/kcwebapply/imemo)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/imemo)
[](https://github.com/gin-gonic/gin/releases)
[![Release](https://img.shields.io/github/release/kcwebapply/imemo.svg?style=flat-square)](https://github.com/kcwebapply/iemo/release)

`imemo` is very simple `memo` tool for terminal.


- [ls (show memo list)](#ls)
- [add (add memo)](#add)
- [rm (delete memo)](#rm)
- [clear (delete all memo)](#clear)

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

default
------------------------------------------------------------------------------------------
| 1| Meething with Tom at 14:00                                                          |
------------------------------------------------------------------------------------------

business
------------------------------------------------------------------------------------------
| 2| Reserve a flight                                                                    |
------------------------------------------------------------------------------------------

```

All memo are listed  grouped by `category` .

You can specify category when you add new memo.



### add
`imemo add` add what you want to write.

<img src="https://imgur.com/WlfzW7Z.gif" width="500px">



```
> imemo add "Meething with Tom at 14:00"
>> Meething with Tom at 14:00
memo saved!

```

you san save memo with `category` .

```
> imemo add "Reserve a flight" business
>> Reserve a flight
memo saved!
```

In this case, 3rd Argument `"business"` is category.

Then, you can check that memo within `"business"` category .

(sample is [here](#ls).)



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
