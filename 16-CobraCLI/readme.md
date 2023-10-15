Repositorio do cobra cli = https://github.com/spf13/cobra

==================
init
==================
cobra-cli init

==================
adicionando um comando
==================
cobra-cli add ping


testeCmd.Flags().String("comando", "", "Escolha ping ou pong") => este comando NÃO permite short hands
testeCmd.Flags().StringP("comando", "c", "", "Escolha ping ou pong") => este comando permite short hands

Parametro 1: Nome do comando
Parametro 2: abreviatura do comando
Parametro 3: valor default
Parametro 4: Descrição

Utilizacao: go run main.go teste -c=pong

testeCmd.MarkFlagRequired("comando") => Obrigada digitar o comando

cobra-cli add create -p 'categoryCmd' => Comando que cria o comando create sendo filho de category

Flags:

Persistent flags: Consegue ser enxergada o comando inteiro e não somente local naquele comando.

Exemplo:
Coloquei a instrução categoryCmd.PersistentFlags().String("name", "", "Name of the category") no category, ao digitar go run main.go category, a flag --name estará lá.
Porem, se eu digitar go run main.go category create, a flag --name também estará lá.

Flags local: 
Agora, coloquei a instrução categoryCmd.Flags().String("descricao", "", "Description of the category") no category, ao digitar go run main.go category, a flag --description estará lá.
Pore, se eu digitar go run main.go category create, a flag --description NÃO estará lá porque ela é uma flag local.

comando para executar com a flag

go run main.go category --name=X
go run main.go category => imprime o valor default da flag

============

Ambos os comando imprimem exists true
go run main.go category -e true
go run main.go category -e
go run main.go category => imprime exists false
============

HOOKS

PreRun => Chamado antes do run
PostRun => Chamado depois do run
RunE => Opcao ao Run, o RunE tem a possibilidade de retornar um erro
PreRunE => Chamado antes do runE e tem a possibilidade de retornar um erro
PostRunE => Chamado depois do runE e tem a possibilidade de retornar um erro