## Youtube Thumbnails Downloader 

is a CLI tool to download YouTube thumbnails by video URLs.

![Image alt](https://github.com/SubochevaValeriya/gRPC-service-loading-youtube-thumbnails-/blob/dev/server/internal/logo/forlogo.gif)

### usage: 
```
- client/client.go [flags] URLs (you can input several URLs divided by backspaces)
- client/client.go [flags] file name.ext 
```
### usage examples:
```
go run client/client.go https://www.youtube.com/yourVideoID
go run client/client.go --async file urls.txt
```

### commands:

``` file name.ext ```

### flags:
```
--help     Show help message
--async    Flag for the program to run asynchronously
```  

### To run an app:

```
make build && make run
```

**Used:** *gRPC, MongoDB, docker-compose.*
