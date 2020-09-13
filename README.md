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

## GO tools
### Build
```
go mod download
./script/build
```

### Run

To run the program with a default data (device, stations), just run:
```
./linkpower
```

If you want to provide a datafile, pass it through `-f` arg, like so:
```
./linkpower -f /path/to/file
```
See an example file structure in `./resources/test_data.json`

### Test
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

## Docker
### Build
```
docker build -t linkpower 
```
### Run
To run with a sample dataset, simply type:
```
docker run linkpower
```

If you want to use your own file with Device and Stations defined, put 
it to the `resources` directory, build and image and pass as an argument to `docker run`, like so:
```
docker run linkpower -f ./resources/test/ok.json
```

### Test
```
./script/test
```