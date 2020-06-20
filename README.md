Simple port scanner

Usage:

go run main.go -ip=... -port=... -T=...

example:

go run main.go -ip=example.org -port=1000 -T=5

ip = ip address

port = port range, default=1000

-T = number of threads (maximum 5), default=1