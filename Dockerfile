FROM ubuntu:latest

WORKDIR /app

COPY main .
COPY web ./web

EXPOSE 7540

CMD ["./main"]