### Golang Log Parser

### Note
- Use predefined log format like `nginx-access.log` file
- Default output format is Plain Text , you can change it with ( -t json or -t Plain Text )

### Example
`go run main.go /blablabla/nginx-access.log -t json`
or
`go run main.go /blablabla/nginx-access.log -t json -o /blablabla/output.txt`
