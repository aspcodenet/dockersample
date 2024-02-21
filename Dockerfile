FROM golang:alpine AS builder
# i denna ligger ju tooling, kompilatorn, 

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v
# dom lagras ju i imagen

RUN go build -o /app/cmd/site

FROM scratch
#den minsta 
COPY --from=builder /app/cmd/site /app/cmd/site

EXPOSE 8080
ENTRYPOINT ["/app/cmd/site"]

