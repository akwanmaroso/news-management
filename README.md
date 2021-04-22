# News Management

## List API

- `api/news` [GET] -> Get All News 
- `api/news?topic=` [GET] -> Get All News Filter topic 
- `api/news?status=` [GET] -> Get All News Filter status 
- `api/news/{news_id}` [GET] -> Get News based on id 
- `api/news` [POST] -> Create A News 
- `api/news` [PUT] -> Update News
- `api/news` [DELETE] -> Delete News

- `api/tags` [GET] -> Get All Tags  
- `api/tags` [POST] -> Create A Tags 
- `api/tags` [PUT] -> Update Tags
- `api/tags` [DELETE] -> Delete Tags
 

## API SPECS

### News
### Save News
Request:
- Method: `POST`
- Endpoint: `/api/news`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`
- Body
```json
{
  "topic": "Investment",
  "content": "Overseas Investment News",
  "status": "draft",
  "tags": ["investment", "wealth"]
 }
```

Response
```json
   "data": {
   "topic": "Investment",
      "content": "Overseas Investment News",
      "status": "draft",
      "tags": [
          {
            "name":"investment" 
          }, {
            "name":"wealth"
          }
      ]
   },
   "status": 201,
   "success": true
```

### Get All News
Request:
- Method: `GET`
- Endpoint: `/api/news`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`
- Query Param:
    - status: string
    - topic: string

- Response
```json
  "data":[
    {
      "topic": "Invesment in the easy times",
      "content": "Berita Investasi Luar negeri",
      "status": "draft",
      "tags": [
          {
            "name":"invesment" 
          }, {
            "name":"wealth"
          }
       ]
     },{
      "topic": "Persija Vs Persib",
      "content": "Football Match News",
      "status": "publish",
      "tags": [
          {
            "name":"football" 
          }, {
            "name":"sport"
          }
        ] 
      }
  ],
  "status": 200,
  "success": true,
    
```

### Get News
Request:
- Method: `GET`
- Endpoint: `/api/news/{news_id}`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`

- Response
```json
    "data": {
      "topic": "Ball Investing",
      "content": "Overseas Investment New",
      "status": "draft",
      "tags": [
          {
            "name":"investasi" 
          }, {
            "name":"wealth"
          }
       ]
    },
    "status": 200,
    "success": true
```

### Update News
Request:
- Method: `PUT`
- Endpoint: `/api/news/{news_id}`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`
- Body
```json
{
  "topic": "string|required",
  "content": "string|required",
  "status": "string[Draft, Publish, Deleted]|required",
  "tags": ["arrayOfNameTag"]
 }
```
- Example
```json
{
  "topic": "Ball Investing",
  "content": "Overseas Investment New",
  "status": "Publish",
  "tags": ["investment", "wealth"]
 }
```

- Response
```json
  "data": {
  "topic": "Ball Investing",
  "content": "Overseas Investment New",
  "status": "draft",
  "tags": [
      {
        "name":"investment" 
      }, {
        "name":"wealth"
      }
    ]
  },
  "status": 200,
  "success": true
```

### Delete News
Request:
- Method: `DELETE`
- Endpoint: `/api/news/{news_id}`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`

- Response
```json
    {
      "data": "News With ID: 7 is successfully deleted",
      "status": 200,
      "success": true
    }
```


## Tag
### Save Tag
Request:
- Method: `POST`
- Endpoint: `/api/tags`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`
- Body
```json
{
  "name": "Investment"
 }
```

Response
```json
   "data": {
    name:"Invesment"
   },
   "status": 201,
   "success": true
```
### Get All Tag
Request:
- Method: `GET`
- Endpoint: `/api/tags`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`
- Response
```json
  "data":[
    {
      "name":"invesment" 
    },{
      "name":"wealth"
    }
  ],
  "status": 200,
  "success": true,
    
``` 
### Get Tag 
Request:
- Method: `GET`
- Endpoint: `/api/tags/{tag_id}`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`
- Response
```json
  "data": {
      "name":"invesment" 
  },
  "status": 200,
  "success": true,
    
``` 
### Update Tag 
### Update News
Request:
- Method: `PUT`
- Endpoint: `/api/tag/{tag_id}`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`
- Body
```json
{
  "name": "Ball Investing Update"
 }
```

- Response
```json
  "data": {
    "name": "Ball Investing Update"
  },
  "status": 200,
  "success": true
```
### Delete Tag 
Request:
- Method: `DELETE`
- Endpoint: `/api/tags/{tag_id}`
- Header: 
    - Content-Type: `application/json`
    - Accept: `application/json`

- Response
```json
    {
      "data": "Tag With ID: 7 is successfully deleted",
      "status": 200,
      "success": true
    }
```