# tangelo-flattener
Flattener Service for Arrays

API that will flatten an input array of arbitrarily nested arrays of values (of mixed types) into a flat
array of values, and give us the level of the deepest element in the array

For example:

Input Output
[[10 ; 20 ; 30] ; 40] [10 ; 20 ; 30 ; 40] ; depth=2
[[A ; 20 ; [B]] ; 40] [A ; 20 ; B ; 40] ; depth=3
[[10 ; [[20 ; [30]]] ; [40]]] [10 ; 20 ; 30 ; 40] ; depth=5
[♣ ; ♦ ; ♥] [♣ ; ♦ ; ♥] ; depth=1

| Input | Output |
| ------ | ------ |
| [[10 ; 20 ; 30] ; 40] | [10 ; 20 ; 30 ; 40] ; depth=2 |
| [[A ; 20 ; [B]] ; 40] | [A ; 20 ; B ; 40] ; depth=3 |
| [[10 ; [[20 ; [30]]] ; [40]]] | [10 ; 20 ; 30 ; 40] ; depth=5|
| [♣ ; ♦ ; ♥] | [♣ ; ♦ ; ♥] ; depth=1 |


## Installation

Use the go compiler from base directory

```bash
go run cmd/flattener/main.go
```

## Usage

-----------------------------------------
Make a post with the array to flatten
```bash
curl --location --request POST 'http://localhost:8000/api/v1/flatArray' \
--header 'Content-Type: application/json' \
--data-raw '[[2634, 9867, true, false, ["hola"]], ["chau"]]'
```

Response Example
```json
{ 
  "flatten": [ 2634, 9867, true, false, "hola", "chau" ],
  "depth": 3
}
```

--------------------------------------------------
Get the last 100 request

```bash
curl --location --request GET 'http://localhost:8000/api/v1/flatArray'
```
Response Example

```json

[
  {
    "timestamp":"2021-02-11T00:34:48.585231-03:00",
    "input-array":[8943598534,42347243247,0,"",true,0.234,"😀","😁","😂","🤣"],
    "flattened-array":[8943598534,42347243247,0,"",true,0.234,"😀","😁","😂","🤣"]
  },
  {
    "timestamp":"2021-02-11T00:23:26.500104-03:00",
    "input-array":[[2634,9867,true,false,["hola"]],["chau"]],
    "flattened-array":[2634,9867,true,false,"hola","chau"]
  }
]

```
All this examples and more can be tested with the postman collection Tangelo-flattener.postman_collection.json