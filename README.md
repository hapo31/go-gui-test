# setup
上から順にやる

```
$ sudo apt install gcc libosmesa6-dev libxcursor-dev libx11-dev xorg-dev
```

```
$ go mod tidy
```

# build & run
```
$ go run main.go
```

# tips
`LIBGL_ALWAYS_INDIRECT=1` が環境変数に設定されていると動かない。  
これは昔のバージョンから入っているWSL2の `~/.bashrc` に設定されていることがある（WSLgがなかった頃のworkaroundらしい）  
されている場合は
```
$ unset LIBGL_ALWAYS_INDIRECT
```
として、 `~/.bashrc` 内の `LIBGL_ALWAYS_INDIRECT=1` を消してしまう。
