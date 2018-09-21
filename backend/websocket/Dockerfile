
# FROM golang

# ENV webserver_path /go/src/github.com/patientplatypus/webserver/
# ENV PATH $PATH:$webserver_path

# WORKDIR $webserver_path
# COPY . .

# RUN go get github.com/lib/pq
# RUN go get golang.org/x/crypto/bcrypt
# RUN go get github.com/gorilla/handlers
# RUN go get -u github.com/gorilla/mux
# RUN go get github.com/rs/cors
# RUN go build .

# ENTRYPOINT ./webserver

# EXPOSE 8081

FROM golang:latest

COPY . /go/src/github.com/patientplatypus/webserver/
WORKDIR /go/src/github.com/patientplatypus/webserver/

ADD . /go/src/github.com/patientplatypus/webserver/

RUN go get github.com/oxequa/realize
RUN go get github.com/jasonlvhit/gocron
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gorilla/sessions
RUN go get github.com/SparkPost/gosparkpost
RUN go get github.com/lib/pq
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/gorilla/handlers
RUN go get -u github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get github.com/gocql/gocql
RUN go get github.com/gorilla/websocket
