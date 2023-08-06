global.path = require("path");

const strParser = require("./src/parseString.js");
const DBSystem = require("./src/database/interface");
console.log("Hello World!");

global.Root = path.join(__dirname, "Records");

let Sys = new DBSystem("./System.json", Root);
Sys.CreateDatabase("Test Database");
