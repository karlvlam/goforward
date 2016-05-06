# goforward
simple port forward server in go

### Usage

```bash
./goforward [CONFIG_FILE]
```

### Config Format (as example)

Forward 8080 to hiauntie.com:80
Forward 8081 to hiauntie.com:443
```
0.0.0.0:8080 hiauntie.com:80
0.0.0.0:8081 hiauntie.com:443
```

### Build

```bash
CGO_ENABLED=0 go build goforward.go
```
