FROM golang:1.17  AS build
WORKDIR /go/src/github.com/chenhuaicong/alertmanaer-feishu-webhook/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webhook cmd/webhook/webhook.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=build /go/src/github.com/yunlzheng/alertmanaer-dingtalk-webhook/webhook .
ENTRYPOINT ["./webhook"]