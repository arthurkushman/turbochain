## TurboChain 
Its primary goal to be able to get up and running blockchain server with thin API layer and local storage, 
to provide data consistency + an ability to restore all data if needed by just running server again
 
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkushman/turbochain)](https://goreportcard.com/report/github.com/arthurkushman/turbochain)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://github.com/golang/gddo/blob/c782c79e0a3c3282dacdaaebeff9e6fd99cb2919/gddo-server/assets/status.svg)](https://godoc.org/github.com/arthurkushman/turbochain)

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