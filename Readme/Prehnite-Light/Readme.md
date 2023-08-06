
## Prehnite-Light

This is a prototype made in NODEJS. As NodeJS is not statically typed it can be easier and quicker to make a working version of the system. This will act as a guide to build the final version in a more efficient language such as `GO` or `RUST`

The prototype doesn't have any security. Tokens and Security will be added at a later date.


### Prehnite
The aim of prehnite is to allow quick development of database systems. Quickly design a relational database in a GUI, add custom methods to modify data before adding to the database. This will allow quick updating of the database schema and automatically create the required databases, tables and records. 
The databases that are created can be accessed with a `API` that is updated with the schema of the database, allowing for development projects to quickly retrieve and store data to model how the system will act with the database model. The ability to have multiple types of databases including `Key-Value`, `Document` and `Relational` allows all required databases to be handled by this one program without the need to spin up multiple databases and servers.
After you have designed the system you can easily generate the required `SQL` code to generate the database. Connect Prehnite to your database, allow prehnite to act as a middleman to create the desired structure directly in your desired database software.

* At the moment no security - Not planned at the moment ( seemed boring and slightly more complicated than my `ADHD` would let me code. )
* Database objects are saved in manual variables at the moment. This will be updated when a Server API is added and a frontend GUI.

P.S ( Sorry I really like using Promises in JS, the use of .catch helps keep everything clean and condensed )

### Default File structure

```dirtree

- /Prehnite-Light
	- /node_modules
	- /Readme ( Obsidian Vault )
	- /Records ( Where the databases are stored by default)
		- /Database Example
			- Config.json
	- /src ( Source code )
		- /database
			- interface.js ( contains the DatabaseSystem )
			- databaseClass.js ( contains the class for the database )
			- `databaseControls.js (  )`
	- System.json ( config file for database System )#
	- main.js  ( Entry point for node app )

```