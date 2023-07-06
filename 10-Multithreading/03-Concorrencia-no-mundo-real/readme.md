==============================
Ferramenta de teste de stress
==============================
ab => apache benchmark

==============================
Como usar
==============================
ab -n X -c Y Z

Onde:
X é o número de vezes que você quer rodar.
Y é a quantidade de pessoas acessando simultaneamente.
Z é a url a ser testada

Exemplo:
ab -n 10000 -c 100 http://localhost:3000/

go run -race main.go => executa e diz para o go verificar se estamos com alguma condição em race condition

