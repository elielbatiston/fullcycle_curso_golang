go build .

go build -o server
-o = output, ou seja o comando gera o arquivo com o nome server

GOOS=linux go build -o server .
GOOS = para qual sistema operacional que ele deve gerar o binário

==========
DWARF = Debugging with arbitrary record format
==========
Todo arquivo binário quando o go vai gerar, ele coloca alguns simbolos que são informações que fazem com que a gente consiga rodar ferramentas de profiling, debbuging e etc.

Se eu gero um binário com essas informações significa que o arquivo vai ser maior por outro lado, se eu quiser reduzir o tamanho do arquivo, eu posso remover essa opção porem se eu precisar fazer debbugging e profiling eu não vou conseguir

LDFLAGS

GOOS=linux go build -ldflags="-w -s" -o server .
-w remove as opções do DWARF
-s remove outros simbolos incluindo as informações de profiling

==========
Gerando imagem docker
==========
docker build -t elielbatiston/21-deploy-k8s:latest -f Dockerfile.prod .

==========
Rodando a imagem
==========
docker run --rm -p 8080:8080 elielbatiston/21-deploy-k8s:latest

==========
CGO
==========
Um dos recursos da linguagem go é conseguir importar bibliotecas da linguagem C, então vamos imaginar que eu estou fazendo o meu programa e eu preciso de uma rotina complicada de criptografia e essa rotina já está escrito em C então, você não vai desenvolver isso em go, você vai importar essa biblioteca aqui.

Porem esse seu programa tem uma dependencia agora em C e esse recurso por padrão é chamado de CGO. Esse recurso CGO já vem habilitado por padrão.

Na nossa imagem scratch eu não tenho nda disso, eu não tenho um compilador C, eu não tenho um monte de biblioteca e no nosso programa, nós não temos dependencia de nenhuma biblioteca C portanto, a gente pode desabilitar o CGO.

para desabilitar é só adicionar o CGO_ENABLED=0

GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server .

CGO_ENABLED por padrão é 1.

==========
subir a imagem para o docker
==========
docker push batistondeoliveira/21-deploy-k8s:latest

==========
Kind
==========
Instalar o Kind para gerenciar o K8S

https://kind.sigs.k8s.io/docs/user/quick-start/#installation

1-Criar um cluster k8s = kind create cluster --name=goexpert

kubectl cluster-info --context kind-goexpert

==========
Kubectl
==========
Instalar o Kubectl

https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/

1- kubectl get nodes => Saber quantos nodes estão rodando na minha maquina

kubectl apply -f k8s/deployment.yaml
kubectl get pods => ver os pods criados
kubectl apply -f k8s/service.yaml
kubectl get svc => ver todos os services
kubectl port-forward svc/serversvc 8080:8080