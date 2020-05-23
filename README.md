### Lightweight blockchain server application 
Its primary goal to be able to get up and running blockchain server with thin API layer and local storage, to provide data consistency + an ability to restore all data if needed by just running server again 

#### Configuration
To configure port and storage file names edit .env file
```
# set any port for block-chain server
ADDR=8080
# set any db file name for block-chain storage
BLOCK_CHAIN_DB=turbochain
```
make sure go environment set up correctly and run: 
```
go build -o turbochain 

./turbochain
```
or 
```
go run server.go
```