FROM golang:1.22
WORKDIR /usr/src/app
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o myapp main.go
CMD ["./myapp"]
