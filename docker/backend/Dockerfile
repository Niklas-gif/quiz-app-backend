#WIP
FROM golang:1.22

WORKDIR /app

RUN git clone -b docker https://github.com/Niklas-gif/quiz-app-backend.git .

RUN go mod download

RUN go build -o main .

#FROM golang:1.22
#COPY --from=build /app/main /app/main
#WORKDIR /app

EXPOSE 3030

CMD ["./main"]
