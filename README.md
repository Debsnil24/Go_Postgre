
# GO_Postgre

This is a Golang project that uses a PostgreSQL backend server & REST API to manage Stock data using CRUD operations


## Roadmap

### Project Structure
![ProjectStructure](https://github.com/user-attachments/assets/63e6608f-9dec-41bc-a399-aa1172365f13)

### Routes
![Routes](https://github.com/user-attachments/assets/b43ffd9f-91c2-426d-926b-e040174ed9cc)



## Testing Demo

Refer to the provided Google Doc for the [Testing Demo](https://docs.google.com/document/d/e/2PACX-1vQ-YSw469_yzCocTWc_d9lv-yc3srWWyYJL7DK6qoIXOpndmLt7heexBgM16qqEWfvu_C3-6IJNe6X8/pub)




## Run Locally

#### Clone the project

```bash
  git clone https://github.com/Debsnil24/Go_Postgre.git
```

#### Go to the project directory

```bash
  cd {Directory}/Go_Postgre
```

#### Install dependencies

```go
  go get "github.com/gorilla/mux"
```
```go
  go get "github.com/joho/godotenv"
```
```go
  go get "github.com/lib/pq"
```

#### Pre-Requisites

- Ensure to have PostgreSQL installed on your Local Machine
*(Note: Don't use the postgres user profile for this project)* 
- Login to Admin Profile
```bash
    psql -U postgres
```
*Enter the Password if prompted*

- Create a New User 
```sql
    CREATE DATABASE username;
    CREATE USER username WITH PASSWORD 'yourpassword';
    ALTER USER username WITH SUPERUSER;
```
*Quit using \q & Relogin using the new user*
- Create a New Database called stockdb
```sql
    CREATE DATABASE stockdb;
```
- Connect to the Database 
```sql
    \c stockdb;
```
- Create a new table
```sql
    CREATE TABLE stocks (
        stockid SERIAL PRIMARY  KEY,
        name TEXT,
        price INT,
        company TEXT
    );
```
- Update the .env file with username and yourpassword

#### Starting the server
- Go into the main folder 
```bash
  cd {Directory}/Go_Postgre/main
```
- Build the project
```go
  go build main.go
```
- Run the main.go file
```go
  go run main.go
```


## Acknowledgements

#### Check Out Akhil Sharma's Youtube Video for Project Tutorial
[GO with PostgreSQL - A Different Way ! (Stocks API)](https://youtu.be/1nLH4J-DRLg?si=sVuS6v7NsJHTy67a)



