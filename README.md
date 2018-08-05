
### Requirement
- Echo Framework (https://echo.labstack.com/)
- Dep (https://github.com/golang/dep)
- Docker (https://www.docker.com/get-docker)
  

### Endpoint list 
```
- [POST] http://localhost:7001/api/v1/exchange
- [GET] http://localhost:7001/api/v1/exchange
- [GET] http://localhost:7001/api/v1/exchange/<exchange_id>
- [POST] http://localhost:7001/api/v1/exchange_rate
- [GET] http://localhost:7001/api/v1/exchange_rate/<yyyy-mm-dd>
- [DELETE] http://localhost:7001/api/v1/exchange/<exchange_id>
```

### How To Run
- Clone this repo
- Go to main folder
- Run this command
```
  docker-compose up
```