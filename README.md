<h1 align="center">
<img src="https://i.imgur.com/VJiJiMO.png">
</h1>

## Description
HTTP Static file server written in Go.

## Usage
1. Create `.env` file with fields:<br>
    `PORT` - port at which http will be served<br>
    `FILES_DIR` - location with files which will be served<br><br>
    Example
    ```
    PORT=8000
    FILES_DIR=/home/user/files/are/here
    ```
2. Use either Docker or Makefile
    * Run with Docker<br>
        `docker-compose up`
    * Run via Makefile<br>
        `make`