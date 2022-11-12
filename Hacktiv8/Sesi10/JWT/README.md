<h1 align="center"> MyGram</h1>
<h2 align="center"> Final Project - Scalable Web Service with Golang - DTS Kominfo </h2>

## About
Aplikasi untuk menyimpan foto dan membuat comment foto orang lain

## Library
- Gorm
- Gin
- jwt-go
- govalidator
- crypto

## Database
Postgress

## Deployment
Heroku

## Endpoints
### Register
```
POST https://mygarm.herokuapp.com/users/register
```
### Login
```
POST https://mygarm.herokuapp.com/users/login  
```
### Get Data User
```
GET https://mygarm.herokuapp.com/users
```
### Update Data User
```
PUT https://mygarm.herokuapp.com/users
```
### Delete Data User
```
DELETE https://mygarm.herokuapp.com/users
```

### Post Photo
```
POST https://mygarm.herokuapp.com/photos
```
### Get All Photo
```
GET https://mygarm.herokuapp.com/photos
```
### Update Photo
```
PUT https://mygarm.herokuapp.com/photos/:photoId
```
### Delete Photo
```
DELETE https://mygarm.herokuapp.com/:photoId
```

### Post Comment
```
POST https://mygarm.herokuapp.com/comments
```
### Get All Comment
```
GET https://mygarm.herokuapp.com/comments
```
### Update Comment
```
PUT https://mygarm.herokuapp.com/comments/:commentId
```
### Delete Comment
```
DELETE https://mygarm.herokuapp.com/:commentId

```
### Post Social Media
```
POST https://mygarm.herokuapp.com/socialmedias
```
### Get All Social Media
```
GET https://mygarm.herokuapp.com/socialmedias
```
### Update Social Media
```
PUT https://mygarm.herokuapp.com/socialmedias/:socialMediaId
```
### Delete Social Media
```
DELETE https://mygarm.herokuapp.com/:socialMediaId
```
## Install
- Git clone https://github.com/vikaputri/MyGram.git 
- Install :
```
go get github.com/asaskevich/govalidator
go get github.com/dgrijalva/jwt-go
go get github.com/gin-gonic/gin
go get golang.org/x/crypto/bcrypt
go get gorm.io/gorm  
go get gorm.io/driver/postgres
```
- Run : go run main.go
