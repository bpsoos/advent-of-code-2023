run day:
    go run main.go day{{day}}

test day:
    go test -v ./day{{day}}

testall:
    go test -v ./...

