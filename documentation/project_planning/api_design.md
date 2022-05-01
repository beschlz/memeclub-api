# Design Draft API Endpoints


## Version

`GET /api/version`

## Posts

`GET /api/posts`

`POST /api/posts`

### Post

`GET /api/posts/:id`

`PUT /api/posts/:id`

`DELETE /api/posts/:id`

#### Reactions

`GET /api/posts/:id/reactions`

`POST /api/posts/:id/reactions/:userID`

`DELETE /api/posts/:id/reactions/:userID`

### Comments

`GET /api/posts/:id/comments`

`POST /api/posts/:id/comments`

#### Comment

`PUT /api/posts/:id/comments/:id`

`DELETE /api/posts/:id/comments/:id`

`GET /api/posts/:id/comments/:id`

`PUT /api/posts/:id/comments/:id`

## Users

`GET /api/users/:id`

`PUT /api/users/:id`

`DELTE /api/users/:id`

## Auth

`POST /api/auth`