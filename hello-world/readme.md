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

* func is short form of function
- works as a traditional function

1. How to run the code?
ans: go run main.go

2. What is fmt mean?
- it means give my package name access to all of the code, all of the functionality,
  that is contained in other package called fmt
- it is inside the Go programming language by default
- fmt == format

3. how did main.go file organize?
- 1. package declaration
- 2. import required packages
- 3. declare function and tell to do things
