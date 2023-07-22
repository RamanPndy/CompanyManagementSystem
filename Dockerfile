# Latest golang image on apline linux
FROM golang:1.18-alpine AS builder

RUN apk add --no-cache ca-certificates && update-ca-certificates

# Work directory
WORKDIR /go/src/company-builder-app

ADD ./ .

# Copying all the files
COPY /cmd/api/main.go .
RUN go mod tidy
RUN GOOS=linux go build -o "company-builder-app"

FROM golang:1.18-alpine
WORKDIR /bin/company-builder-app

COPY --from=builder /go/src/company-builder-app/company-builder-app .
RUN mkdir config
COPY --from=builder /go/src/company-builder-app/config/tier ./config/tier

CMD ["./company-builder-app"]
EXPOSE 80