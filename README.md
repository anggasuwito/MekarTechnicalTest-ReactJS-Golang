# MekarTechnicalTest-ReactJS-Golang

1. import db_mekar.sql in backend-golang/database to your database
2. setting configuration value in backend-golang/configvalue.env
3. at directory backend-golang type command "go run main.go"
4. at directory frontend-reactjs type command "npm install" to install all depedencies
5. manual hit point if not using react js
  OPEN POSTMAN
  1. ADMIN
  ADMIN LOGIN
  - POST
    localhost:port/admin/login
    EXAMPLE
    localhost:8001/admin/login
  - BODY JSON
    {
      "username": "string",
      "password": "string"
    }
    EXAMPLE
    {
      "username": "admin",
      "password": "admin"
    }
    
  2. WORKS
  GET ALL WORKS
  - GET (REQUIRE Header : {Authorization : "Bearer {token}"})
    localhost:port/work/works
    EXAMPLE
    localhost:8001/work/works
    
  3. STUDY
  GET ALL STUDIES
  - GET (REQUIRE Header : {Authorization : "Bearer {token}"})
    localhost:port/study/studies
    EXAMPLE
    localhost:8001/study/studies
    
  4. USER
  GET ALL USER
  - GET (REQUIRE Header : {Authorization : "Bearer {token}"})
    localhost:port/user/users?keywords={keywords}&page={page}&limit={limit}
    EXAMPLE (WITHOUT KEYWORDS)
    localhost:8001/user/users?keywords=&page=1&limit=3
    EXAMPLE (WITH KEYWORDS)
    localhost:8001/user/users?keywords=user&page=1&limit=3
    
  GET USER BY ID
  - GET (REQUIRE Header : {Authorization : "Bearer {token}"})
    localhost:port/user/{userId}
    EXAMPLE
    localhost:8001/user/1
    
  CREATE NEW USER
  - POST (REQUIRE Header : {Authorization : "Bearer {token}"})
    localhost:port/user/createuser
    EXAMPLE
    localhost:8001/user/1
    
  - BODY JSON
    {
      "name": string,
      "birthDate": string (format date yyyy-mm-dd),
      "numberIdCard": string,
      "workId": int,
      "studyId": int
    }
    EXAMPLE
    {
      "name": "user A",
      "birthDate": "2020-03-20",
      "numberIdCard": "717107290001",
      "workId": 1,
      "studyId": 1
    }
    
  UPDATE USER
  - PUT (REQUIRE Header : {Authorization : "Bearer {token}"})
    localhost:port/user/updateuser/{userId}
    EXAMPLE
    localhost:8001/user/updateuser/1
    
  - BODY JSON
    {
      "name": string,
      "birthDate": string (format date yyyy-mm-dd),
      "numberIdCard": string,
      "workId": int,
      "studyId": int
    }
    EXAMPLE
    {
      "name": "user A (update)",
      "birthDate": "1996-03-25",
      "numberIdCard": "123456",
      "workId": 2,
      "studyId": 2
    }
    
  DELETE USER
  - DELETE (REQUIRE Header : {Authorization : "Bearer {token}"})
    localhost:port/user/deleteuser/{userId}
    EXAMPLE
    localhost:8001/user/deleteuser/1
    
