FROM golang:1.22-alpine AS build

WORKDIR /app

COPY ./controllers/ /app/controllers/
COPY ./database/ /app/database/
COPY ./models/ /app/models/
COPY ./routes/ /app/routes/
COPY ./main.go /app/main.go/
COPY ./go.mod /app/go.mod/
COPY ./go.sum /app/go.sum/

RUN go build main.go

FROM alpine:latest AS production

WORKDIR /app

ENV PORT 8080
ENV DB_HOST postgress
ENV DB_USER root
ENV DB_PASSWORD root
ENV DB_NAME 5432

EXPOSE 8080


COPY ./assets/ /app/assets/
COPY ./templates/ /app/templates/

COPY --from=build /app/main /app/main

CMD ["./main"]