SB
===


### Instalation

```shell
git clone https://github.com/nipeharefa/silver-parakeet
cp silver-parakeet
cp config.yml.dist config.yml

# and replace apiKey
```

### Run Api Gateway

```shell
go run *.go api-gateway
```

### Run Search Service

```shell
go run *.go movie-service
```