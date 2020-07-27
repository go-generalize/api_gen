# api_gen

## Introduction

### 前提条件

[struct2ts](https://github.com/OneOfOne/struct2ts) のインストールが必須。
※ client_generateで必要。

```shell
go get -u -v github.com/OneOfOne/struct2ts/...
```



### インストール

#### リリースよりダウンロード
[Release](https://github.com/go-generalize/api_gen/releases/)

#### 自分の環境でビルドする場合
```shell script
go get -u github.com/go-generalize/api_gen
go install github.com/go-generalize/api_gen/server_generator
go install github.com/go-generalize/api_gen/client_generator
```