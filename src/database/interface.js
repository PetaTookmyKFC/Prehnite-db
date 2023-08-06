const fs = require("fs");
const path = require("path");
const uuid = require("uuid");
const Database = require("./databaseClass.js");
const {GenerateDatabaseInfo} = require("./databaseControls.js");


class DBSystem {
  Databases = {};
  constructor(ConfigFile, DefaultRoot) {

    this.ConfigLocation = ConfigFile;
    this.DefaultRoot = DefaultRoot;
    
    if (fs.existsSync(ConfigFile)) {
      fs.readFile(ConfigFile, "utf-8", (err, data) => {
        if (err) throw err;

        let conf = JSON.parse(data);
        this.Databases = conf.Databases;
      });
    } else {
      fs.writeFileSync(this.ConfigLocation, JSON.stringify({ Databases: {}, DefaultRoot: this.DefaultRoot }), "utf-8");
      this.Databases = {};
    }
  }
  writeConfig() {
    return new Promise((resolve, reject) => {
      let data = JSON.stringify({Databases: this.Databases, DefaultRoot: this.DefaultRoot});
      fs.writeFile(this.ConfigLocation, data, (err) => {
        if (err) {
          reject(err);
        } else {
          resolve();
        }
      });
    });
  }
  /**
   *
   * @param {Database} database
  */
 addDatabase(database) {
   return new Promise((resolve, reject) => {
     let DBInfo = {
       name: database.name,
       Root: database.Root,
      };
      if (this.Databases[database.name] == undefined) {
        this.Databases[database.name] = DBInfo;
        this.writeConfig()
        .then((data) => {
          resolve(data);
        })
        .catch((err) => {
          reject(err);
        });
      } else {
        reject("Database name already exists");
      }
    });
  }
  /**
   *
   * @param {String} name
   * @param {String} rootPath
   * @returns {Database} databaseObject
  */
  CreateDatabase(name, OProotPath) {
   let rootPath = !(OProotPath == undefined || OProotPath == "") ? OProotPath : this.DefaultRoot
   
   return new Promise((resolve, reject) => {
     if (this.Databases[name] == undefined) {
       let data = new Database(name, rootPath);
       this.addDatabase(data).then(result => {
         data.Write();
         resolve()
        }).catch(err => reject(err))
      } else {
        reject("Name already exists")
      }
    });
  }
  /**
   * 
   * @param {String} name 
   * @returns {Database} databaseObject
   */
  SelectDatabase(name) {
    return new Promise((resolve, reject) => {
      if ( this.Databases[name] == undefined ) {
        reject("404");
      } else {
          let database = new Database(name, this.Databases[name].Root);
          resolve(database);
      }
    })
  }

  GenDatabase (configPath) {
    return GenerateDatabaseInfo(configPath)
  }
}

module.exports = DBSystem