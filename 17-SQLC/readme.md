============================
1- GOLANG MIGRATE
============================
https://github.com/golang-migrate/migrate

Instalação: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

Comandos:
migrate create -ext=sql -dir=sql/migrations -seq init
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down


============================
2 - SQLX
============================
https://github.com/jmoiron/sqlx

Deixa menos verboso, ele facilita o uso do scan por exemplo

============================
3- SQLC
============================
Documentação: https://sqlc.dev
github: https://github.com/kyleconroy/sqlc/blob/main/docs/tutorials/getting-started-postgresql.md

Comandos:
sqlc generate