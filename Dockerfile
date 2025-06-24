FROM golang:1.24.4-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

#COPY /uploads /app/uploads
RUN apk --no-cache add tzdata
ENV TZ=Asia/Jakarta

EXPOSE 3000
CMD ["./main"]