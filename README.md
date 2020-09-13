# linkpower


The program solves the most suitable (with most power) link station for a 
device at givenpoint [x,y].

The solution resolves the problem in 2-dimensional space. Link stations 
have reach and power.

A link station’s power can be calculated:
```
power = (reach - device's distance from linkstation)^2
if distance > reach, power = 0
```

## Build & run

### local 
```
go mod download
go run cmd/linkpower/main.go
```

### docker
```
docker build . -t linkpower
docker run linkpower
```

## Test
```
go test ./...
```

Coverage
```
go test -coverprofile cover.out ./...
```
Show covered lines in a browser
```
go tool cover -html=cover.out
```

Show covered lines in a browser
```
go tool cover -html=cover.out
```
