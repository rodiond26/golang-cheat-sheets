FROM golang:1.20.1-alpine3.16
# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/go-mysql
COPY . .
RUN go build -o main .
# Start the application.
CMD ["./main"]
