package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	recaptcha "github.com/dpapathanasiou/go-recaptcha"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/sessions"
)

var err error
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func CheckGoogleCaptcha(response string) bool {
	var googleCaptcha string = "6Lc3XT4UAAAAABac5-cbX23gBDnzUd9_TUYNbWQF"
	req, err := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", nil)
	q := req.URL.Query()
	q.Add("secret", googleCaptcha)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	var googleResponse map[string]interface{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &googleResponse)
	return googleResponse["success"].(bool)
}

func mainpage(res http.ResponseWriter, req *http.Request) {
	//validate recaptcha; if false then reload page
	/*captcha := req.FormValue("g-recaptcha-response")
	valid := CheckGoogleCaptcha(captcha)

	if valid != true {
		http.ServeFile(res, req, "bad_captcha_login.html")
	}*/
	http.ServeFile(res, req, "mainpage.html")

}

func errorHandler(res http.ResponseWriter, req *http.Request, status int) {
	res.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(res, "404 Page Not Found")

		//For now, dont server any html
		//http.ServeFile(res, req, "404.html")
	}
}

func ediapi(res http.ResponseWriter, req *http.Request) {
	//validate recaptcha; if false then reload page
	/*captcha := req.FormValue("g-recaptcha-response")
	valid := CheckGoogleCaptcha(captcha)

	if valid != true {
		http.ServeFile(res, req, "bad_captcha_login.html")
	}

	*/

	// Check if user is authenticated
	/*session, _ := store.Get(req, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(res, "Forbidden", http.StatusForbidden)
		return
	}*/

	//just a basic Struct which will be converted into Json
	type User struct {
		Username  string
		Password  string `json:"-"`
		IsAdmin   bool
		CreatedAt time.Time
	}

	//If the request is not GET then redirect to homepage
	//if req.Method != "GET" {
	//	http.ServeFile(res, req, "EDIManager.html")

	//	return
	//}

	//user := User{} //initialize empty user

	//Marshal or convert user object back to json and write to response
	//userJson, err := json.Marshal(user)
	//if err != nil{
	//	panic(err)
	//}

	//Set Content-Type header so that clients will know how to read response
	//res.Header().Set("Content-Type","application/json")
	//res.WriteHeader(http.StatusOK)
	//Write json response back to response
	//res.Write(userJson)

	//Validate user credentials against sql database
	//username := req.FormValue("username")
	//password := req.FormValue("password")

	//var databaseUsername string
	//var databasePassword string

	//err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	/*if err != nil {
		//http.Redirect(res, req, "/index", 301)
		res.Write([]byte("Error, Can't process request"))
		return
	}*/

	//err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	//if err != nil {
	//	res.Write([]byte("Failed to Login"))
	//	return
	//}

	//Create Cookie
	//session, _ := store.Get(req, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	//session.Values["authenticated"] = true
	//session.Save(req, res)

	http.ServeFile(res, req, "media/test.json")

}

func robot(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "robots.txt")

}

