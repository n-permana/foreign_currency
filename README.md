
  

### Requirement
- Echo Framework (https://echo.labstack.com/)
- Dep (https://github.com/golang/dep)
- Docker (https://www.docker.com/get-docker)

### Endpoint list
-[POST] http://localhost:7001/api/v1/exchange
> Sample Request
``` 
{
	"from":"USD",
	"to":"IDR"
}
```
> Sample Response
```
{
    "message": "the exhange has successfully saved",
    "succes": "true"
}
```

- [GET] http://localhost:7001/api/v1/exchange
> Sample Response
```
[
    {
        "id": 1,
        "from": "USD",
        "to": "IDR"
    },
    {
        "id": 2,
        "from": "GBP",
        "to": "USD"
    }
]
```

- [GET] http://localhost:7001/api/v1/exchange/<exchange_id>
> Sample Response
```
{
    "from": "GBP",
    "to": "USD",
    "average": 0.20242857142857143,
    "variance": 0,
    "Rates": [
        {
            "date": "2018-07-08",
            "rate": 1.417
        }
    ]
}
```

- [POST] http://localhost:7001/api/v1/exchange_rate
> Sample Request

``` 
{
	"date":"2018-07-08",
	"from":"GBP",
	"to":"USD",
	"rate":1.417
}
```
	
> Sample Response
```
{
    "message": "the exhange rate has successfully saved",
    "succes": "true"
}
```


- [GET] http://localhost:7001/api/v1/exchange_rate/'yyyy-mm-d'
> Sample Response
```
[
    {
        "id": "2",
        "from": "GBP",
        "to": "USD",
        "rate": "1.417",
        "average": "1.4170000553131104"
    },
    {
        "id": "3",
        "from": "USD",
        "to": "IDR",
        "rate": "insufficient data",
        "average": "insufficient data"
    }
]
```

- [DELETE] http://localhost:7001/api/v1/exchange/<exchange_id>
> Sample Response
```
{
    "message": "the exhange has successfully deleted",
    "succes": "true"
}
```
  
### How To Run
- Clone this repo
- Go to main folder
- Run this command
```
docker-compose up
```