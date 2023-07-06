template no repositorio do git => http://github.com/golang-standards/project-layout

======================
Estrutura do projeto
======================
* internal => Tudo que esta relacionado ao projeto e que nao podera ser compartilhado com outros projetos

* pkg => Oposto do internal, aqui vem tudo que podera ser compartilhado com outros projetos

* cmd => Aqui vem o seu programa (main)

* configs => Aqui vem as configurações padrao (templates) para o projeto funcionar. Para algumas pessoas, essa pasta
também pode conter os arquivos de configuração para fazer o boot da aplicação

* test => Arquivos adicionais que vao nos ajudar a realizar os testes da nossa aplicação como:
  - Documentacao de testes
  - Exemplo de testes
  - stubs de testes
  - Arquivos Http para você poder fazer testes end2end
  - Declarações de uma biblioteca do postman para você conseguir rodar
OBS: Nao necessariamente esses arquivos devem ser arquivos .go mas podem ser scripts por exemplo

* api => Onde fica a documentacao da nossa api

======================
Web Frameworks
======================
Focado mais na parte da Web

* golang echo (http://echo.labstack.com)
* golang fiber (http://gofiber.io) => Parecido com express do node
* gin golang (http://gin-gonic.com)

======================
Frameworks
======================
Vai alem de Http

* buffalo golang (http://gobuffalo.io) => tipo laravel
* iris golang (http://github.com/kataras/iris)

======================
Roteadores
======================
* gorilla golang (http://gorillatoolkit.org)
* golang chi (http://github.com/go-chi/chi) => mais leve

======================
Como funciona o Middleware
======================
Request -> Middleware(usa os dados, faz alguma coisa. Continua)|outro Middleware|outro Middleware -> Handler -> Response

======================
Documentacao
======================
swaggo / swag (http://github.com/swaggo/swag)

* swag init -g cmd/server/main.go