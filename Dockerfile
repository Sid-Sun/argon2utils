FROM golang:buster as build

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go main.go
RUN go build -o argon2util .

FROM alpine
WORKDIR /root/app
COPY --from=build /build/argon2util .

CMD [ "/root/app/argon2util" ]
