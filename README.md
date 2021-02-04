# imdb_example
## Run Project
 ~~~
 go run main.go
 ~~~

## Project uses Mongodb as database
1. Make sure to create database with name __IMDB__
2. Create 2 collections 1 is __user__ with admin user and other collection name is __movies__
 ~~~
  {"username":"admin","password":"admin", "role":"admin"}
 ~~~
3. Import the imdb.json file to __movies__ collection.


## endpoint description.
1. Get - http://localhost:3032/movies - Available to all
2. Post - http://localhost:3032/movies - requires authorization(bearer) token
3. Delete - http://localhost:3032/movies/{id} - requires authorization(bearer) token
4. Patch - http://localhost:3032/movies/{id} - requires authorization(bearer) token and allows only score and popularity field changes.
5. Post - http://localhost:3032/auth/login

Changes added
