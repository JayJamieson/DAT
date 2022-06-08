package pb_test

import (
	"bufio"
	"fmt"
	"mime"
	"net/mail"
	"os"
	"testing"

	"github.com/JayJamieson/pb"
)

func TestParseEmail(t *testing.T) {

	f, _ := os.Open("./testdata/pricebook.email.txt")
	defer f.Close()
	reader := bufio.NewReader(f)
	msg, err := mail.ReadMessage(reader)

	if err != nil {
		panic(err.Error())
	}

	header := msg.Header

	// body, _ := io.ReadAll(msg.Body)
	// var a = string(body)
	// fmt.Println(a)
	// Content-Type: multipart/mixed; boundary="_004_SY6PR01MB8522D9BEDE755259711E5DA7E9A49SY6PR01MB8522ausp_"
	// Retrieve the Content-Type header and parse it to get the boundary value
	_, params, _ := mime.ParseMediaType(msg.Header.Get("Content-Type"))

	// Instantiate a new MIME reader from the body of the message using the boundary
	// mreader := multipart.NewReader(msg.Body, params["boundary"])
	// part, err := mreader.NextPart()
	// fmt.Println(part.FileName())
	pb.ParsePart(msg.Body, params["boundary"])

	fmt.Println("Content-Type:", header.Get("Content-Type"))
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))
}
