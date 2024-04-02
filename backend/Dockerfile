FROM golang:alpine

WORKDIR /app

# COPY entrypoint.sh /app/entrypoint.sh
# RUN chmod +x /app/entrypoint.sh

# COPY go.sum ./
# RUN go mod download

# COPY *.go ./
# RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.40.4


# ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["air"]