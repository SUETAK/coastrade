# 2020/04/21最新versionを取得
FROM golang:1.16.3-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git && apk add --no-cache gcc && apk add --no-cache musl-dev
# appディレクトリの作成
RUN mkdir /go/coastrade
# ワーキングディレクトリの設定
WORKDIR /go/coastrade

COPY go.mod /go/coastrade
COPY go.sum /go/coastrade

RUN go mod download