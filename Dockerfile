# Build the backend with the golang image
FROM golang:1.13 AS backend-builder

COPY backend /backend
WORKDIR /backend

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app

# Build the frontend with the node image
FROM node:13.2 AS frontend-builder

COPY frontend /frontend
WORKDIR /frontend

RUN npm install
RUN npm run build


# Build the final image with the results from the builders
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=backend-builder /backend/app .
COPY --from=frontend-builder /frontend/public public

EXPOSE 8080
CMD ["./app"]
