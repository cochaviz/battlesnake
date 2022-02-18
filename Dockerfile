FROM golang:latest

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /bin/run_snek
RUN export PORT=80

EXPOSE 80

CMD [ "env PORT=80 /bin/run_snek" ]
