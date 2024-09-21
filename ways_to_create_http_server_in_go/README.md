# Ways to create HTTP server in go
Here is the simple example of _5 ways_ of running **HTTP** server in golang.
1. go [`net/http`](https://pkg.go.dev/net/http) package
2. [`fasthttp`](https://github.com/valyala/fasthttp) framework
3. [`gin`](https://gin-gonic.com/) framework 
4. [`fiber`](https://gofiber.io/) framework
5. [`echo`](https://echo.labstack.com/) framework

## Functionality
For each sample, we have: 
- index route handler
- not found routes handler
- add built-in logger for frameworks (if exists)
- add panic-recovery for frameworks (if exists)

and return a sample static json response.

## How to run:
You can simply run each file independent by using `go run <filename>`, or using the `make` tool:
| name | command |
| :------ | :-------- |
| net/http | `make run_nethttp` |
| fasthttp | `make run_fasthttp` |
| gin | `make run_gin` |
| fiber | `make run_fiber` |
| echo | `make run_echo` |

## More info
For more information, visit each one home page.
