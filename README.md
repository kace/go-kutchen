# go-kutchen

Webserver that servers up templated HTML with an embedded "chocolate cake" themed GIF from Giphy using the GiphyAPI. Written in Golang.

## Misc Golang Notes

### install go
sudo apt-get install golang-go

### websocket info
https://github.com/golang/go/wiki/GOPATH
http://www.gorillatoolkit.org/pkg/websocket
https://github.com/gorilla/websocket/blob/master/examples/echo/client.go
https://github.com/crsmithdev/goenv

### Build for windows x64
GOOS=windows GOARCH=amd64 go build -o <output_exe> <source_go>
