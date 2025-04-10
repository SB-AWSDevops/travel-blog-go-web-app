FROM golang:1.22 as base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

# Using distroless image to reduce final docker image size

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=base /app/main /app

COPY --from=base /app/static /app/static

EXPOSE 8081

CMD ["./main"]