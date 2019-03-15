APP=ktwig
GO-FLAGS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build-linux:
	${GO-FLAGS} go build -o ${APP}
clean:
	rm -rf ${APP}
