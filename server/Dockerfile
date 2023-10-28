
## Build golang
FROM golang:1.21.1 AS build_go

WORKDIR /app
RUN cd /app
COPY ./ ./
RUN go mod download
RUN go build -o /stopmotion_api ./cmd/stopmotion_api/main.go

## DEPLOYMENT
FROM golang:1.21.1

WORKDIR /app
COPY --from=build_go /stopmotion_api /app/stopmotion_api
EXPOSE 5000
CMD [ "/app/stopmotion_api" ]