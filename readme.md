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
#Note for testing
Before running any tests, be it with `go test` or `goconvey`, be sure to export the `library_app_root` environment variable first. 

If you're in the project directory, simply: ``export library_app_root=`pwd` && goconvey``

It should be set to the root of this application, where this very MD file is located (e.g: /home/YourUsername/go/src/yellowroad_library )

--------
More details (like a description of what the app actually is) to come ... 