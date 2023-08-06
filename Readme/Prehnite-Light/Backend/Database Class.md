

This is the main object that contains the functions to handle tables within the database. This object handles the folder the records are stored in. 

Constructor 
```js
DB = new Database("Database name", "Path to where database is stored")
```

|fucntion|Params|Description|returns|
|-|-|-|-|
| MkRoot| N/A | Creates the folder used to contain the database folder if it doesn't already exist | N/A |
|Write| N/A | Writes the config file for the database | Promise() resolve: undefined, reject : error |


### Properties

| PropertyName | Description |
|-|-|
| path | Returns the path to the database folder that contains the tables ( Root + Database.Name ) |
| name | The name of the database ( this is also the name of the folder that contains the database ) |
| Root | The root path where the folder containing the database is stored |]
