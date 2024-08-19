# REST API with Go, GORM, and PostgreSQL

This project is a simple REST API built with Go, GORM. The API includes JWT authentication using RSA keys.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [RSA Key](#rsa)

## Requirements

- Go 1.23 or later
- PostgreSQL
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [OpenSSL](https://www.openssl.org/) (for generating RSA keys)

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/thaksharadhananjaya/go_rest_api.git
cd go-restapi
```
### 2. Setup evironment variables using `sample.env` file
### RSA key generate 
``` bash
openssl genrsa -out private.pem [time]
openssl rsa -in key.pem -outform PEM -pubout -out public.pem
```

`` Not required private key integration into application! Required only Public key``


