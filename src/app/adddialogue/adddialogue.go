package adddialogue
import (
	"net/http"
	// "strings"
	"io/ioutil"
	"fmt"

	"github.com/rs/xid"

	"app/db"
	"encoding/json"
)

type Dialogue struct {

	Id string `json:"id"`
	Dialogue string `json:"dialogue"`
	Keywords []string `json:"keywords"`
	Character string `json:"character"`
	// Id map[string]interface {
	// 	Dialogue string `json:"dialogue"`
	// 	Keywords []string `json:"keywords"`
	// 	Character string `json:"character"`
	// }
}

type Profiel struct {
	Id struct {
		Name string `json:"name"`
    KnownAs string `json:"knownas"`
    Weapon []string `json:"Weapons"`
    Description string `json:"description"`
    Speciality []string `json:"speciality"`
    Defeated []string `json:"defeated"`
    SavedStone string `json:"savedstone"`
	} `json:"id"`

}

func AddDialogue (w http.ResponseWriter, r *http.Request){
  if r.Method == "POST" {
		id1 := xid.New()
		id := id1.String()
		fmt.Println("id->",id)
		var d Dialogue
		// dialogue := r.FormValue("dialogue")
		// keywords := r.FormValue("keywords")
		character := r.FormValue("character")
		fmt.Println(character)
		body, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &d)
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("body->", string(body))
		if err != nil {
			fmt.Println("canr read body", body)
		}
		// fmt.Println("char->", d.Character)
	 	db.HSet("h:"+d.Character, id, string(body))
   	fmt.Println("method->", r.Method)
 }
 if r.Method == "DELETE"{
 	  character := r.FormValue("character")
 	  id := r.FormValue("id")
 		db.HDel("h:"+character, id)

 }
 if r.Method == "GET" {
 		character := r.FormValue("character")
 		// var p  Dialogue
 		// dialogue, err := json.Marshal(p)
 		fmt.Println("char->", character)
 		// fmt.Println(db.HGetAll("h:"+character))
 		// w.Header().Set("Content-Type", "application/json")
 		// json.NewEncoder(w).Encode(p)\
 		map1, err := db.HGetAll("h:"+character)
 		if err != nil {
 			fmt.Println("err getting all dialogue", err)
 		}
 		fmt.Println("map->", map1)
 		body, err := json.Marshal(map1)
 		if err != nil{
 			fmt.Println("error in marshalling", err)
 		}
 		w.Header().Set("Content-Type","application/json")
 		w.WriteHeader(http.StatusOK)
 		w.Write(body)

 }
}

func AddProfiel(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {
		bodyObj, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("err in reading body", err)
		}
		var p Profiel
		err = json.Unmarshal(bodyObj, &p)
		if err != nil {
			fmt.Println("cant unmarshal bodyObj in profiel", err)
		}
		// p.Id = p.Id.Name+p.Id.KnownAs
		fmt.Println("p-->", p)
	}
}
