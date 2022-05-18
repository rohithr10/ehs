FROM golang:latest as build

WORKDIR D:/srm_project/BE/ehs

COPY . .

WORKDIR D:/srm_project/BE/ehs/cmd/db

RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /db

FROM alpine:3.7
COPY --from=build /db /db

EXPOSE 8080
ENTRYPOINT ["/db"]