# go-api
go api deploy on Azure Function

##Sunmary module to install
```bash
go get github.com/gofiber/fiber/v2
go get -u gorm.io/gorm gorm.io/driver/sqlserver
go get -u github.com/joho/godotenv

```
##step-by-step[Reccomance Link](https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows)

##On local
1. Create a New Function

2. Create a component

```bash
project/
├── main.go        
├── .env            
├── database/
│   └── connection.go   
├── controllers/        
│   └── hello_controller.go  
│   └── user_controller.go
├── routes/             
│   └── routes.go  
└── models/
    └── user.go      
```

3. Create go.mod
```bash
go mod init <module-name>
```

4. Update host.json
```bash
  "customHandler": {
    "description": {
      "defaultExecutablePath": "main.exe",  //file name need to run frist
      "workingDirectory": "",
      "arguments": []
    },
    "enableForwardingHttpRequest": true
  }
```
6. install all module

7. build
```bash
go build main.go
```

##On Azure

8. Create Azure Function

9. Deploy code to Azure Function

10. Open system assigned set status on

11. Set App settings on Environment variables: add app setting name "DB_CONNECTION_STRING" 
```bash
DB_CONNECTION_STRING=sqlserver://your-username:your-password@your-server.database.windows.net:1433?database=your-database&encrypt=true
```

##On Azure SQL
12. Create SQL server (set Authentication method is Use both SQL and Microsoft Entra authentication then add username and password of admin)

13. Create SQL Database

14. Create table
```bash
CREATE TABLE Users (
    ID INT PRIMARY KEY IDENTITY(1,1),  -- Set ID as PRIMARY KEY and have auto-increment value.
    Name VARCHAR(255) NOT NULL,         -- The Name field is set to a VARCHAR of size 255 and cannot be NULL.
    Email VARCHAR(255) NOT NULL UNIQUE  -- The Email field is set to a VARCHAR of size 255, cannot be NULL, and is unique.
);

```

15. Insert data
```bash
INSERT INTO Users (Name, Email)
VALUES 
    ('John Doe', 'john.doe@example.com'),
    ('Jane Smith', 'jane.smith@example.com'),
    ('Alice Johnson', 'alice.johnson@example.com');
```

16. Set networking of sql server (config Firewall rules)

17. deploy to Azure Function



