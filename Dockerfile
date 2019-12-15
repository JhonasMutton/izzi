##Build Image
FROM golang:1.13-stretch AS builder
COPY . /izzi
WORKDIR /izzi
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application

#Run Image
FROM scratch
COPY --from=builder /izzi/bin/application application
COPY .env .env
EXPOSE 8107
ENTRYPOINT ["./application"]