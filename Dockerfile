FROM ubuntu:latest

COPY ./main .
COPY ./edges.json .

#ARG GODEBUG=madvdontneed=1
#ENV GODEBUG madvdontneed=1

CMD ["/main"]
