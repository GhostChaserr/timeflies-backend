FROM golang:1.15.7-buster

ENV GO111MODULE=on
ENV GOFLAGS=-mod=readonly

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ENV PORT=5001

COPY . .
EXPOSE 5001

CMD ["bash", "start.sh"]