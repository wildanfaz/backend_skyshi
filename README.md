# Backend Test - Skyshi

## Installation

### 1. Clone
```bash
git clone https://github.com/wildanfaz/backend_skyshi.git
```

### 2. Open Folder
```bash
cd backend_skyshi
```

### 3. Dependencies
```bash
go mod tidy
```

### 4. Run Golang App On Port 3030
```bash
go run main.go
```

## Docker Image

### 1. Pull Image

```bash
docker pull muhamadwildanfaz/backend_skyshi:v1
```

### 2. Create Network

```bash
docker network create skyshi
```

### 3. MySQL

Powershell
```bash
docker run --network=skyshi --name db_skyshi -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=xxxx -e MYSQL_PASSWORD=xxxxx -e MYSQL_DATABASE=todo4 -v ${pwd}/init.sql/:/docker-entrypoint-initdb.d/init.sql -dp 3306:3306 mysql
```

### 4. Run Image
```bash
docker run --network=skyshi --name backend_skyshi -e MYSQL_HOST=db_skyshi -e MYSQL_USER=xxxx -e MYSQL_PASSWORD=xxxxx -e MYSQL_DBNAME=todo4 -e MYSQL_PORT=3306 -p 8090:3030 muhamadwildanfaz/backend_skyshi:v1
```