func createhtmlguideline(res http.ResponseWriter, req *http.Request) {
	// Check if user is authenticated
	session, _ := store.Get(req, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(res, "Forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(res, req, "create_html_guideline.html")
}

//This is the starting page of the application
func start(res http.ResponseWriter, req *http.Request) {

	//validate recaptcha; if false then reload page
	/*captcha := req.FormValue("g-recaptcha-response")
	valid := CheckGoogleCaptcha(captcha)

	if valid != true {
		http.ServeFile(res, req, "captcha.html")
	}*/

	if req.Method != "POST" {
		//http.ServeFile(res, req, "login.html")
		res.Write([]byte("Error"))
		return
	}

	//no user validation!
	http.ServeFile(res, req, "main.html")

}

func makenewDirectory(style string, format string, org string, version string, name string) {

	//make dir if not exist
	//0755  = read, write execuite
	/*
		0777 Everyone can read write and execute. On a web server, it is not advisable to use ‘777’ permission for your files and folders, as it allows anyone to add malicious code to your server.

		0644 Only the owner can read and write. Everyone else can only read. No one can execute the file.

		0655 Only the owner can read and write, but not execute the file. Everyone else can read and execute, but cannot modify the file.
	*/
	//os.Mkdir("./"+path, 0755)
	fmt.Println("about to make directory")
	//os.Mkdir("./"+path, os.FileMode(0522))
	//err = os.Mkdir("./edi_guidelines/test3", 0755)
	fmt.Println(err)
	//path = "./" + path + "/"
	//os.Mkdir("\""+path+"\"", 0755)
	//err = os.MkdirAll(path, os.ModePerm)

	//,style,format,org,version,name

	os.MkdirAll("./edi_guidelines/"+style, os.ModePerm)
	os.MkdirAll("./edi_guidelines/"+style+"/"+format, os.ModePerm)
	os.MkdirAll("./edi_guidelines/"+style+"/"+format+"/"+org, os.ModePerm)
	os.MkdirAll("./edi_guidelines/"+style+"/"+format+"/"+org+"/"+version, os.ModePerm)
	os.MkdirAll("./edi_guidelines/"+style+"/"+format+"/"+org+"/"+version+"/"+name, os.ModePerm)

	//jsonResponse(res, http.StatusCreated, "File uploaded successfully!.")
	fmt.Println("end of file was saved...")
}

func saveFile(res http.ResponseWriter, file multipart.File, handle *multipart.FileHeader, path string, style string, format string, org string, version string, name string) {
	fmt.Println("in save file wtf")
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(res, "%v", err)
		return
	}

	//make dir if not exist
	//0755  = read, write execuite
	/*
		0777 Everyone can read write and execute. On a web server, it is not advisable to use ‘777’ permission for your files and folders, as it allows anyone to add malicious code to your server.

		0644 Only the owner can read and write. Everyone else can only read. No one can execute the file.

		0655 Only the owner can read and write, but not execute the file. Everyone else can read and execute, but cannot modify the file.
	*/
	//os.Mkdir("./"+path, 0755)
	fmt.Println("about to make damn directory")
	//os.Mkdir("./"+path, os.FileMode(0522))
	//err = os.Mkdir("./edi_guidelines/test3", 0755)
	fmt.Println(err)
	//path = "./" + path + "/"
	//os.Mkdir("\""+path+"\"", 0755)
	//err = os.MkdirAll(path, os.ModePerm)

	//,style,format,org,version,name

	os.MkdirAll("./edi_guidelines/"+style, os.ModePerm)
	os.MkdirAll("./edi_guidelines/"+style+"/"+format, os.ModePerm)
	os.MkdirAll("./edi_guidelines/"+style+"/"+format+"/"+org, os.ModePerm)
	os.MkdirAll("./edi_guidelines/"+style+"/"+format+"/"+org+"/"+version, os.ModePerm)
	os.MkdirAll("./edi_guidelines/"+style+"/"+format+"/"+org+"/"+version+"/"+name, os.ModePerm)

	//err = os.Mkdir("./edi_guidelines/kjh", 0755)
	path = "./edi_guidelines/" + style + "/" + format + "/" + org + "/" + version + "/" + name
	//fmt.Println("made directory: " + "\"" + path + "\"")

	err = ioutil.WriteFile(path+"/"+handle.Filename, data, 0666)
	fmt.Println("finished writing  wtf" + path + handle.Filename)
	if err != nil {
		fmt.Fprintf(res, "%v", err)
		return
	}
	//jsonResponse(res, http.StatusCreated, "File uploaded successfully!.")
	fmt.Println("end of file was saved...")
}

//nice function to send JSON responses!
// jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func RequestLogger(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// log request by who(IP address)
	requesterIP := r.RemoteAddr

	log.Printf(
		"%s\t\t%s\t\t%s\t\t%v",
		r.Method,
		r.RequestURI,
		requesterIP,

		time.Since(start),
	)

}

func googleSearchConsole(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "google259e7adf5a143f76.html")
}

func serviceworker(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "sw1.js")
}

func manifest(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "manifest.json")
}

func gcim(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "4041.html")
}
func oedi(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "4041.html")
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		errorHandler(res, req, http.StatusNotFound)

		//log ip
		//log.Println(req.RemoteAddr)

		return
	}
	log.Println("index was accesssed")
	// log request by who(IP address)
	start := time.Now()
	requesterIP := req.RemoteAddr
	log.Printf(
		"%s\t\t%s\t\t%s\t\t%v",
		req.Method,
		req.RequestURI,
		requesterIP,
		time.Since(start),
	)
	//end log
	log.Println("successfully served index!")

	http.ServeFile(res, req, "index.html")
}

