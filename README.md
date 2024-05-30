# pos-golang-di

Repositório para as aulas de golang sobre DI (Dependency Injection)

## Anotações
Containers de injeção de dependência.


Instalar o google wire:

```bash

go install github.com/google/wire/cmd/wire@latest

```

O wire é executado na raiz do projeto e vai pegar as tags wireinject para identificar os arquivos de definição. Por convenção, o nome do arquivo 'wire.go'. 

Ao executar o comando será gerado o arquivo `wire_gen.go`.


```bash

wander@bsnote283:~/pos-golang-di$ wire
wire: github.com/wandermaia/pos-golang-di: wrote /home/wander/pos-golang-di/wire_gen.go
wander@bsnote283:~/pos-golang-di$

```

Neste caso, o main fica dividido em dois arquivos. Dessa forma, temos que executar o programa dessa forma:

```bash

go run main.go wire_gen.go

```

### Bibliotecas para Injeção de Dependências

FX

Trabalha por reflection.

Google Wire

Gera código comum das dependências.




## Referências

Fx is a dependency injection system for Go
https://github.com/uber-go/fx

Get started with Fx
https://uber-go.github.io/fx/get-started/


Wire: Automated Initialization in Go
https://github.com/google/wire