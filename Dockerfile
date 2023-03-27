FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .



FROM gcr.io/distroless/static-debian11
COPY --from=build /app/main .
EXPOSE 80
CMD ["/main"]
