build:
	go build -ldflags="-w -s" -o gohttpd.exe 
	upx gohttpd.exe