//XML END

func jsonBeautifyPage(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "jsonbeautifypage.html")
}

func jsonbeautify(res http.ResponseWriter, req *http.Request) {

	in := req.FormValue("jsonDocumentTextBox1")
	//fmt.Println(in)
	//xy, _ := strconv.Unquote(input)

	input := json.RawMessage(in)

	x, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		panic(err)
	}
	res.Write([]byte(x))

	//fmt.Println(x)
}

func main() {

	//REad config file
	config()

	//Initialize captcha key
	recaptcha.Init("6Lc3XT4UAAAAABac5-cbX23gBDnzUd9_TUYNbWQF")
	//db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	//db, err = sql.Open("mysql", "appuser1:ZBhgdgemPvOH2nVubhQV@tcp(192.168.178.46:3306)/baseglobe") //Live: appuser1:ZBhgdgemPvOH2nVubhQV@/users TEST:"appuser1:yoyos4life@/erp"
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer db.Close()

	if err != nil {
		panic(err.Error())
	}
	log.Println("Listening...")

	http.HandleFunc("/", use(myHandler, basicAuth))
	//http.HandleFunc("/", index)
	//http.HandleFunc("/google259e7adf5a143f76.html", googleSearchConsole)
	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	//http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))
	//http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))
	//http.Handle("/edi_guidelines/", http.StripPrefix("/edi_guidelines/", http.FileServer(http.Dir("edi_guidelines"))))

	//log file system
	fileName := "webrequests.log"
	// https://www.socketloop.com/tutorials/golang-how-to-save-log-messages-to-file
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	// direct all log messages to webrequests.log
	log.SetOutput(logFile)

	// Start the HTTP server
	//write a small error catch

	/*if err := http.ListenAndServe(":80", nil); err != nil { //LIVE: 80
		log.Fatal("failed to start server", err)
	}*/

	// Start the HTTPS server in a goroutine
	if err := http.ListenAndServeTLS(":443", "edigenerator.com.pem", "edigenerator.key", nil); err != nil {
		log.Fatal("failed to start server", err)
	}

	// Cerbot Free SSL instruction: https://certbot.eff.org/lets-encrypt/windows-other
}

/* Author:  Caleb Crickette 2021

This basic template has a lot of ideas to quickly get working in Go which will serve easily HTTP/S website or RESTful API's.

Some good articles on Go:
https://stackoverflow.blog/2020/11/02/go-golang-learn-fast-programming-languages/
https://dev.to/techschoolguru/implement-restful-http-api-in-go-using-gin-4ap1
https://medium.com/helidon/can-java-microservices-be-as-fast-as-go-5ceb9a45d673

For all documentation visit www.golang.org

Inside there are examples of database inserts and reads as well as several examples of serving and serializing/unserializing JSON data.


HOW TO RUN AND COMPILE THE CODE TO AN .EXE (binary)
You need to have Go installed. (golang.org) then to test it open a terminal/cmd and type "go" you should get some information if not. Recheck installation.const
There are different ways to run Go code. While quickly developing it is better to just run the following command in your terminal or IDE:
"Go run whatever_the_name_of_your_program.go"
It will then run in your console.  If you want to create an executable and compile to machine code to deploy its very similar, but instead of RUN use "build":
"Go build whatever_the_name_of_your_program.go"
NOTE: if you want to deploy on other environments you need to set an environment variblae on your computer AND THEN run the  build command.
REMBMER to set it back if you are working on your local machine!

For  mac for example run the following code in your terminal console or cmd:
1. $ GOOS=darwin GOARCH=386
2. Go build whatever_the_name_of_your_program.go


Then if you want windows exe:
1. $ GOOS=windows GOARCH=amd64
2. Go build whatever_the_name_of_your_program.go

Sometimes Windows gives a problem so try the following:
1. $env:GOOS = "linux"
2. read variable: $env:GOOS
3. Go build whatever_the_name_of_your_program.go


For a complete list visit here:
https://github.com/ccrickette/Go/blob/master/ps_cross_compile.txt
https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04


TIP: if you do not want the console open just add the following to your run/build command:
-ldflags -H=windowsgui

*/