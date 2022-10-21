# Project (Scalable Web Service with Golang - DTS Kominfo) - MyGram
Aplikasi untuk menyimpan foto dan membuat comment foto orang lain

## Register
Path: [http://localhost:8080/users/register ](http://localhost:8080//users/register)'

Method: Post

Request 

1. body:
```
{

    "age":20,
    
    "email":"vika@gmail.com",
    
    "password":"password",
    
    "username":"vika"
    
}
```

Response

1. Status 201

2. Data
```
{

    "age":20,
    
    "email":"vika@gmail.com",
    
    "id":1,
    
    "username":"vika"
    
}
```
## Login
Path: [http://localhost:8080/users/login ](http://localhost:8080//users/login)

Method: Post

Request 

1. body:
```
{
  
    "email":"vika@gmail.com",
    
    "password":"password",
    
}
```

Response

1. Status 200

2. Data
```
{

    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZpa2FAZ21haWwuY29tIiwiaWQiOjV9.LCpNoDe29cp6KUjclllQ0EPNwqqQbol8ibvaYC9GVMQ",
    
}
```

## Update User
Path: [http://localhost:8080/users ](http://localhost:8080//users)

Method: Put

Request 

1. headers: Authorization (Bearer token string)

2. body:
```
{

    "age":19,
    
    "email":"vikaputri@gmail.com",
    
    "password":"passw0rd",
    
    "username":"vikaputri"
    
}
```

Response

1. Status 200

2. Data
```
{

    "age":19,
    
    "email":"vikaputri@gmail.com",
    
    "id":1,
    
    "username":"vikaputri"
    
    "updatedAt":"2022-10-17T11:37:19.901Z"
    
}
```
## Delete User
Path: [http://localhost:8080/users ](http://localhost:8080//users)

Method: Delete

Request

1. headers: Authorization (Bearer token string)

Response :
1. Status 200

2. Data:
```
{

    "message": "Your account has been successfully deleted"
    
}
```
