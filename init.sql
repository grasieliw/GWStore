CREATE TABLE IF NOT EXISTS produtos (
  id SERIAL PRIMARY KEY,
  nome VARCHAR,
  descricao VARCHAR,
  preco DECIMAL,
  quantidade INTEGER
);
