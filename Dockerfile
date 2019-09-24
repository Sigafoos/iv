FROM golang:1.12 AS build

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY main.go .
COPY model ./model

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"'

FROM scratch
ENV PORT 4080
EXPOSE 4080
COPY --from=build /app/iv .
COPY data ./data
ENTRYPOINT ["./iv"]
