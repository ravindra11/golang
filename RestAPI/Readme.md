Gin simplifies many coding tasks associated with building web applications, including web services.

You'll use `Gin` to route requests, retrieve request details, and marshal JSON for responses.

This includes the following sections.

1. Design API endpoints
2. Create a folder for your code
3. Create the data
4. Write a handler to return all items.
4. Write a handler to add a new item.
5. Write a handler to return a specific item.

# Design API endpoints

Build an API that provides access to a store selling vintage recordings on vinyl. 

So you'll need to provide endpoints through which a client can get and add albums for users.

When developing an API, you typically begin by designing the end point. Your API's users will have more success if the endpoints are easy to understand.

Here are the endpoints, you'll create in this tutorial

/albums

1. `GET` - Get a list of all albums, returned as JSON
2. `Post` - Add a new album from request data sent as JSON

/albums/:id

1. `GEt`- Get an album by its ID, returning the album data as JSON

# Create the data
To keep things simple, we'll store data in memory. A more typical API would interact with a database.

##### Note
Note that storing data in memory means that set of albums will be lost each time you stop the server, then recreated when you start it.

#### write the code
1. Create a file called `main.go`.
2. Into main.go, at the top of the file, paste the following package declaration.

```
package main

```

A standalone program (as opposed to a library)is always in package main.

3. Beneath the package declaration, declare the `album` struct. Will use this to store album data in memory.

`
// album represents data about a record album.
type Album struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Price float64 `json:"price"`
}
`

Struct tags such as `json:"artist"` specify what a field's name should be when the struct's contents are serialized into JSON.
Without them, the JSON would use the struct's capitalized field names - a style not as common in json.

4. Beneath the struct declaration you just added, albums array data.

### Write a handler to return all items

When the client makes a requirest at `GET /albums`, you want to return all the albums as JSON.

To do this, 
1. Logic to prepare a response
2. Code to map the request path to your logic

#### write the code
1. Beneath the struct code you added in the preceding section, add `getAlbums` writing the JSON into the response.

 `getAlbums` function would take `gin.Context` as parameter and it is the most important part of Gin. It carries request details, validates and serilizes JSON, and more.

Call `Context.IndentedJSON` to serilize the struct into JSON and add it to the response.
The function's first argument is the HTTP status code you want to send to the client. Here, you're passing the `StatusOK` constant from the `net/http` package to indicate 200 ok.

### Note ###
You can replace `Context.IndentedJSON` with a call to `Context.JSON` to send more compact JSON.
In practice, the indented form is much easier to work with when debugging and the size difference is usually small.

2. Neat the top of main.go, just beneath the albums slice declaration, declare a handler function to an endpoint path.

This sets up an association in which `getAlbums` handles requests to the `/albums` endpoint path.

`

func main(){

    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
}

`

### Run the code
1. Begin tracking the Gin module as a dependcy.
At the command line, use `go get` to add the github.com/gin-gonic/gin module as a dependency for your module.

Use a dot argument means `get dependencies for code in the current directory`

`
go get .

go get: added github.com/gin-gonic/gin v1.
`

2. from command line at main.go, run the code, use a dot argument to mean `run code in the current directory`

`go run .`

### Write a handler to add a new item(post)

When the client makes a POST request at /albums, you want to add the album described in the request body to the existing albums data.

