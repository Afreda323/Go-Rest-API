# Go Todo API

I know, it's another todo app.  Oh well, it's a good way to get familiar with API development using a new language

## Usage

Create a ```.env``` file using the ```.env.schema```.  Make sure all of the values are filled out.

Then in the terminal run ```go run main.go```

## Routes 

## Sign Up

Create an account, saved to database. JWT response

#### HTTP Request

```POST - /api/v1/users/signup```

##### Request body

| Name   | Type       | Required  |
|:-------|:-----------|:----------|
|email   |```String```| ```true```|
|password|```String```| ```true```|

#### Example

##### Request

```json
{
    "email": "user@email.com",
    "password": "abc123"
}
```

##### Response

```json
{
    "message": "User created",
    "status": true,
    "user": {
        "ID": 2,
        "CreatedAt": "2018-12-20T15:15:19.489462-05:00",
        "UpdatedAt": "2018-12-20T15:15:19.489462-05:00",
        "email": "user@email.com",
        "token": "..."
    }
}
```

---

## Log In

Log into an existing account. JWT response

#### HTTP Request

```POST - /api/v1/users/login```

##### Request body

| Name   | Type       | Required  |
|:-------|:-----------|:----------|
|email   |```String```| ```true```|
|password|```String```| ```true```|

#### Example

##### Request

```json
{
    "email": "user@email.com",
    "password": "abc123"
}
```

##### Response

```json
{
    "message": "Logged in",
    "status": true,
    "user": {
        "ID": 2,
        "CreatedAt": "2018-12-20T15:15:19.489462-05:00",
        "UpdatedAt": "2018-12-20T15:15:19.489462-05:00",
        "email": "user@email.com",
        "token": "..."
    }
}
```

---

## Create Todo

Save a todo to the database.

#### HTTP Request

```POST - /api/v1/todos/```

##### Request Headers
| Name        | Type         | Required  |
|:------------|:-------------|:----------|
|Authorization|Bearer {Token}| ```true```|

##### Request body

| Name   | Type       | Required  |
|:-------|:-----------|:----------|
|value   |```String```| ```true```|

#### Example

##### Request

```Coming soon...```

##### Response

```Coming soon...```

---

## Get Todos

Retrieve all of your todos from the database

#### HTTP Request

```GET - /api/v1/todos/```

##### Request Headers
| Name        | Type         | Required  |
|:------------|:-------------|:----------|
|Authorization|Bearer {Token}| ```true```|

#### Example

##### Request

```Coming soon...```

##### Response

```Coming soon...```

---

## Get Todo

Retrieve a single todo from the database

#### HTTP Request

```GET - /api/v1/todos/{id}```

##### Request Headers
| Name        | Value         | Required  |
|:------------|:-------------|:----------|
|Authorization|Bearer {Token}| ```true```|

##### URL Parameters

| Name   | Type       | Required  | Description|
|:-------|:-----------|:----------|:---------- |
|id      |```Int```   | ```true```| ID of Todo |

#### Example

##### Request

```Coming soon...```

##### Response

```Coming soon...```

---

## Update Todo

Edit and save a todo to the database.

#### HTTP Request

```PATCH - /api/v1/todos/{id}```

##### Request Headers
| Name        | Value         | Required  |
|:------------|:-------------|:----------|
|Authorization|Bearer {Token}| ```true```|

##### URL Parameters

| Name   | Type       | Required  | Description|
|:-------|:-----------|:----------|:---------- |
|id      |```Int```   | ```true```| ID of Todo |

##### Request body

| Name   | Type       | Required  |
|:-------|:-----------|:----------|
|value   |```String```| ```true```|

#### Example

##### Request

```Coming soon...```

##### Response

```Coming soon...```

---

## Delete Todo

Remove your todo from the database

#### HTTP Request

```DELETE - /api/v1/todos/{id}```

##### Request Headers
| Name        | Value         | Required  |
|:------------|:-------------|:----------|
|Authorization|Bearer {Token}| ```true```|

##### URL Parameters

| Name   | Type       | Required  | Description|
|:-------|:-----------|:----------|:---------- |
|id      |```Int```   | ```true```| ID of Todo |

##### Request body

| Name   | Type       | Required  |
|:-------|:-----------|:----------|
|value   |```String```| ```true```|

#### Example

##### Request

```Coming soon...```

##### Response

```Coming soon...```

