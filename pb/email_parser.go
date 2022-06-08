package pb

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"
)

func ParsePart(mime_data io.Reader, boundary string) {
	// Instantiate a new io.Reader dedicated to MIME multipart parsing
	// using multipart.NewReader()
	reader := multipart.NewReader(mime_data, boundary)
	if reader == nil {
		return
	}

	// Go through each of the MIME part of the message Body with NextPart(),
	for {

		new_part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error going through the MIME parts -", err)
			break
		}

		mediaType, params, err := mime.ParseMediaType(new_part.Header.Get("Content-Type"))
		fmt.Println("Content-Type", mediaType)
		if err == nil && strings.HasPrefix(mediaType, "multipart/") {

			// This is a new multipart to be handled recursively
			ParsePart(new_part, params["boundary"])

		} else {

			fmt.Println("filename", new_part.FileName())
			fmt.Println("content disposition", new_part.Header.Get("Content-Disposition"))
			// Not a new nested multipart.
			// We can do something here with the data of this single MIME part.

		}

	}

}
