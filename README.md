# go-gorilla-sqlboiler-mysql

## HTTP Request List
```
$ curl -X GET http://localhost:8082/user/{id}
$ curl -X GET http://localhost:8082/user/list
$ curl -X GET http://localhost:8082/review/list
$ curl -X GET http://localhost:8082/review/{id}
```

## Article
 - gorilla  
    - https://github.com/gorilla/mux  
    - https://qiita.com/gold-kou/items/99507d33b8f8ddd96e3a  

 - sqlboiler  
    - https://github.com/volatiletech/sqlboiler#configuration  
    - https://zenn.dev/sagae/articles/c6b2e460201d31  
    - https://note.crohaco.net/2020/golang-sqlboiler/  
 - grpc  
    - https://grpc.io/docs/languages/go/quickstart/  
    - https://github.com/grpc/grpc-go  
    - https://dev.classmethod.jp/articles/golang-grpc-sample-project/  

## sqlboiler
    1. path  
    ```
    $ export GOPATH=$HOME/go;
    $ export PATH=$PATH:$GOPATH/bin;
    ```
    2. command  
    ```
    $ sqlboiler mysql -c sqlboiler/sqlboiler.toml -o pkg/models --no-tests
    ```
    3. check sqlboiler verion. if exist error files. 
    ```
    sqlboiler --version
    ```

