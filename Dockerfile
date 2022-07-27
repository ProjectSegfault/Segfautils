FROM golang:1.18-alpine3.16

ENV SEGFAUTILITIES_PORT 6893

RUN mkdir /segfautilities 
WORKDIR /segfautilities
COPY . /segfautilities/
RUN go mod download

EXPOSE 6893

CMD ["go", "run", "main.go"]