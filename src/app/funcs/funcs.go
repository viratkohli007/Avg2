package funcs
import (
	"html/template"
	"net/http"
	"app/db"
	"github.com/gomodule/redigo/redis"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"github.com/alexedwards/scs"
)

type Regform struct{
	fname string `json:"fname"`
	lname string `json:"lname"`
	email string `json: "email"`
	password string `json:"password"`
}

type Jsonapi struct{
	Id int `json:"id"`
	UserId int `json:"userId"`
	Title string `json:"title"`
	Body string `json:"body"`
}

var sessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

func Api(){
	//var res2  []Jsonapi
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	// err =json.Unmarshal(responseData, &res2)
	// if err != nil {
	// 	fmt.Println("??>>",err)
	// }

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, responseData, "", " ")
	if error != nil {
		fmt.Println(">>>", error)
	}

	// for _, v := range res2 {
	// 	out, err := json.Marshal(v.Title)
	// 	if err != nil {
	// 		fmt.Println("??>>",err)
	// 	}
	// 	fmt.Println("\n\n\n>", string(out))

	// }
	 // fmt.Println(string(prettyJSON.Bytes()))
}

func Abc(w http.ResponseWriter, r *http.Request){
    // w.Write([]byte("Hahahahahahah"))
}

func Login(w http.ResponseWriter,  r *http.Request){
	fmt.Println(ping(db.Conn))

	session := sessionManager.Load(r)
	err_s := session.PutString(w, "authenticated", "true")
	if err_s != nil{
		fmt.Println("session was not set")
	}

	t, _ := template.ParseFiles("login.html")
	t.Execute(w,"")
}

func Registration(w http.ResponseWriter,  r *http.Request){
	//fmt.Println(insertreg(db.Conn))
	t, _ := template.ParseFiles("register.html")
	t.Execute(w, "")
}

func Welcome(w http.ResponseWriter, r *http.Request){

	session := sessionManager.Load(r)

	/*
		isMail = true
		isPw = true
	*/

	msg, err_s2 := session.GetString("authenticated")
	if err_s2 != nil {
		fmt.Println(err_s2)
	}
	fmt.Println("msg", msg)
	formEmail := r.FormValue("email")
	formPassword := r.FormValue("password")

	if formEmail == "" && formPassword == "" {
		// w.Write("enter valid ")
		http.Redirect(w,r, "/login", 307)

	}

	fmt.Println(db.HGet(formEmail, "email"))
	fmt.Println(db.HGet(formEmail, "password"))
	email, err := db.HGet(formEmail, "email")
	if err == redis.ErrNil {
		fmt.Println("email doesnt exists", err)
		// isMail = false

	}
	if err != nil {
		fmt.Println("Error :",err)

	}

	password, err := db.HGet(formEmail, "password")
	if err == nil {
		fmt.Println("email doesnt exists", err)
	}
	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println("Email", email)
	fmt.Println("formEmail", formEmail)
	fmt.Println("password", password)
	fmt.Println("formPassword", formPassword)

	if password == formPassword && email == formEmail && msg == "true"{
		t, _ := template.ParseFiles("welcome.html")
		t.Execute(w, "")
	}else{
		http.Redirect(w,r, "/login", 307)
	}
}

func AfterReg(w http.ResponseWriter, r *http.Request){
	var regformobj Regform
	regformobj.fname = r.FormValue("fname")
	regformobj.lname = r.FormValue("lname")
	regformobj.email = r.FormValue("email")
	regformobj.password = r.FormValue("password")

	// err2 := json.Unmarshal()

	err := db.HSet(regformobj.email, "firstname", regformobj.fname)
	if err != nil {
		fmt.Println("afterreg func>>>",err)
	}
	err = db.HSet(regformobj.email, "lastname", regformobj.lname)
	if err != nil {
		fmt.Println("afterreg func>>>",err)
	}
	err = db.HSet(regformobj.email, "email", regformobj.email)
	if err != nil {
		fmt.Println("afterreg func>>>",err)
	}
	err = db.HSet(regformobj.email, "password", regformobj.password)

	http.Redirect(w,r,"/login",307)
}

func ping(c redis.Conn) error{
	pong, err := c.Do("Ping")
	if err != nil {
		return err
	}
	s,err := redis.String(pong,err)
	if err != nil {
		return err
	}
	fmt.Println(">>>",s)
	return nil
}

func Logout(w http.ResponseWriter, r *http.Request){
	session := sessionManager.Load(r)
	session.Remove(w, "authenticated")


	msg, err_s2 := session.GetString("authenticated")
	if err_s2 != nil {
		fmt.Println(err_s2)
	}
	fmt.Println("logout msg", msg)
	http.Redirect(w, r, "/login", 307)
}

// func insertreg (c redis.Conn) error{

// 	res, err := c.Do("SET user qqq")
// 	if err != nil {
// 		return err
// 	}
// 	s, err := redis.String(res,err)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(">>>", s)
// 	return nil

// }
