### INSTALAÇÃO

1- Instalar o protoc

https://grpc.io/docs/protoc-installation/

```
sudo apt update
sudo apt upgrade
```

```
sudo apt-get install autoconf automake libtool curl make g++ unzip
```

```
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.21.9/protobuf-3.21.9.tar.gz
```

```
tar -xzvf protobuf-3.21.9.tar.gz
```

```
./autogen.sh
./configure
make
```

```
sudo make install
```

```
sudo ldconfig
```

```
protoc --version
```

2- Instalar os plugins do go

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

3- Atualizar o path

export PATH="$PATH:$(go env GOPATH)/bin"

4- Gerar arquivos proto

protoc --go_out=. --go-grpc_out=. proto/course_category.proto

5- Client Evans 

https://github.com/ktr0731/evans

a. Instalar: go install github.com/ktr0731/evans@latest
b. evans -r repl
c. Selecionar o serviço: service <nome_servico>, exemplo: service CategoryService

Comandos: 
package pb
service CategoryService
show package
show service
call <nome_servico>, exemplo call CreateCategory

6- Entrar no sqlite3

sqlite3 db.sqlite

7- Criar tabela
create table categories (id string, name string, description string);