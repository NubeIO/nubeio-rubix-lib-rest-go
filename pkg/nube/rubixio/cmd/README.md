# rubix-io rest client

## cmd/cli usage

### start the app

download and install https://github.com/NubeIO/nubeio-rubix-app-pi-gpio-go

#### start the app in test mode

start the app first in debug mode, this will load dummy data

```
go run app.go -debug=true
```

#### start the app in normal mode

this will only work when this is running on the rubix-io

```
go run app.go -debug=false
```

usage of cmd app

```
cd cmd
go run main.go --help
```

### ping

```
go run main.go ping --ip=0.0.0.0 --port=5001 
```

### read all

```
go run main.go read --ip=192.168.15.10 --port=5001
```

### write one

```
go run main.go write --ip=0.0.0.0 --port=5001 --point=UO1 --value=100
```
