# 2020/04/21最新versionを取得
FROM golang:1.16.3-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src

COPY go.mod /go/src

RUN go mod download