const fs = require("fs");
const path = require("path");



class Database {
    /**
     * @param {String} name
     * @param {path} path ( to root )
     */
    constructor (  name, RootPath ) {
        this.name = name;
        this.Root = RootPath;
    }
    get path () {
        return path.join(this.Root, this.name)
    }

    MkRootRoot () {
        let needMake = true;
        if ( fs.existsSync( this.Root ) ) {
            if ( fs.statSync(this.Root).isDirectory() ) {
                needMake = false;
            }
        }
        if (needMake) {
            fs.mkdirSync(this.Root)
        }
    }
    Write () {

        this.MkRootRoot()

        let needMake = true;
        if ( fs.existsSync( this.path ) ) {
            if ( fs.statSync( this.path ).isDirectory() ) {
                needMake = false;
            }
        }

        if (needMake) {
            fs.mkdirSync(this.path )
        }
        
        fs.writeFile(path.join(this.path, "Config.json"), JSON.stringify({
            name: this.name,
            rootDir: this.Root        
        }), "utf-8", (data) => {
            console.log(data);
        })
    }
    LogData () {
        console.log(`Database : ${this.name} - ${this.path}`)
    }
}

module.exports = Database;