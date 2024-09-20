# Gin OAuth2 secured

## cURL requests

### Books

- Get all Books

  ```bash
  curl --location 'localhost:8080/api/books'
  ```

- Create a Book

  ```bash
  curl --location 'localhost:8080/api/books' \
  --header 'Content-Type: application/json' \
  --data '{
  "title": "Sample Book Title",
  "author": "John Doe"
  }'
  ```

- Get a Book by ID

  ```bash
  curl --location 'localhost:8080/api/books/1'
  ```

- Update a Book

  ```bash
  curl --location --request PATCH 'localhost:8080/api/books/1' \
  --header 'Content-Type: application/json' \
  --data '{
  "title": "Updated Book Title",
  "author": "John Doe"
  }'
  ```
