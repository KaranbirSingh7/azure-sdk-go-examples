# Azure-go-sdk

A simple usage of azure-go-sdk to scale a SQL DB

## Setup 
```bash
cp .env.tpl .env

# modify .env and modify variables (requires service prinicpal)
# load env variables
./load_dotenv.sh

# build
go build -o bin

# run 
./bin

## SAMPLE OUTPUT
# ❯ go run main.go
# 2021/12/20 21:11:54 InProgress
# 2021/12/20 21:12:09 DB scale operation to S4 Succeeded
```

## References:
- https://curatedsql.com/2020/03/06/managing-azure-sql-database-with-golang/

- https://github.com/Azure/azure-sdk-for-go/#client-new-releases

- https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/b49c4162aa1d96bc2b1b42afecbf4a21b420e568/sql/sql.go

