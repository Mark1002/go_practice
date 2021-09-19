
## Build
FROM golang:1.17-buster AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /ginAPI

## Deploy
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /ginAPI /ginAPI

EXPOSE 80
ENTRYPOINT [ "/ginAPI" ]
