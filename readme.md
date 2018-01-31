TODO: Clean this readme up. 

An unfinished webapp project (previously written in Node.js) that I'm reviving and rewriting from scratch in Go. Part of the goal is of course to finish the project itself - another part of the Go(al) is to master Golang itself. 

This is the API server of the app - the frontend will eventually be written in its own repo.

Current major libraries used :
* gin (Routing)
* gorm (ORM)

The project adopts a "[Route]-[Service]-[Repository]" architecture instead of traditional MVC:

Layer | Purpose
----- | -------
Route | Routes are responsible for **getting request parameters, passing them to the necessary Services/Repositories, and returning responses**.
Service | All "**business logic**" is handled by Services, which make use of Repositories to handle database interaction
Repository | All the nitty-gritty **database handling** is handled by Repositories, which take and return *Model structs*, which are POGS (plain old Go structs)


-------


# Unit Tests


This project uses GoConvey in order to run tests and expose them via a webpage on port 8080. In order to run these tests ... :

If you've `go build`'ed the executable:
    
    `$ [executable] test`
    
If you want it running quick-and-dirty using `go run`:
    
    `$ go run main.go test`

Note:
The tests internally use the `library_app_root` to define the root of the application. This will very probably change in future, but if you've set this environment variable for whatever reason, be warned of unexpected behaviour.

--------

# Migrations


## How to run them
If you've `go build`'ed the executable:
    
    `$ [executable] migrate`
    
If you want it running quick-and-dirty using `go run`:
    
    `$ go run main.go migrate`


## Adding new migrations
Database migrations should be stored in the `database/migrations` directory. They should follow this naming convention:
    
    `{YYYY}{MM}{DD}{HH}{MM}_{description}.up.sql`

For instance, a migration written on the 2nd of October 2017 at exactly 12:06AM should have a name like the following:
    
    `201710020006_add_requirements_and_effect_to_chapter_path.up.sql`


--------


More details (like a description of what the app actually is) to come ... 
