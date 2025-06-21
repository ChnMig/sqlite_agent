# API

## Get the connection status of db

``` bash
curl --request GET \
  --url http://192.168.124.65:10086/api/v1/info/status \
  --header 'auth: your_api_auth_key'
```

## Add a user table

``` bash
curl --request POST \
  --url http://192.168.124.65:10086/api/v1/exec \
  --header 'auth: your_api_auth_key' \
  --header 'content-type: application/json' \
  --data '{
  "sql": "CREATE TABLE users (\n    id INTEGER PRIMARY KEY AUTOINCREMENT,\n    name TEXT NOT NULL,\n    email TEXT UNIQUE\n);"
}
'
```

## Add a user to the user table

``` bash
curl --request POST \
  --url http://192.168.124.65:10086/api/v1/exec \
  --header 'auth: your_api_auth_key' \
  --header 'content-type: application/json' \
  --data '{
  "sql": "INSERT INTO users (name, email) VALUES ('\''Alice'\'', '\''alice@example.com'\'')"
}'
```

## Query the user table

``` bash
curl --request POST \
  --url http://192.168.124.65:10086/api/v1/exec/query \
  --header 'auth: your_api_auth_key' \
  --header 'content-type: application/json' \
  --data '{
  "sql": "SELECT * FROM users"
}'
```

## Delete a user from the user table

``` bash
curl --request POST \
  --url http://192.168.124.65:10086/api/v1/exec \
  --header 'auth: your_api_auth_key' \
  --header 'content-type: application/json' \
  --data '{
  "sql": "DELETE FROM users WHERE name = '\''Alice'\''"
}'
```

## Drop the user table

``` bash
curl --request POST \
  --url http://192.168.124.65:10086/api/v1/exec \
  --header 'auth: your_api_auth_key' \
  --header 'content-type: application/json' \
  --data '{
  "sql": "DROP TABLE users"
}'
```