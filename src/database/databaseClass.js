const fs = require("fs");
const path = require("path");

class Database {
  /**
   * @param {String} name
   * @param {path} path ( to root )
   */
  constructor(name, RootPath) {
    this.name = name;
    this.Root = RootPath;
    this.MkRoot();

    // Load Config File! - If exists
    this.reloadConfig()
  }

  reloadConfig() {
    return new Promise((resolve, reject) => {
      this.MkRoot();

      let needMake = true;
      if (fs.existsSync(this.path)) {
        if (fs.statSync(this.path).isDirectory()) {
          needMake = false;
        }
      }

      if (needMake) {
        reject("Directory for database Doesn't exists");
        return;
      }

      fs.readFile(path.join(this.path, "Config.json"), "utf-8", (err, data) => {
        if (err != undefined) {
          reject(err);
          return;
        }

        data = JSON.parse(data);
        this.name = data.name;
        this.Root = data.rootDir;
      });
    });
  }

  get path() {
    return path.join(this.Root, this.name);
  }

  MkRoot() {
    let needMake = true;
    if (fs.existsSync(this.Root)) {
      if (fs.statSync(this.Root).isDirectory()) {
        needMake = false;
      }
    }
    if (needMake) {
      fs.mkdirSync(this.Root);
    }
  }
  Write() {
    return new Promise((resolve, reject) => {
      this.MkRoot();

      let needMake = true;
      if (fs.existsSync(this.path)) {
        if (fs.statSync(this.path).isDirectory()) {
          needMake = false;
        }
      }

      if (needMake) {
        fs.mkdirSync(this.path);
      }

      fs.writeFile(
        path.join(this.path, "Config.json"),
        JSON.stringify({
          name: this.name,
          rootDir: this.Root,
        }),
        "utf-8",
        (err) => {
          if (err != undefined) {
            reject(err);
            return;
          }
          resolve();
        }
      );
    });
  }
  LogData() {
    console.log(`Database : ${this.name} - ${this.path}`);
  }
}

module.exports = Database;
