package main

import (
"net/http"
	"log"
	"html/template"
		"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"fmt"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
 
func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "\\"))
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i]), nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request){

    t, err := template.ParseFiles("resources/html/index.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}

func main() {
	
	strPath, _:=GetCurrentPath()
	fmt.Printf("Current Path: %s\n", strPath)
	
	parentPath := getParentDirectory(strPath)
	
	fmt.Printf("Parent Path: %s\n", parentPath)	
	
    http.Handle("/bootstrap-4.3.1/", http.StripPrefix("/bootstrap-4.3.1/", http.FileServer(http.Dir(parentPath + "/resources/bootstrap-4.3.1"))))
    http.Handle("/jquery-3.3.1/js/", http.StripPrefix("/jquery-3.3.1/js/", http.FileServer(http.Dir(parentPath + "/resources/jquery-3.3.1"))))

    http.HandleFunc("/",IndexHandler)
    http.ListenAndServe(":8181", nil)

    
    
}