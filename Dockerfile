FROM golang:latest

RUN apt-get update && apt-get install -y ffmpeg

WORKDIR /app

RUN git clone https://github.com/gregidonut/VEWorkflowAutomation

WORKDIR /app/VEWorkflowAutomation

RUN go build -o app skim/cmd/web/main.go

EXPOSE 8080

CMD ["./app"]