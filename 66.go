package main
import ("fmt"
		"net/http"
		"os"
		"io/ioutil"
		"log"
		"encoding/json")


		type Response struct {
			Name    string    `json:"name"`
			Pokemon []Pokemon `json:"pokemon_entries"`
		}
		
		// A Pokemon Struct to map every pokemon to.
		type Pokemon struct {
			EntryNo int            `json:"entry_number"`
			Species PokemonSpecies `json:"pokemon_species"`
		}
		
		// A struct to map our Pokemon's Species which includes it's name
		type PokemonSpecies struct {
			Name string `json:"name"`
		}		

func main(){
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil{
		fmt.Println("Unable to connect API endpoint...!")
		os.Exit(1)
	}
	responsedata, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println("Unable to Read the response from API..!")
		log.Fatal(err)
	}
	//fmt.Println(string(responsedata))
	var responseObject Response
	json.Unmarshal(responsedata, &responseObject)
	//fmt.Println(responseObject.Name)
	for i:=0;i<len(responseObject.Pokemon);i++{
		fmt.Println(responseObject.Name)
	}
}