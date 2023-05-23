# EulabsApi

<!---Esses são exemplos. Veja https://shields.io para outras pessoas ou para personalizar este conjunto de escudos. Você pode querer incluir dependências, status do projeto e informações de licença aqui--->

![GitHub repo size](https://img.shields.io/github/repo-size/iuricode/README-template?style=for-the-badge)
![GitHub language count](https://img.shields.io/github/languages/count/iuricode/README-template?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/iuricode/README-template?style=for-the-badge)
![Bitbucket open issues](https://img.shields.io/bitbucket/issues/iuricode/README-template?style=for-the-badge)
![Bitbucket open pull requests](https://img.shields.io/bitbucket/pr-raw/iuricode/README-template?style=for-the-badge)

<img src="https://i.imgur.com/mRy3JVj.png" alt="exemplo imagem">

>Resumo:
Este projeto foi desenvolvido a fim de solucionar um desafio proposto pela empresa eulabs. Foi feito uma api em golang usando o echo framework, contendo: testes, documentação, docker 

## 💻 Pré-requisitos

Antes de começar, verifique se você atendeu aos seguintes requisitos:
<!---Estes são apenas requisitos de exemplo. Adicionar, duplicar ou remover conforme necessário--->
* Você instalou o  golang?.
* Você instalou o docker-compose?.
* Você leu o README do projeto?.

## ☕ Usando eulabsApi

Para usar EulabsApi, siga estas etapas:


- Instale as dependências do projeto utilizando o comando “go mod tidy”

- rode o comando “go get -u swag” para baixar as dependências do swagger 

- rode o comando “docker-compose up -d” para subir o banco de dados utilizado no docker

- execute a função main.go do path “./cmd/migrations/main.go”, para subir as migrações do banco de dados 

- Execute a função main.go do path “./cmd/api/main.go”, para iniciar a api

- Execute o comando “go test —“ para visualizar todos os testes unitários e de integração 




Para contribuir com EulabsApi, siga estas etapas:

1. Bifurque este repositório.
2. Crie um branch: `git checkout -b <nome_branch>`.
3. Faça suas alterações e confirme-as: `git commit -m '<mensagem_commit>'`
4. Envie para o branch original: `git push origin <nome_do_projeto> / <local>`
5. Crie a solicitação de pull.

Como alternativa, consulte a documentação do GitHub em [como criar uma solicitação pull](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).
