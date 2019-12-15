# Izzi
Izzi é um serviço que expõe apis para a fácil contratação de um seguro de vida por autonomos.
## Development

### Fazer Download do Wires
O Wire é responsável pela injeção de dependências do projeto. Faça o download do mesmo:
```shell
go get -u github.com/google/wire/cmd/wire
```

### Injeção de dependências
Para realizar a injeção de dependências, execute o comando abaixo no diretório raiz do projeto:
```shell
wire
```
A execução do comando anterior produzirá um arquivo chamado `wire_gen.go`.

---

## Execução

### Via Aplicação
Após realizar a injeção das dependências, execute o comando abaixo na raiz do projeto:
```shell
go run ./
```

### Via Docker
- Na raiz do projeto existe um arquivo `Dockerfile`. Gere uma imagem a partir do mesmo:
    ```shell
    docker build -t <image-name> .
    ```
- Ao executar a imagem, fique atento as variáveis de ambiente que devem ser informadas:
    ```shell
    docker run -it --rm \
        --name=izi \
        -e "ENV=production" \
        -e "LOG_LEVEL=INFO" \
        -e "MONGERAL_AEGON_HOST=https://gateway.gr1d.io/production/mongeralaegon/v1/" \
        -e "COMPLINE_HOST=https://gateway.gr1d.io/production/compline/signature/v1/" \
        -e "BIG_ID_HOST=https://gateway.gr1d.io/production/bigdata/bigid/ocr/v1/" \
        -e "SERVER_PORT=8107"
        <image-name>
    ```

### Via Docker (sh)
- Gere e execute a imagem através da execução dos arquivos `generate-image.sh` e `run-docker.sh` respectivamente.

Obs: Nas execuções via docker, caso uma ou mais variáveis de ambiente não sejam informadas, os valores padrões virão do arquivo `.env`, que está na raiz do projeto.

## Testes Unitários
- Para rodar os testes do diretório atual e subdiretórios gerando o resultado em um arquivo coverage.out, execute: 
    ```shell
    go test -coverprofile=coverage.out  ./...
    ```
- Para analisar o arquivo coverage.out gerado, execute:
    ```shell
    go tool cover -func=coverage.out
    ```