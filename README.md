# League Backend Challenge


## How to access the solution

You need download the .zip file and extract it all files in your computer. 

To run this code you will need to have [Golang](https://golang.org/dl/) installed in your computer.
To follow these instructions you need include csv file into the directory fixtures. 

Your directory you will have the following structure:

```
league-code-challenge-solution
├─ controllers
   └─ controllers.go
   └─ controllers_test.go
├─ fixtures
   └─ matrix.csv
   └─ matrix_empty.csv
    └─ matrix_not_square.csv    
├─ matrix
    └─ matrix.go
    └─ matrix_test.go
├─ routes
    └─ routes.go 
├─ main.go
├─ README.md
├─ go.mod

```
## Rules
To start run solution you need open terminal into the root of 
directory. 
> ``go build``
> ``go run main.go``

## Expected Endpoints



#### /echo

To run this /echo you need open a new terminal into root the project:

> ``curl -F 'file=@./fixtures/matrix.csv' "localhost:8000/echo"``

   
    // Expected output
    1,2,3
    4,5,6
    7,8,9
     

#### /invert

To run this /echo you need open a new terminal into root the project:


> ``curl -F 'file=@./fixtures/matrix.csv' "localhost:8080/invert"``

    // Expected output
    1,4,7
    2,5,8
    3,6,9


#### /flatten

To run this /echo you need open a new terminal into root the project:
 

> ``curl -F 'file=@./fixtures/matrix.csv' "localhost:8080/flatten"``

   // Expected output
    1,2,3,4,5,6,7,8,9

#### /sum

To run this /echo you need open a new terminal into root the project:
 

> ``curl -F 'file=@./fixtures/matrix.csv' "localhost:8080/sum"``

    // Expected output
    45

#### /multiply

To run this /echo you need open a new terminal into root the project:

> ``curl -F 'file=@./fixtures/matrix.csv' "localhost:8080/invert"``

    // Expected output
    362880


## Run Test

> matrix
```cd matrix```
``` go test . ```

> controllers
```cd controllers```
``` go test . ```

## About The Challenge

In main.go you will find a basic web server written in GoLang. It accepts a single request _/echo_. Extend the webservice with the ability to perform the following operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```
_______________________________
