[![Xybor founder](https://img.shields.io/badge/xybor-huykingsofm-red)](https://github.com/huykingsofm)
[![Go Reference](https://pkg.go.dev/badge/github.com/xybor-x/xyselect.svg)](https://pkg.go.dev/github.com/xybor-x/xyselect)
[![GitHub Repo stars](https://img.shields.io/github/stars/xybor-x/xyselect?color=yellow)](https://github.com/xybor-x/xyselect)
[![GitHub top language](https://img.shields.io/github/languages/top/xybor-x/xyselect?color=lightblue)](https://go.dev/)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/xybor-x/xyselect)](https://go.dev/blog/go1.18)
[![GitHub release (release name instead of tag name)](https://img.shields.io/github/v/release/xybor-x/xyselect?include_prereleases)](https://github.com/xybor-x/xyselect/releases/latest)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/d767100a442f4e7c826a2d029e2a090d)](https://www.codacy.com/gh/xybor-x/xyselect/dashboard?utm_source=github.com&utm_medium=referral&utm_content=xybor-x/xyselect&utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/d767100a442f4e7c826a2d029e2a090d)](https://www.codacy.com/gh/xybor-x/xyselect/dashboard?utm_source=github.com&utm_medium=referral&utm_content=xybor-x/xyselect&utm_campaign=Badge_Coverage)
[![Go Report](https://goreportcard.com/badge/github.com/xybor-x/xyselect)](https://goreportcard.com/report/github.com/xybor-x/xyselect)

# Introduction

Package xyselect defines custom select statements.

# Features

The main object in the library is `Selector`, a custom usage of `select`
statement.

There are two types of `Selector`, which are `R` (stand for `reflect`) and `E`
(stand for exhausted).

`E` selector uses a center channel to receive all selected channels, so that it
doesn't support to wait on a sending channel. Moreover, `E` selector creates a
goroutine to wait on the center channel, and that goroutine only stops when all
selected channels closed. For this reason, you should call `Select` until get
an error of `ExhaustedError` to ensure the goroutine stopped.

`R` selector uses the built-in library, `reflect`, to customize `select`
statement. It supports to wait on both receiving and sending channels. It also
does not create any goroutine while using.

Each selector has its own advantage, while `R` selector more flexible, `E`
selector is faster.

Visit [pkg.go.dev](https://pkg.go.dev/github.com/xybor/xyplatform/xyselect) for
more details.

# Benchmark

| op name   | time per op |
| --------- | ----------- |
| RSelector | 728ns       |
| ESelector | 679ns       |

# Example

```golang
var c = make(chan int)
go func() { 
    c <- 10
    close(c)
}()

var eselector = xyselect.E()
eselector.Recv(xyselect.C(c))

var _, v, _ = eselector.Select(false)
fmt.Println(v)

// Output:
// 10
```

```golang
var rselector = xyselect.R()
var c = make(chan int)
var rc = xyselect.C(c)

go func() { c <- 10 }()
rselector.Recv(rc)
var _, v, _ = rselector.Select(false)
fmt.Println("receive", v)

rselector.Send(c, 20)
rselector.Select(false)
fmt.Println("send", <-rc)

// Output:
// receive 10
// send 20
```
