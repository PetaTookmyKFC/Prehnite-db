const fs = require("fs")
const path = require("path")

function GenerateDatabaseInfo (configPath) {
    return new Promise((resolve, reject) => {

        // Does that config exist?
        if (!fs.existsSync(configPath)) {
            reject("File Not Found")
            return
        }
        if (!fs.statSync(configPath).isFile())
        {
            reject("File Not Found")
            return 
        }

        fs.readFile(configPath, "utf-8", (err, data) => {
            if (err != undefined) {
                reject(err);
                return
            }

            let pathInfo = path.dirname(configPath)
            let name = path.basename(pathInfo)
            let root= path.dirname(pathInfo)
            
        })

    })
}


module.exports = {GenerateDatabaseInfo}