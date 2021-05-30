# go-fiber-gorm-example
this is simple rest api using gorm and fiber

to run it just type  
```
go run main.go in the project
```
in the folder
and you can check it on
```
localhost:3000/api/v1/book
```
to get all the books

```
localhost:3000/api/v1/book/1
```
to get data book where the id is equal 1

to make new book just execute
```
localhost:3000/api/v1/book
```
and then fill the author name and title name via postman and execute post

to delete book by id type
```
localhost:3000/api/v1/book/1
```
as example it will delete book where id is equal 1 and then execute delete methode via postman
