## Youtube Thumbnails Downloader 

is a CLI tool to download YouTube thumbnails by video URLs.

![Image alt](https://github.com/SubochevaValeriya/gRPC-service-loading-youtube-thumbnails-/blob/dev/server/internal/logo.png)

### usage: 
```
- client/client.go [flags] URLs (you can input several URLs divided by backspaces)
- client/client.go file name.ext 
```
### usage examples:
```
go run client/client.go https://www.youtube.com/yourVideoID
go run client/client.go file urls.txt
```

### commands:

``` file name.ext ```

### flags:
```
--help     Show this help message
--async    Flag for the program to run asynchronously
```  

### To run an app:

```
make build && make run
```

**Used:** *gRPC, MongoDB, docker-compose.*

// собрать всю логику вместе
запуск сервера и бд (можно в отдельных функциях)

затем по запросу ищем в бд
не находим - делаем все шаги
находим - просто грузим по ссылке

// добавить логирование где нужно

// конфиг файл отредактировать

// кли программа
флаги

-h - хэлп

--async - добавить запуск в горутинах с блоком (мьютекс?)

как принимаем URL?

списком
файлом 
по одному?

