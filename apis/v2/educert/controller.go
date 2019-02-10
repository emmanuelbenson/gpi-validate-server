package educert

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/emmanuelbenson/gpi-validate-v2/utils"
)

// New for adding new educational certificate record
func New(w http.ResponseWriter, r *http.Request) {
	uid, _ := strconv.Atoi(r.FormValue("userId"))
	ec := &EducationalCertificate{
		UserID:    uid,
		FirstName: r.FormValue("firstName"),
		OtherName: r.FormValue("otherName"),
		LastName:  r.FormValue("lastName"),
		Title:     r.FormValue("title"),
		Type:      r.FormValue("type"),
		Document:  "doc",
	}

	// Validate all except document
	errs := ec.Validate()
	if len(errs) > 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(errs)
		return
	}

	// Check if document was submitted
	_, doc, err := r.FormFile("document")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		field := "Document"
		errorMsg := "require"
		docErr := utils.Error{}
		docErr[field] = errorMsg

		errs := append(errs, docErr)

		json.NewEncoder(w).Encode(errs)
		return
	}

	// Validate document type
	typeError := utils.ValidateFileExtension(doc)
	if typeError != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		errs = append(errs, typeError)
		json.NewEncoder(w).Encode(errs)
		return
	}

	// Validate document size
	sizeError := utils.ValidateFileSize(doc)
	if sizeError != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		errs = append(errs, sizeError)
		json.NewEncoder(w).Encode(errs)
		return
	}

	// File upload path
	uploadPath := "./static/uploads/"

	fileName, err := utils.WriteFile(doc, uploadPath)
	if err != nil {
		uploadErr := utils.Error{}

		uploadErr["document"] = "could not upload document"
		errs = append(errs, uploadErr)
		return
	}

	ec.Document = fileName

	errz := Create(ec)
	if len(errz) > 0 {
		log.Println(err)
		createErr := utils.Error{}
		createErr["error"] = "Internal server error"
		errs = append(errs, createErr)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs)
		return
	}

	success := utils.Success{}
	success["message"] = "Request sent."
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
	return
}

// All list users Educational Certificate check requests
func All(w http.ResponseWriter, r *http.Request) {
	user := struct {
		ID int `json:"id"`
	}{}

	param, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error reading request parametr")
		return
	}

	json.Unmarshal(param, &user)

	if user.ID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad request")
		return
	}

	allRequests, err := FetchAll(user.ID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allRequests)
	return
}

// Show for displaying a single educational certificate
func Show(w http.ResponseWriter, r *http.Request) {

}
