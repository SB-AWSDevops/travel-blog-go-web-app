FROM golang:1.22 as base

WORKDIR /app

COPY go.mod . 
RUN go mod download

COPY . .

RUN go build -o main .

# Final distroless image
FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=base /app/main .

EXPOSE 8081

CMD ["./main"]
