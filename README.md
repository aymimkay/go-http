## Usage
```
go-http.exe -h
Usage of go-http.exe:
  -address string
        Listening address (default "0.0.0.0")
  -port string
        Listening port (default "80")
  -status int
        Returned HTTP status code (default 200)
```
## Examples
```
$ go-http.exe page.html
$ go-http.exe -port 80 -status 400 "<center><h1>400 Bad Request</h1></center>"
$ go-http.exe .\www\
```
