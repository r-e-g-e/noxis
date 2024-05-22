FROM golang:1.22.3-alpine3.19

WORKDIR /app

COPY . .

ENV PORT=8080
EXPOSE 8080

RUN go build

CMD ./noxis