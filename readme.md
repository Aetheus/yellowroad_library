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


More details (like a description of what the app actually is) to come ... 