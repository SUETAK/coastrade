# 2020/04/21最新versionを取得
FROM golang:1.16.3-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src
# ワーキングディレクトリの設定
WORKDIR /go/src
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src