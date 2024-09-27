# 5 ways to create an HTTP server in go
This directory holds the codes for the [5 ways to create an HTTP server in Golang](https://m8e.ir/l/?p=2536308343) blog post.
Here is a simple example of _5 ways_ of running **HTTP** server in Golang.
1. go [`net/http`](https://pkg.go.dev/net/http) package
2. [`fasthttp`](https://github.com/valyala/fasthttp) framework
3. [`gin`](https://gin-gonic.com/) framework 
4. [`fiber`](https://gofiber.io/) framework
5. [`echo`](https://echo.labstack.com/) framework

## Functionality
For each sample, we have: 
- index route handler
- not found routes handler
- add a built-in logger for frameworks (if exists)
- add panic recovery for frameworks (if exists)

and return a sample static JSON response.

## How to run:
You can simply run each file independently by using `go run <filename>`, or using the `make` tool:
| name | command |
| :------ | :-------- |
| net/http | `make run_nethttp` |
| fasthttp | `make run_fasthttp` |
| gin | `make run_gin` |
| fiber | `make run_fiber` |
| echo | `make run_echo` |

## More info
For more information, visit each home page.
