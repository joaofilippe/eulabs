# Eulabs

![Badge License](https://img.shields.io/badge/LICENSE-MIT-green)

API desenvolvida para atender case do processo seletivo da EuLabs.

Tecnologias utilizadas para o desenvolvimento:

- Golang;
- Echo Framework;
- MySQL;
- GORM;
- GoDotenv;
- Shortuuid;

## Utilização da API:

Link para a documentação no POSTMAN:
https://documenter.getpostman.com/view/16820128/UVsPPQBa

## Funcionalidades:

1. Criar um novo produto – Método POST;
2. Requisitar as informações de um produto pela ID – Método GET;
3. Editar as informações de um produto;
4. Deletar um produto;

## Estrutura do Projeto:

```
.
├── business                   
│   ├── products.go                   
├── controllers             
│   ├── products.go                   
├── database               
│   ├── db.go                   
├── migrations
│   ├── migrations.go 
├── models                   
│   ├── product.go 
├── routes                   
│   ├── routes.go 
├── services
│   ├── hashManager.go 
├── LICENSE
├── .gitignore
├── .gitattributes
├── go.mod
├── go.sum
├── main.go
└── README.md
```

### Autor

João Filippe Rossi Rodrigues

[![Github Badge](https://img.shields.io/badge/-Github-000?style=flat-square&logo=Github&logoColor=white&link=LINK_GIT)](https://github.com/joaofilippe)

[![Linkedin Badge](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white&link=LINK_LINKEDIN)](https://www.linkedin.com/in/joaofilippe/)
