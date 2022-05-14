package main
 
import (
    "log"
    "net/http"
    "os"
)
var(
  port string
)
func main(){
    log.Println("info: init")
    dir := os.Args[1]
    log.Println("info: running from dir:",dir)
    http.Handle("/", http.FileServer(http.Dir(string(dir))))
    if os.Args[2] == ""{
      port = ":5000"
    }else{
      port = ":"+ os.Args[2]
    }
    log.Println("info: running on port:",port)

    log.Fatal(http.ListenAndServe(port, nil))
}
