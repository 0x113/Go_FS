<h1 align="center">
<img src="https://i.imgur.com/VJiJiMO.png">
</h1>

## Description
HTTP Static file server written in Go.

<img src="https://i.imgur.com/5CHrSnq.png">

## How To Use
1. Clone this repository
	```
	$ git clone https://github.com/0x113/GO_FS
	$ cd GO_FS
	```
2. Run using either Docker or Golang
    * Run with Docker
		```
		$ docker build -t gofs .
		$ docker run -p 5000:8000 -it gofs
		```
		Open `localhost:5000` and enjoy.
    * Run via Makefile
		```
		$ make
		$ ./bin/gofs
		```
		Open `localhost:8000` and enjoy.

## Options
### `-path`
Serve specified directory <br>`$ gofs -path="/home/user/files/are/here"` - this will serve specified directory on default port (`8000`)
### `-port`
Serve directory on specified port <br>
`$ gofs -port="5000"` - this will serve current directory on port `5000`

## TODO
- [x] Error page
- [ ] Browse files
- [ ] Search files
- [ ] Upload files
- [x] Create new folder
- [ ] Browse files by type
	* [ ] Movies
	* [ ] TV Shows
	* [ ] Music
	* [ ] Docs
- [ ] QR code
