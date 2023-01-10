======================
Configurar variavel GOPRIVATE
======================
go env | grep PRIVATE

export GOPRIVATE=github.com/batistondeoliveira/fcutil-secret

colocar entre virgulas se tiver mais de um
exemplo:
export GOPRIVATE=github.com/batistondeoliveira/fcutil-secret,github.com/batistondeoliveira/xpto

======================
Baixar github via http
======================
vim ~/.netrc => Se não tiver vc cria

Adicionar as linhas abaixo no arquivo

machine github.com
login batistondeoliveira
password <token_github>

======================
Baixar github via ssh
======================
vim .git/config => edita as configurações somente do projeto que você está trabalhando
vim ~/.gitconfig => edita as configurações gerais do github (serve para qq projeto)

Adicionar as linhas abaixo no arquivo

[url "ssh://git@github.com/"]
  instedOf=https://github.com/

======================
Vendor
======================
Cria uma copia de todas as suas dependencias para caso a dependencia original seja apagada, você tenha a copia

Execute:
go mod vendor