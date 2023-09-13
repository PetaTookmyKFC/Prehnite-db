package server

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

func (ser *Server) Static(WebRoute string, fileSource string) {
	err := ser.Static_HandleDir(WebRoute, fileSource)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// routeRoot - request path e.g localhost/index.html
// folder - The location of the folder within the file system
func (ser *Server) Static_HandleDir(routeRoot string, folder string) error {
	files, err := os.ReadDir(folder)
	if err != nil {
		return err
	}

	files = []fs.DirEntry(files)

	for _, file := range files {
		if file.IsDir() {
			ser.Static_HandleDir(path.Join(routeRoot, file.Name()), path.Join(folder, file.Name()))
		} else {
			// Is file index
			if NameWithoutExt(file.Name()) == "index" {
				if path.Ext(file.Name()) == ".html" {
					if routeRoot == "/" {
						ser.Get(path.Join(routeRoot, "/"), staticFile(path.Join(folder, file.Name())))
						ser.Get(path.Join(routeRoot, "/")+"/", staticFile(path.Join(folder, file.Name())))
					} else {
						ser.Get(path.Join(routeRoot, "/")+"/", staticFile(path.Join(folder, file.Name())))
						fmt.Printf("GET: %s, FILE: %s \n", path.Join(routeRoot, "/")+"/", path.Join(folder, file.Name()))
					}
				} else {
					ser.Get(path.Join(routeRoot, file.Name()), staticFile(path.Join(folder, file.Name())))
					fmt.Printf("GET: %s, FILE: %s \n", path.Join(routeRoot, file.Name()), path.Join(folder, file.Name()))
				}
			} else {
				if path.Ext(file.Name()) == ".html" {
					ser.Get(path.Join(routeRoot, NameWithoutExt(file.Name())), staticFile(path.Join(folder, file.Name())))
					fmt.Printf("GET: %s, FILE: %s \n", path.Join(routeRoot, "/"), path.Join(folder, file.Name()))
				} else {
					ser.Get(path.Join(routeRoot, file.Name()), staticFile(path.Join(folder, file.Name())))
					fmt.Printf("GET: %s, FILE: %s \n", path.Join(routeRoot, file.Name()), path.Join(folder, file.Name()))
				}
			}
		}
	}
	return nil
}

func staticFile(path string) http.HandlerFunc {
	function := func(res http.ResponseWriter, req *http.Request) {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			res.WriteHeader(404)
			return
		}

		filebytes, err := os.ReadFile(path)
		if err != nil {
			res.Write([]byte(err.Error()))
			res.WriteHeader(500)
			return
		}

		res.Write(filebytes)
		res.WriteHeader(200)
	}
	return function
}

func NameWithoutExt(name string) string {
	return strings.Replace(name, path.Ext(name), "", 1)
}
