/**
 * 
 * @param {String} name 
 * @param {String} rootPath 
 * @returns {Database} databaseObject
 */
function CreateDatabase (name,  rootPath) {
    let id = uuid.v4().toString()
    let data = new Database(id, name, rootPath);
    data.Write()
    return data;
} 
/**
 * 
 * @param {String} id 
 * @param {String} optional_Root 
 * @returns {Promise}
 */
function SelectDatabase (id, RootPath) {
    return new Promise((resolve, reject) => {


    // Check if optional_Root
    let rootDir = "";
    if (RootPath != undefined && RootPath != "") {
        rootDir = RootPath;
    }
    rootDir = path.join(rootDir, id);
    // Find the file in system

    if (!fs.existsSync(rootDir) || !fs.statSync(rootDir).isDirectory() ) {
            throw new Error(`Couldn't find database at "${rootDir}`)
        }
    

    // Read Config file
    if (!fs.existsSync(path.join(rootDir, "Config.json"))) {
        throw new Error(`Couldn't find config for database at ${path.join(rootDir, "Config.json")}`)
    }

    fs.readFile(path.join(rootDir, "Config.json"), "utf-8", (err, data) => {
            console.log(err, data);
        try {

            data = JSON.parse(data)
            
            
            dbo = new Database(data.id, data.name, RootPath)
            
            
            resolve(dbo);
        } catch (err) {
            reject(err)
        }

        })
    });
}
