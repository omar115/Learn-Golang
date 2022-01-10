here i will explain the hello-world program in go.
and some general F/Q

* some important go commands:
- go build => compiles a bunch of go source code files
- go run => compiles and executes one or two files
- go fmt => formats all the code in each file in the current directory
- go install => compiles and installs a package
- go get => downloads the raw source code of someone else's package
- go test => runs any tests associated with the current project

* go build will build the main.go file and build a executable file
- run go build main.go and you will see a new executable file

* a package is a collection of common source code files
- the ground rule is if you have 3 files under package main, 
- you will write package main in first line of your code then

* executable vs reusable
- executable: generates a file that we run
- reusable: code used as helpers, good place to put reusable logic

* package main is an executable package
- defines a package that can be compiled and then "executed",
- must have a function called "main"

* any other package without main is called reusable package
- defines a package that can be used as a dependency (helper code)

1. How to run the code?
ans: go run main.go

2. what does package main mean ?
ans: 