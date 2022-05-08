FROM golang:1.18 AS build
WORKDIR /Users/xj/go/src/learning-k8s/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver learning-k8s/cmd/learning-k8s-02

FROM busybox
COPY --from=build /Users/xj/go/src/learning-k8s/httpserver /Users/xj/go/src/learning-k8s/httpserver
EXPOSE 8360
ENV ENV local
WORKDIR /Users/xj/go/src/learning-k8s/
ENTRYPOINT ["./httpserver"]
