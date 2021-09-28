
# Rest API - Notes CRUD

This was my first ever CRUD done with Golang. It was a challenge cause i hadn't got experience with backend development in general. In spite of that, i am preety confortable with the result.


## API Reference

#### Create a new note

```http
  POST /api/v1/notes/create
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `owner_name` | `string` | **Required** |
| `title` | `string` | **Required** |
| `details` | `string` |  Optional |

#### Get all notes

```http
  GET /api/v1/notes/readall
```


#### Get note by id

```http
  GET /api/v1/notes/readbyid
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int` | **Required** |


#### Update note by id

```http
  PUT /api/v1/notes/update
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int` | **Required** |


#### Delete note by id

```http
  DELETE /api/v1/notes/delete
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int` | **Required** |

## Author

- [@ramirocuencasalinas](https://www.linkedin.com/in/ramiro-cuenca-salinas/)

  
