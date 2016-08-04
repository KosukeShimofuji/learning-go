# GO言語の学習

## 開発環境の整備

 * 環境の情報

```
kosuke@go ~ $ lsb_release -a
No LSB modules are available.
Distributor ID: Debian
Description:    Debian GNU/Linux 8.5 (jessie)
Release:        8.5
Codename:       jessie
```

 * go言語のインストールと設定

```
$ wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.6.linux-amd64.tar.gz
$ mkdir $HOME/go
$ cat /etc/profile.d/goenv.sh
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
$ source /etc/profile.d/goenv.sh
```

## Vimの設定

 * gocode

```
$ go get github.com/nsf/gocode
```

 * .vimrc

```
filetype off
filetype plugin indent off
set runtimepath+=$GOROOT/misc/vim
filetype plugin indent on
syntax on
autocmd FileType go autocmd BufWritePre <buffer> Fmt
exe "set rtp+=".globpath($GOPATH, "src/github.com/nsf/gocode/vim")
set completeopt=menu,preview
```

## Hello World

 * hello.go

```
package main

import "fmt"

func main() {
        fmt.Println("Hello, 世界")
}
```

 * 実行

```
kosuke@go ~ $ go run hello.go
Hello, 世界
```

 * build

```
kosuke@go ~ $ go build hello.go
kosuke@go ~ $ ./hello
Hello, 世界
kosuke@go ~ $ file ./hello
./hello: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
```

バイナリはstatic linkされている。

# 参考文献

 * https://golang.org/dl/
 * https://go-tour-jp.appspot.com/welcome/1


