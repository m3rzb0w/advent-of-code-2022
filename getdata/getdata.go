package getdata

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Getdata(url string, sessionCookie string) string {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Cookie", sessionCookie)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[INFO] Req status : ", res.Status)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func Grabsession() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	session := os.Getenv("SESSION")
	fmt.Println("[INFO] Current cookie : ", session)
	return session
}
