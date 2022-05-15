# Microservice Auth and Fetch API Specs

## Auth App
`http://103.157.27.164:3000` [Production]
<br>
`http://localhost:3000` [Local]

### Register

`POST /register`

*Request*
> Body Params : JSON

```json
{
  "username": "rambo_kriting",
  "name": "Dani Prasetya",
  "phone": "082312345678",
  "role": "super-admin"
}
```

*Response*

> HTTP 201 : Example Response

```json
{
  "message": "User created!",
  "data": {
    "_id": "fd9b9d4363524bf4bb8a5ebc7fcb2c12",
    "username": "rambo_kriting",
    "name": "Dani Prasetya",
    "phone": "123456789",
    "role": "super-admin",
    "password": "123123"
  },
  "success": true
}
```

### Signin

`POST /signin`

*Request*
> Body Params : JSON

```json
{
  "phone": "082312345678",
  "password": "lalayeye"
}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "User Authenticated!",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....."
  },
  "success": true
}
```

> HTTP 401 : Example Response

```json
{
  "success": false,
  "data": null,
  "message": "Wrong phone or password !"
}
```

### Profile

`GET /profile`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "OK",
  "data": {
    "username": "rambo_kriting",
    "phone": "123456789",
    "role": "super-admin",
    "timestamp": 45234523453,
    "iat": 45234523453,
    "exp": 45234523453
  },
  "success": true
}
```

> HTTP 401 : Example Response

```json
{
  "success": false,
  "data": {
    "Authorization": "Invalid Token!"
  },
  "message": "Invalid Token!"
}
```
<br></br>
## Fetch App 
`http://103.157.27.164:9090` [Production]
<br>
`http://localhost:9090` [Local]

### Validate
`GET /v1/profile`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "OK",
  "success": true,
  "data": {
    "username": "astalavista",
    "name": "John Travolta",
    "phone": "4522342134",
    "role": "admin",
    "timestamp": 1646669672444,
    "iat": 1646669672444,
    "exp": 1646669672444
  }
}
```

> HTTP 401 : Example Response

```json
{
  "message": "Invalid Token!",
  "success": false
}
```

### Fetch
`GET /v1/fetch`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "OK",
  "success": true,
  "data": [
    {
      "uuid": "23-123-123-93f7-1234",
      "komoditas": "Lele",
      "area_provinsi": "JAWA TIMUR",
      "area_kota": "MADIUN",
      "size": "120",
      "price": "200000",
      "price_usd": "12.12312",
      "tgl_parsed": {},
      "timestamp": null
    }
  ]
}
```

> HTTP 401 : Example Response

```json
{
  "message": "Invalid Token!",
  "data": null,
}
```

> HTTP 403 : Example Response
```json
{
  "message": "Role not valid to access!",
  "data": null,
  "success": false
}
```

### Fetch
`GET /v1/aggregate`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "OK",
  "success": true,
  "data": [
    {
      "year": 2022,
      "month": 3,
      "week": 2,
      "province": "JAWA TIMUR",
      "total_data": 2,
      "size": {
        "max": 23,
        "min": 2,
        "med": 8,
        "avg": 4
      },
      "price": {
        "max": 23,
        "min": 2,
        "med": 8,
        "avg": 4
      }
    }
  ]
}
```

> HTTP 401 : Example Response

```json
{
  "message": "Invalid Token!",
  "success": false
}
```

> HTTP 403 : Example Response
```json
{"message":"Insufficient access level !",
"success":false}
```