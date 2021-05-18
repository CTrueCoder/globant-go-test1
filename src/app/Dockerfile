FROM golang:alpine

WORKDIR /src

COPY conf conf/
COPY controllers controllers/
COPY dbapi dbapi/
COPY models models/
COPY routers routers/
COPY static static/
COPY tests tests/
COPY views views/
COPY main.go .
COPY go.mod .

#RUN go mod donwload
RUN go mod tidy
RUN go build

EXPOSE 8080
CMD ["./globantapp"]