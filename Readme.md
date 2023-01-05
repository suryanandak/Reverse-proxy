
# Reverse proxy cli in GoLang

  This is simple Reverse proxy cli tool written in GO.

## A fully functional cli tool to route http traffic from localhost to target IP on specified port!  

This project is a simple implementation of reverse proxy written in GO.  

* Create a simple reverse proxy  
* Routing traffic from host to the target host

## How to build   

### Build for window

Generate an executable **`revproxy.exe`** for windows environment.   

```bash
env GOOS=windows GOARCH=amd64 go build -o revproxy.exe -ldflags "-s -w"
```

### Build for linux

```bash
go build -o revproxy -ldflags "-s -w"
```

## Example command

```bash
./revproxy -lhost 127.0.0.1 -lport 9011 -thost 127.0.0.1 -tport 9008
```

```
$ ./revproxy help
-help  
      Print default help  
-lhost string  
      Listening Host IP Address  
-lport int  
      Listening port  
-thost string  
      Pointing to target Host IP Address  
-tport int  
      Pointing to the target port  
```


## Find a bug?

If you found an issue or would like to submit an improvement to this project, please submit an issue using the issues tab above.

