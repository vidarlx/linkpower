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

## Build
```
go mod download
./script/build
```

## Run

To run the program with a default data (device, stations), just run:
```
./linkpower
```

If you want to provide a datafile, pass it through `-f` arg, like so:
```
./linkpower -f /path/to/file
```
See an example file structure in `./resources/test_data.json`

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
