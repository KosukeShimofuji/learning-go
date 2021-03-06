# GO言語の学習

## 開発環境の整備

 * deploy

```
$ git clone git@github.com:KosukeShimofuji/learning-go.git 
$ cd ansible
$ ansible-playbook -i development site.yml
```

 * 環境の情報

```
kosuke@go ~ $ lsb_release -a
No LSB modules are available.
Distributor ID: Debian
Description:    Debian GNU/Linux 8.6 (jessie)
Release:        8.6
Codename:       jessie
```

 * go言語のインストールと設定

```
$ go version
go version go1.7.1 linux/amd64
```

## vim-go-ide

go言語用のvim設定を作成する

```
git clone git@github.com:farazdagi/vim-go-ide.git ~/.vim_go_runtime
sh ~/.vim_go_runtime/bin/install
mkdir -p ~/.vim/autoload ~/.vim/bundle && curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
alias govim='vim -u ~/.vimrc.go'
```

go言語用のtagsを用意する

```
go get -u github.com/jstemmer/gotags
gotags -R * > tags
```

```
:GoInstallBinaries
```

```
cat /home/kosuke/.vim_go_runtime/custom_config.vim
nnoremap <C-[> :pop<CR>
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

## ディレクトリ構成

```
mkdir -p go/src/github.com/KosukeShimfouji/
cd go/src/github.com/KosukeShimfouji/
git clone git@github.com:KosukeShimofuji/learning-go.git
```

# 参考文献

 * https://golang.org/dl/
 * https://go-tour-jp.appspot.com/welcome/1
 * https://github.com/farazdagi/vim-go-ide
 * https://github.com/STNS/libnss_stns
 * http://qiita.com/yugui/items/e71d3d0b3d654a110188
 * http://qiita.com/yugui/items/cc490d080e0297251090
 * https://speakerdeck.com/pyama86/goyan-yu-deshi-zhuang-surulinuxnameserver

