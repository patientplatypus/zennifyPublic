package util

import (
	// "crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	// "bytes"
	// "runtime"
	// "mime"
	"net/http"
	"strings"
	"os"
	"path/filepath"
	// "path"
	"github.com/gorilla/mux"
)

const maxUploadSize = 999999 * 1024 // 2 mb
const uploadPath = "./files"


func ServeFiles(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside ServeFiles")
	vars := mux.Vars(req)
	fmt.Println("%v", vars["filename"])

	files, err := ioutil.ReadDir("/go/src/github.com/patientplatypus/webserver/files")
    if err != nil {
        log.Fatal(err)
	}
	
	fileloc:=""

    for _, f := range files {
			name := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			if name == vars["filename"]{
				fileloc = "/go/src/github.com/patientplatypus/webserver/files"+"/"+vars["filename"] + filepath.Ext(f.Name())
				break
			}
	}

	fmt.Println("extension before if: ", filepath.Ext(files[0].Name()))
	
	if vars["filename"]=="add"{
		fmt.Println("inside wasm if statement~")
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, req, fileloc)
		// w.Write([]byte(fileloc))
	}else{
		http.ServeFile(w, req, fileloc)
	}
}

func UploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside uploadFileHandler")

		// validate file size
		fmt.Println("got to -2")
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		fmt.Println("got to -1")
		err := r.ParseMultipartForm(maxUploadSize);
		fmt.Println("got to 0")
		if err != nil {
			fmt.Println("inside err not nil for UploadFileHandler")
			log.Panic(err)
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		fmt.Println("got to 1")

		// parse and validate file and post parameters

		fmt.Println("got to 1.1")
		file, _, err := r.FormFile("uploadFile")
		fmt.Println("got to 1.2")
		fmt.Println("file: ", file)
		if err != nil {
			log.Panic(err)
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}


		fmt.Println("got to 2")

		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		specFileType := ""
		switch filetype {
		case "image/jpeg":
			specFileType = "jpeg"
		case "image/jpg":
			specFileType = "jpg"
		case "image/gif":
			specFileType = "gif"
		case "image/png":
			specFileType = "png"
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}

		fmt.Println("got to 4")

		fileName := r.PostFormValue("name")
		fileEmail := r.PostFormValue("email")

		fmt.Println("got to 5")

		newPath := filepath.Join("/go/src/github.com/patientplatypus/webserver/files", fileEmail + "_" + fileName + "." + specFileType)
		fmt.Println("uploadPath: ", uploadPath)
		fmt.Println("fileName: ", fileName)
		fmt.Println("filetype: ", filetype)
		fmt.Println("newPath: ", newPath)
		

		fmt.Println("got to 6")

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			fmt.Println("error in newFile creation")
			log.Panic(err)
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}

		fmt.Println("got to 7")

		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}

		fmt.Println("got to 8")

		w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
