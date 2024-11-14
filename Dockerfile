FROM golang:alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go


# Run stage

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
EXPOSE 8811

CMD [ "/app/main" ]