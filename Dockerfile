FROM ubuntu:20.04

RUN apt update

WORKDIR /app
ENV BIRTHDAY_DB_USER=root
ENV BIRTHDAY_DB_PASS=secret

COPY ./build/gres-birthday /app/gres-birthday
COPY keys/jwt.pem app/keys/jwt.pem
COPY keys/jwt.pub app/keys/jwt.pub

ENTRYPOINT ["/app/gres-birthday"]
