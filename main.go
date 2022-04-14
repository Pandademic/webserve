package main
 
import (
    "log"
    "net/http"
    "os"
)
var(
  port string
)
func main() {
    log.Println("info: starting")
    dir := os.Args[1]
    if(len(os.Args) < 1){
     log.Fatal("fatal: no directory provided\nUSAGE:\nwebserve directory port")
    }  
    log.Println("info running from dir:",dir)
    http.Handle("/", http.FileServer(http.Dir(string(dir))))
    if os.Args[2] == "" {
      port = ":5000"
    } else{
      port = ":"+ os.Args[2]
    }
    log.Println("info: starting on port ",os.Args[2])
    log.Fatal(http.ListenAndServe(port, nil))
}
