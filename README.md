# GoSqlServer
Initial project to learn Go with Gin, Gorm, Sql Server

##
For Go, there are several tools that provide static code analysis, similar to what ReSharper offers for C#. Here are some of the best static code analysis tools for Go:

1. GolangCI-Lint
   Description: GolangCI-Lint is a fast, configurable, and extensible linter for Go. It aggregates multiple linters into a single tool, making it a comprehensive solution for static code analysis.
   Features:
   Aggregates multiple linters (e.g., govet, golint, staticcheck, ineffassign, errcheck, etc.)
   Highly configurable
   Fast and parallel execution
   Installation:
   Usage:
2. Staticcheck
   Description: Staticcheck is a state-of-the-art linter for Go that focuses on finding bugs, performance issues, and other problems in your code.
   Features:
   Detects bugs and performance issues
   Provides suggestions for code improvements
   Integrates well with CI/CD pipelines
   Installation:
   Usage:
3. Go Vet
   Description: Go Vet is a built-in tool that examines Go source code and reports suspicious constructs, such as those likely to be bugs or performance issues.
   Features:
   Part of the Go toolchain
   Detects common mistakes and potential issues
   Usage:
4. Gosec (Go Security Checker)
   Description: Gosec is a security-focused static analysis tool that scans your Go code for common security issues.
   Features:
   Identifies security vulnerabilities
   Provides detailed reports
   Integrates with CI/CD pipelines
   Installation:
   Usage:
5. Revive
   Description: Revive is a fast, configurable, and extensible linter for Go, designed to be a drop-in replacement for golint.
   Features:
   Highly configurable
   Extensible with custom rules
   Fast execution
   Installation:
   Usage:
   Integrated Development Environments (IDEs)
   GoLand: JetBrains GoLand is a powerful IDE for Go that provides many features similar to ReSharper, including code analysis, refactoring, and navigation.
   Visual Studio Code: With the Go extension, Visual Studio Code provides a good development experience for Go, including linting, code navigation, and debugging.
   Conclusion
   For comprehensive static code analysis in Go, GolangCI-Lint is highly recommended as it aggregates multiple linters and provides extensive configuration options. 
6. Additionally, using Staticcheck and Gosec can help identify bugs, performance issues, and security vulnerabilities in your code. 
7. For an IDE experience similar to ReSharper, GoLand by JetBrains is a great choice.



Packages and/or third party Go libraries executed in this project so far:

$ go get -u github.com/gin-gonic/gin
$ go get -u gorm.io/gorm
$ go get gorm.io/driver/postgres  | go get -u gorm.io/driver/sqlserver
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag init
$ go get github.com/stretchr/testify
$ go get github.com/stretchr/testify/assert@v1.10.0

Run the tests

Navigate to the directory containing your test files: Open a terminal and navigate to the directory where your test files are located. 
For example, if your test files are in the controller directory, navigate to that directory:
$ cd path/to/your/project/controller

Run the tests: Use the go test command to run the tests. This command will automatically discover and run all test files (files ending with _test.go) in the current 
directory and its subdirectories.
$ go test

Run tests with verbose output: If you want more detailed output, you can use the -v flag:
$ go test -v

Run a specific test file: If you want to run tests in a specific file, you can specify the file name:
$ go test -v userController_test.go

Run tests with coverage: If you want to see the test coverage, you can use the -cover flag:
$ go test -cover

Example Output
When you run the tests, you should see output similar to this:

This indicates that the TestFindById test passed successfully.

Summary
Use go test to run all tests.
Use go test -v for verbose output.
Use go test -v userController_test.go to run tests in a specific file.
Use go test -cover to see test coverage.
These commands will help you run and verify your tests in Go.

*******************************************************************************
Running the test
PS C:\Projects\Others\Go\GoSqlServer\controller> go test -v userController_test.go
# command-line-arguments
C:\Users\joaolopes\go\pkg\mod\github.com\stretchr\testify@v1.10.0\mock\mock.go:14:2: missing go.sum entry for module providing package github.com/davecgh/go-spew/spew (imported by github.com/stretchr/testify/assert); to add:    
go get github.com/stretchr/testify/assert@v1.10.0
FAIL    command-line-arguments [setup failed]
FAIL
PS C:\Projects\Others\Go\GoSqlServer\controller> 

________
GitHub Copilot explanation

The error indicates that there is a missing go.sum entry for a dependency. You can resolve this by running the go get command to fetch the required dependencies.

1. Fetch the required dependencies:
$ github.com/stretchr/testify/assert@v1.10.0

2. Run the tests again:
$ go test -v userController_test.go

Full Steps:
1. Open your terminal and navigate to your project directory:
$ cd C:\Projects\Others\Go\GoSqlServer\controller

2. Fetch the required dependencies:
$ go get github.com/stretchr/testify/assert@v1.10.0

3. Run the tests:
$ go test -v userController_test.go

This should resolve the missing dependency issue and allow you to run your tests successfully.

___________
It continued to show error
I went to root folder and:
PS C:\Projects\Others\Go\GoSqlServer> go build main.go
C:\Users\joaolopes\go\pkg\mod\github.com\stretchr\testify@v1.10.0\mock\mock.go:16:2: missing go.sum entry for module providing package github.com/stretchr/objx (imported by github.com/stretchr/testify/mock); to add:
go get github.com/stretchr/testify/mock@v1.10.0

So in the root folder I ran again:
PS C:\Projects\Others\Go\GoSqlServer> go get github.com/stretchr/testify/mock@v1.10.0
go: downloading github.com/stretchr/objx v0.5.2
PS C:\Projects\Others\Go\GoSqlServer> 

Now it got
downloading github.com/stretchr/objx v0.5.2

______
Correndo o teste individualmente ele falha (??? Porque não consegue resolver todas as dependencias e imports ???)
PS C:\Projects\Others\Go\GoSqlServer\controller> go test -v userController_test.go
# command-line-arguments [command-line-arguments.test]
.\userController_test.go:17:20: undefined: NewUsersController
FAIL    command-line-arguments [build failed]
FAIL

Indo na mesma para a pasta controller mas dando um comando mais abrangente o teste já passa
PS C:\Projects\Others\Go\GoSqlServer\controller> go test -v
=== RUN   TestFindById
[GIN] 2025/01/16 - 22:00:43 | 200 |            0s |                 | GET      "/user/1"
--- PASS: TestFindById (0.00s)
PASS
ok      github.com/lopesoliveira/GoSqlServer/controller 0.237s
PS C:\Projects\Others\Go\GoSqlServer\controller> 
_____
Descobri isto porque se fizer run no IDE ao teste sem ser por command line o teste estava a passar
