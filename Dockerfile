FROM golang:1.10

WORKDIR /go/src/football-squares

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o app ." -command="./app"
