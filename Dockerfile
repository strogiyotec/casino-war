FROM golang:1.16-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
# Disable CGO to link all libs statically and create an executable called 'main'
# that we will copy into next stage to decrease size of image
RUN CGO_ENABLED=0 GOOS=linux go build -o main 


#Next build stage, just use smallest alpine
FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./main"]

