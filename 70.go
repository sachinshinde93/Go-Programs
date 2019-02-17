package main
import "fmt"
import "os"
import "io/ioutil"

func main(){
	f,_ := ioutil.TempFile("/home/osboxes","tempfile-*.png")
	defer os.Remove(f.Name())
	fmt.Println(f.Name())
}
