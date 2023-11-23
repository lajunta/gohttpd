win:
	GOARCH=amd64 GOOS=windows go build -ldflags="-w -s" -o bin/gohttpd.exe 
	upx bin/gohttpd.exe
linux:
	GOARCH=amd64 GOOS=linux go build -ldflags="-w -s" -o bin/gohttpd
	upx bin/gohttpd

all: win linux