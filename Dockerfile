FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY blockchain-go .

CMD ["./blockchain-go"]
EXPOSE 8080
