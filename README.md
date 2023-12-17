# golang
project
This is basic CRUD app, whihc uses to GO/Go Lang as web server. MySql is used as Databse to store records and for Front end I am using Plain old AngularJs.

Installtion
Below command will Install all the dependencies recursively.

go get -d ./...
Starting the GO server
Below command will start the GO server.

go run *.go

API endpoints
The following endpoints are available:

GET /movies : returns a list of all movies in the store

GET /movies/{id} : returns a single movie with the specified ID

POST /movies : creates a new movie based on the data sent in the request body

PUT /movies/{id} : updates an existing movie with the specified ID based on the data sent in the request body

DELETE /movies/{id} : deletes the movie with the specified ID from the store

Each movie has the following fields:

ID : a unique ID for the movie
Isbn : the ISBN for the movie
Title : the title of the movie
Director : the director of the movie (first name and last name)
Example usage
Here is an example of how to use the API using curl:

# Add a new movie
$ curl -X POST -H "Content-Type: application/json" -d '{"isbn":"448743","title":"Movie 1","director":{"firstname":"Director","lastname":"One"}}' http://localhost:8000/movies

# Get a list of all movies
$ curl http://localhost:8000/movies

# Get a single movie
$ curl http://localhost:8000/movies/1

# Update a movie
$ curl -X PUT -H "Content-Type: application/json" -d '{"isbn":"448744","title":"Updated Movie 1","director":{"firstname":"Director","lastname":"One"}}' http://localhost:8000/movies/1

# Delete a movie
$ curl -X DELETE http://localhost:8000/movies/1
