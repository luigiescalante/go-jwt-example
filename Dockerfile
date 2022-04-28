FROM 220361521127.dkr.ecr.us-east-1.amazonaws.com/go:1.18.1
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY ./src/ .
COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download && \
    go build -o main .
WORKDIR /dist
RUN cp /build/.env .env
RUN cp -R /build/files files/
RUN cp /build/main .
EXPOSE 80
CMD ["/dist/main"]