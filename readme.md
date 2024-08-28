# Sync-Datasus

### Descrição:

O `sync-datasus` é uma aplicação desenvolvida em Go para automatizar o processo de coleta e sincronização de dados da base FTP do DATASUS com tabelas em um banco de dados PostgreSQL. Este projeto facilita a obtenção de dados de saúde pública, garantindo que as informações estejam sempre atualizadas e prontas para análise.

### Funcionalidades:

* Conexão e autenticação automática com o servidor FTP do DATASUS.
* Download de arquivos de dados diretamente do FTP.
* Processamento e transformação dos dados.
* Inserção dos dados processados em tabelas específicas do PostgreSQL.
* Monitoramento e logging das operações para garantir a integridade dos dados.

### Requisitos:

* [Go](https://golang.org/) 1.19 ou superior.
* [PostgreSQL](https://www.postgresql.org/) 13 ou superior.
* Conexão com a internet para acessar o FTP do DATASUS.

### Instalação:

1. Clone o repositório:

```shell
git clone https://github.com/CriticalNoob02/sync-datasus.git
cd sync-datasus
```

2. Instale as dependências:

```shell
go mod tidy
```

3. Configure as variáveis de ambiente:

```shell
FTP_HOST="ftp.datasus.gov.br"
POSTGRES_USER="seu_usuario_postgres"
POSTGRES_PASSWORD="sua_senha_postgres"
POSTGRES_DB="seu_banco_de_dados"
BATCH_LIMIT="Limite de arquivos para cada worker"
NUM_WORKS="Numero de Workers rodando em paralelo"
MODULE_LIMIT_DATE="Data filtragem dos arquivos"
```

4. Configure o mapper com base nos dados que deseja processar:

```go
	case "RAAS":
		os.Setenv("MODULE_TYPE", "PS")
		os.Setenv("MODULE_REMOTE_DIR", "dissemin/publicos/SIASUS/200801_/Dados")
		os.Setenv("MODULE_TABLE_NAME", "tb_fat_importacoes_raas")
		os.Setenv("MODULE_TABLE_SCHEMA", "public")
		os.Setenv("QUERY_LIMIT", "1000")
	}
```

5. Compile o projeto:

```shell
go build -o build/myapp ./cmd
```

### Uso:

Execute o binário gerado para iniciar o processo de sincronização:

`./sync-datasus/build`

O aplicativo fará o download dos arquivos de dados, processará as informações e as inserirá nas tabelas do PostgreSQL conforme configurado.

### Contribuição:

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.


> **Observacoes técnicas:**
>
> 1. No modelo que desenvolvi, as colunas da sua tabela no banco de dados devem ter os mesmos nomes das colunas nos arquivos do DATASUS. Caso você não queira importar uma coluna específica do DATASUS, basta que ela não exista na sua tabela. O serviço mapeia automaticamente as tabelas do banco e coleta apenas as informações correspondentes no DATASUS.
> 2. É possível filtrar os dados para um Estado específico aplicando uma filtragem no nome do arquivo. No entanto, será necessário implementar um loop na task Reader para aplicar essa filtragem corretamente.




<p align="center">
  <a href="https://github.com/radarsaude/api-ia">
    <img src="https://images2.alphacoders.com/133/1335141.png" width="100%" height="200" alt="Banner">
  </a>
<p/>
