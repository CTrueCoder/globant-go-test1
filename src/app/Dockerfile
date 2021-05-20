FROM golang:alpine

WORKDIR /src

COPY conf conf/
COPY controllers controllers/
COPY dbapi dbapi/
#COPY models models/
COPY routers routers/
COPY static static/
COPY tests tests/
COPY views views/
COPY main.go .
COPY go.mod .

#RUN go mod tidy
RUN go mod download
#RUN go build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o globantapp

EXPOSE 8080
CMD ["./globantapp"]