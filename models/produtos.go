package models

import "GWStore/db"

type Produtos struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produtos {
	db := db.ConectaBancoDados()
	selectTodosProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produtos{}
	produtos := []Produtos{}

	for selectTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos

}
