![logo](https://i.imgur.com/VJiJiMO.png)

## Description
HTTP Static file server written in Go.

## Usage
1. Create `.env` file with fields:
    `PORT` - port at which http will be served
    `FILES_DIR` - location with files which will be served<br>
    Example
    ```
    PORT=8000
    FILES_DIR=/home/user/files/are/here
    ```
2. Use either Docker or Makefile
    * Run with Docker
        `docker-compose up`
    * Run via Makefile
        `make`