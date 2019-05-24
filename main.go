package main
import (
	"net/http"
	"fmt"
	"app/db"
	"app/funcs"
	"app/adddialogue"
)

func main(){
	defer db.Conn.Close()
	fmt.Println(&db.Conn)
	fmt.Println("pool", db.Pool)
	const port = ":8080"
	funcs.Api()
	http.HandleFunc("/", funcs.Abc)
	http.HandleFunc("/login", funcs.Login)
	http.HandleFunc("/registration", funcs.Registration)
	http.HandleFunc("/welcome", funcs.Welcome)
	http.HandleFunc("/afterreg", funcs.AfterReg)
	http.HandleFunc("/logout", funcs.Logout)
	http.HandleFunc("/adddialogue", adddialogue.AddDialogue)
	http.HandleFunc("/addprofiel", adddialogue.AddProfiel)
	//http.HandleFunc("/api", funcs.Api)
	fmt.Println("server is running on port", port)
	http.ListenAndServe(port, nil)
}
