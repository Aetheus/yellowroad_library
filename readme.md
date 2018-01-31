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
#Unit Tests
Before running any tests, be it with `go test` or `goconvey`, be sure to export the `library_app_root` environment variable first. 

If you're in the project directory, simply: ``export library_app_root=`pwd` && goconvey``

It should be set to the root of this application, where this very MD file is located (e.g: /home/YourUsername/go/src/yellowroad_library )

--------
#Migrations
##How to run them
In order to run migrations, simply execute `go run migrate.go`. The migrate tool will expect a config file to be present (see `_sample_config.json`), and this config file can be specified using the `config_path` environment variable. 

In the absence of such an environment variable, the tool will attempt to find a `config.json` file in the same directory that it is being executed from (i.e: the project's root directory).


##Adding new migrations
Database migrations should be stored in the `database/migrations` directory. They should follow this naming convention:
    
    `{YYYY}{MM}{DD}{HH}{MM}_{description}.up.sql`

For instance, a migration written on the 2nd of October 2017 at exactly 12:06AM should have a name like the following:
    
    `201710020006_add_requirements_and_effect_to_chapter_path.up.sql`


--------


More details (like a description of what the app actually is) to come ... 