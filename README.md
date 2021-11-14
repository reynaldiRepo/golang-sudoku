# Sudoku Solver with Golang



### for running test 
```
go main.go main_test.go
```

endpoint list : 
- "/" is index endpoint just only welcome message
- "/sudoku" the main endpoint to solving sudoku, only receive post request with json data

example curl request :
```
curl --location --request POST 'localhost:8080/sudoku' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Input" : [
        [0,2,0,0,0,0,0,0,0],
        [7,0,0,0,0,0,0,0,0],
        [0,0,0,0,7,0,8,6,0],
        [0,0,0,0,0,0,0,0,8],
        [0,0,0,0,0,5,0,3,6],
        [3,0,4,0,1,0,0,0,0],
        [0,0,0,0,2,0,0,0,9],
        [0,0,3,6,0,0,0,0,0],
        [0,7,0,0,0,0,0,2,0]
    ]
}'
```

Postman JSON URL : https://www.getpostman.com/collections/7c456a19db64b47e9204

example result : 
```
{
    "is_solved":true,
    "message":"sudoku is solved",
    "sudoku":[
        [1,2,5,3,6,8,4,9,7],
        [7,6,8,1,4,9,3,5,2],
        [4,3,9,5,7,2,8,6,1],
        [6,5,2,9,3,7,1,4,8],
        [9,1,7,4,8,5,2,3,6],
        [3,8,4,2,1,6,9,7,5],
        [8,4,6,7,2,3,5,1,9],
        [2,9,3,6,5,1,7,8,4],
        [5,7,1,8,9,4,6,2,3]
    ]
}
```