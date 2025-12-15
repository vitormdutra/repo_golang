create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)

insert into produtos (nome, descricao, preco, quantidade) values ('Camiseta', 'Preta', '29', '15'),('Computador', 'I5', '59', '5');

select * from produtos