FROM golang:1.10-alpine
COPY . .
CMD sh "script.sh"