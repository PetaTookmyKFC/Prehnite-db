
## Database-System

This is the main object, it contains all the databases that are registered in the system. The data for this object is saved and loaded from config.json

```JSON
{
	"Databases": {
		"DATABASE-NAME/ID" : { 
			"ROOT" : "PATH TO Folder that contains database folder"
		 }
	},
	"DefaultRoot" : "Default Location to store records"
}
```

---
 
constructor
```js
DBS = new DBSystem("Path to Config file", "default path for records to be stored")
```

### Functions

|fucntion|Params|Description|returns|
|-|-|-|-|
|writeConfig| N/A | writes config for database system| Promise() resolve: undefined, reject: error|
| addDatabase| database ( class Database )| Adds a database to the active memory, them calls `writeConfig`, if name is already taken rejects error | Promise() resolve: undefined, reject: error|
| CreateDatabase| name: String, OProotPatgh: String| Creates database with name if available, if unavailable rejects error. Creates database, runs `Databasesystem addDatabase`, writes database to records folder. | Promise() resolve: undefined, reject: error|
| SelectDatabase | name: String | Returns the database Object for the desired database if available | Promise () resolve: Database, reject: error |
|GenDatabase | configPath: String | Return a generated database object for the desired database, this is generated from the config file that the path leads to. The config file should be within the database folder. This allows for database to be modified even if not in the system | Promise (), resolve: Database, reject: error|
| RenameDatabase|name: String, newName: string| Renames the database - allows for the database to be easily renamed for adding into the system even if another database exists | Promise() resolve: undefined, reject: error|


### Properties

| PropertyName | Description |
|-|-|
| Databases | A Object that contains all the database objects |
| ConfigLocation|The path to the config file used by the database system |
| DefaultRoot| The path that is given to new created databases if no override is pressent |
