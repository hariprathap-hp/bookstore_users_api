#start from base image 1.16.4
FROM golang:1.16.4

#configure our work directory and gopath
ENV REPO_URL=github.com/hari/bookstore_users_api
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORK_PATH=$APP_PATH/src
COPY src $WORK_PATH
WORKDIR $WORK_PATH

RUN go build -o users_api .
EXPOSE 8080

CMD ["./users_api"]