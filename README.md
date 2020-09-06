# League Backend Challenge

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

##How to access the solution
###File
You need download the .zip file and extract it all files in your computer. Pay attention that this directory needs to be at the $GOPATH.

Your directory you will have the following structure:

league-code-challenge
├─ README.md
├─ cmd
   └─ main.go


###To compile the build
You need run in your terminal the follow in commands.

If you use Windows(terminal): go build ./go.exe

If you use Unix like computer(terminal): go build and after go run .

After this your go server will be listening at localhost:8080. You can open in other terminal and send the following commands to tests all the endpoints:

curl -F 'file=@path/matrix.csv' "localhost:8080/echo"

curl -F 'file=@path/matrix.csv' "localhost:8080/invert"

curl -F 'file=@path/matrix.csv' "localhost:8080/flatten"

curl -F 'file=@path/matrix.csv' "localhost:8080/sum"

curl -F 'file=@path/matrix.csv' "localhost:8080/multiply"

###Testing
You can run the test typing in your terminal go test, after running the build command. In the next section you can see the original declaration of the problem.

