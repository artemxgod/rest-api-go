# REST-API-GO

* All materials are taken from [this course](https://youtube.com/playlist?list=PLehOyJfJkFkJ5m37b4oWh783yzVlHdnUH)

## Github repositories that might help

* [Code structuring in golang](https://github.com/golang-standards/project-layout)
* [different code practices in go](https://github.com/codeship/go-best-practices)
* [extended package for db in go](https://github.com/jmoiron/sqlx)
* [migration for db](https://github.com/golang-migrate/migrate)
* [package for data validation](https://github.com/go-ozzo/ozzo-validation)
* [httpie allows to do request and get respond easily](https://github.com/httpie/httpie)



## Tutorials

* [extended package for db in go tutorial](https://github.com/jmoiron/sqlx)
* [standart package for db in go tutorial](http://go-database-sql.org)
* [postgres for mac tutorial](https://www.codementor.io/@engineerapart/getting-started-with-postgresql-on-mac-osx-are8jcopb) 

## Helpful commands

> Some commands from the course did not work for various reasons 
> Here are the commands that I used instead:

* example of creating migrations 
`migrate create -ext sql -dir migrations create_users`
* example of using migration command 
`migrate -path migrations/ -database postgres://postgres:postgres@localhost/restapi_dev?sslmode=disable up`
* example of curl usage with json 
`curl -X POST -d '{"email":"ac@mail.ru", "password":"password"}' -v -i 'http://localhost:8080/users'`
* creating session with cookie
`http -v --session=user POST http://localhost:8080/sessions email=a1@mail.ru password=password`
* request to private info with cookie 
`http -v --session=user http://localhost:8080/private/whoami`
* request with certain origin and cookie
`http -v --session=user http://localhost:8080/private/whoami "Origin: google.com"`


## Apps
* Postman (for requests)
* Openserver (for database)


