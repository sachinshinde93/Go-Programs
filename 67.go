package main
//import "fmt"
import "net/http"
//import "strconv"
//import "sync"

//var count int
//var mutex = &sync.Mutex{}

/*func inc_func(w http.ResponseWriter, r* http.Request){
	mutex.Lock()
	count++
	fmt.Fprintf(w,"You are %vth user visiting this site..Thanks",strconv.Itoa(count))
	mutex.Unlock()
}*/

func main(){
	/*http.HandleFunc("/",func(w http.ResponseWriter, r* http.Request){
		fmt.Fprintf(w,"Hello to the New Project")
	})
	http.HandleFunc("/increment",inc_func)*/
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8090",nil)
}