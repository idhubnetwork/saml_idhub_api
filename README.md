# IDHub SAML API
APIs for login AWS with IDHUB account and authority management.

## Feature Overview
* SAML Response
* SAML Identity Provider
* SAML Metadata

## Requirements
* Go
* MySQL
* OpenSSL

## Installation
1. Create a MySQL database with name "saml_idhub_api" and services at localhost:3306.
2. Git clone
3. Install go packages
``` 
go get github.com/astaxie/beego
go get github.com/astaxie/beego/orm
go get github.com/go-sql-driver/mysql
go get github.com/goware/saml
go get github.com/RobotsAndPencils/go-saml/util
go get github.com/stretchr/testify/assert
go get github.com/kardianos/osext 
```
4. ` cd saml_idhub_api `
5. ` go build main.go `
6. ` ./main `

## License
[The MIT License (MIT)](./LICENSE)