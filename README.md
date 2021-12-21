# Azure-go-sdk

A simple usage of azure-go-sdk to scale a SQL DB

## Setup 
```bash
cp .env.example .env

# modify .env and modify variables (requires service prinicpal)
# load env variables
./load_dotenv.sh

# build
go build -o bin

# run 
./bin

```

## References:
https://curatedsql.com/2020/03/06/managing-azure-sql-database-with-golang/
https://github.com/Azure/azure-sdk-for-go/#client-new-releases
https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/b49c4162aa1d96bc2b1b42afecbf4a21b420e568/sql/sql.go

