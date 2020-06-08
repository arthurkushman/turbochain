## TurboChain 
Its primary goal to be able to get up and running blockchain server with thin API layer and local storage, 
to provide data consistency + an ability to restore all data if needed by just running server again
 
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkushman/turbochain)](https://goreportcard.com/report/github.com/arthurkushman/turbochain)
[![codecov](https://codecov.io/gh/arthurkushman/turbochain/branch/master/graph/badge.svg)](https://codecov.io/gh/arthurkushman/turbochain)
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
#### Using API to get/put blocks from/to chain
To put a new block to a chain:
```http request
curl --data '{"bpm":82, "title": "abc", "description": "abc abc abc"}' http://localhost:8080/
```

```
{
  "index": 1,
  "timestamp": "2020-05-24 10:41:52.429649 +0300 MSK m=+26.655275710",
  "hash": "8f3e317922b6accb60f7d51bb58aa1a1b12fd87af082e26021736a38881c5fb1",
  "prev_hash": "d4a1a47f15d55e81304e22fc52db6090dd5cae5e8d544911393737d38d193de4",
  "message": {
    "bpm": 82,
    "title": "abc",
    "description": "abc abc abc"
  }
}
```
Getting blocks by an index and hash:
```http request
curl http://localhost:8080/one\?index\=5\&hash\=32ddce756b09c74cfcd24a4a3c7b7cb3aea2cd131f7a11f29d8043ae12ed424a
```

```
{
  "index": 5,
  "timestamp": "2020-05-24 10:43:57.869286 +0300 MSK m=+152.093356109",
  "hash": "32ddce756b09c74cfcd24a4a3c7b7cb3aea2cd131f7a11f29d8043ae12ed424a",
  "prev_hash": "f7740abe30fe1540d23ed288fc945d4ccc172c09bdc3de732d86441e36a37074",
  "message": {
    "bpm": 23,
    "title": "abc 5",
    "description": "abc abc abc"
  }
}
```
To get last N blocks:
```http request
curl http://localhost:8080/last?limit=3
```
Restrictions: 0 <= limit <= 1000 (default is 1000)
```
[
  {
    "index": 5,
    "timestamp": "2020-05-24 10:43:57.869286 +0300 MSK m=+152.093356109",
    "hash": "32ddce756b09c74cfcd24a4a3c7b7cb3aea2cd131f7a11f29d8043ae12ed424a",
    "prev_hash": "f7740abe30fe1540d23ed288fc945d4ccc172c09bdc3de732d86441e36a37074",
    "message": {
      "bpm": 23,
      "title": "abc 5",
      "description": "abc abc abc"
    }
  },
  {
    "index": 4,
    "timestamp": "2020-05-24 10:43:44.147761 +0300 MSK m=+138.372001638",
    "hash": "f7740abe30fe1540d23ed288fc945d4ccc172c09bdc3de732d86441e36a37074",
    "prev_hash": "5de3b32b6b52a3f73be63e6955e4471d884123e099b80484eb310ba517292717",
    "message": {
      "bpm": 13,
      "title": "abc 4",
      "description": "abc abc abc"
    }
  },
  {
    "index": 3,
    "timestamp": "2020-05-24 10:42:24.800961 +0300 MSK m=+59.026186210",
    "hash": "5de3b32b6b52a3f73be63e6955e4471d884123e099b80484eb310ba517292717",
    "prev_hash": "a227a5950287495d1f9952ef72ba8ab640bc7a1240a199043ec488f9c03c154f",
    "message": {
      "bpm": 31,
      "title": "abc 3",
      "description": "abc abc abc"
    }
  }
]
```