FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o program .

# Generate init.sql dari migrations
RUN rm -f init.sql && \
    for file in $(find migrations -name "*.up.sql" | sort); do \
        echo "-- $file" >> init.sql; \
        cat "$file" >> init.sql; \
        echo "" >> init.sql; \
    done

# =================================================

FROM postgres:alpine

WORKDIR /app

ENV POSTGRES_USER=program
ENV POSTGRES_PASSWORD=program
ENV POSTGRES_DB=program

COPY --from=builder /app/program .
COPY --from=builder /app/init.sql /docker-entrypoint-initdb.d/
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

CMD ["./entrypoint.sh"]