# ビルド用ステージ
FROM golang:1.23.3-bullseye as builder

WORKDIR /go-bookmark/

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
# -o : 実行ファイルの生成場所指定
# -ldflags '-s -w': 実行ファイルにアプリの動作に関係ないものを入れないためのオプションを付与
# -tags timetzdata: タイムゾーンデータ組み込み
RUN go build \
    -o /go/bin/go-bookmark \
    -ldflags '-s -w' \
    -tags timetzdata

# 実行用ステージ
FROM scratch as runner

COPY --from=builder /go/bin/go-bookmark /app/go-bookmark

EXPOSE 8080

ENV TZ=Asia/Tokyo

ENTRYPOINT ["/app/go-bookmark"]
CMD ["web"]
