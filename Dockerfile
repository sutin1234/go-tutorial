FROM golang:1.20.5-alpine
# RUN apk update && apk add gcc git 

WORKDIR /myApp
COPY . .
RUN go build ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /myApp/app .
CMD ["./app"]