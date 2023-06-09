# Meli-bootcamp-projetinho

Resolução das atividades:

#Exercício 1 - Gerar pacote interno

A estrutura do projeto deve ser separada e como primeiro passo gerando o pacote interno,
todas as funcionalidades que não dependem de pacotes externos devem estar no pacote
interno.
Dentro do pacote devem estar as camadas:
1. Service, deve conter a lógica da nossa aplicação.
a. O arquivo service.go deve ser criado.
b. A interface Service com todos os seus métodos deve ser gerada.
c. A estrutura de serviço que contém o repositório deve ser gerada.
d. Deve ser gerada uma função que retorne o Serviço.
e. Todos os métodos correspondentes às operações a serem executadas (GetAll,
Create, etc.) devem ser implementados.

2. Repository, você deve ter acesso à variável armazenada na memória.
a. O arquivo repository.go deve ser criado
b. A estrutura da entidade deve ser criada
c. Variáveis globais devem ser criadas para armazenar as entidades
d. A interface Repository deve ser gerada com todos os seus métodos
e. A estrutura do repositório deve ser gerada
f. Deve ser gerada uma função que retorne o Repositório
g. Todos os métodos correspondentes às operações a serem executadas (GetAll,
Store, etc.) devem ser implementados.

#Exercício 2 - Gerar pacote do server

A estrutura do projeto deve ser separada, como segundo passo deve ser gerado o pacote do
servidor onde serão adicionadas as funcionalidades do projeto que dependem de pacotes
externos e o principal do programa.

Dentro do pacote deve estar:
1. O main do programa.
a. O repositório, serviço e handler devem ser importados e injetados
b. O roteador deve ser implementado para os diferentes endpoints
2. O pacote handler com o controlador da entidade selecionada.
a. A estrutura request deve ser gerada
b. A estrutura do controller que tem o serviço como campo deve ser gerada
c. A função que retorna o controlador deve ser gerada
d. Todos os métodos correspondentes aos endpoints devem ser gerados
