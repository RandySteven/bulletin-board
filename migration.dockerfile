FROM golang:alpine

ENV TZ=Asia/Jakarta
WORKDIR /app

COPY go.mod go.sum ./
RUN rm -rf vendor/* bin/*

RUN go clean -mod=mod
RUN go mod tidy
RUN go mod download && go mod verify
RUN go mod vendor

COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/bulletin-migration cmd/task_mission/migration/main.go

EXPOSE 8888
ENTRYPOINT ["./bin/bulletin-migration"]
CMD ["-config=/app/files/yml/configs/task.docker.yml"]