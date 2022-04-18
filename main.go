package main
 
import (
    "log"
    "net/http"
    "os"
)
var(
  port string
  portNum int
  dir string
)
var usage string = `
USAGE: webserve DIR PORT
   DIR is the directory you want to serve
   PORT is the port to serve it to
   if no PORT is provided , it will default serve to port 5000
`
func main() {
    log.Println("info: starting")
    portNum = os.Args[2]
    dir = os.Args[1]
    if(len(os.Args) < 2){
     log.Println("fatal: no directory provided")
     log.Fatal(usage)
    }  
    log.Println("info: running from dir:",dir)
    http.Handle("/", http.FileServer(http.Dir(string(dir))))
    if os.Args[2] == "" {
      port = ":5000"
    } else{
      port = ":"+ portNum
    }
    log.Println("info: starting on port ",portNum)
    log.Fatal(http.ListenAndServe(port, nil))
}
