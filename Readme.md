### Golang Log Parser

### Note
- Use predefined log format like `nginx-access.log` file
- Default output format is Plain Text , you can change it with ( -t json or -t Plain Text )

### Example
- `./main /home/dimas/personal/LogicalTest/nginx-access.log -t json`

- `./main /home/dimas/personal/LogicalTest/nginx-access.log -t json -o /blablabla/output.txt ` 

- `go run main.go /blablabla/nginx-access.log -t json`

- `go run main.go /blablabla/nginx-access.log -t json -o /blablabla/output.txt`
