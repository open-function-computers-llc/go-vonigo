package vonigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// NoteResponse - This is the shape of the response vonigo sends when attempting to create a note.
type NoteResponse struct {
	Company       string `json:"company"`
	SecurityToken string `json:"securityToken"`
	ErrNo         int    `json:"errNo"`
	ErrMsg        string `json:"errMsg"`
	DateNow       string `json:"dateNow"`
	Note          Note   `json:"Notes"`
}

// Note this is an object that help save string notes in relation to Vonigo Clients(Accounts/Leads).
type Note struct {
	ObjectID       string `json:"objectID"`
	Name           string `json:"name"`
	Title          string `json:"title"`
	DateCreated    string `json:"dateCreated"`
	DateLastEdited string `json:"dateLastEdited"`
	FieldID        string `json:"fieldID"`
	FieldValue     string `json:"fieldValue"`
	OptionID       string `json:"optionID"`
}

// CreateNote - Create a note for a single client(Account/Lead agnostic)
func CreateNote(clientID string, noteString string) (Note, error) {
	note := Note{}
	noteResponse := NoteResponse{}
	client := &http.Client{}

	if len(clientID) == 0 || len(noteString) == 0 {
		return note, errors.New("Paraemters empty when creating a note")
	}

	log.Info("Creating note")

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return note, err
		}
	}

	params, _ := getBaseParams("create") // this used to work. vonigo is terrible and the mode needs to be set in each individual place now
	params["method"] = "3"
	fields, err := createFields("9291", noteString)

	if err != nil {
		return note, err
	}
	params["clientID"] = clientID
	params["Fields"] = fields

	reqURL, err := buildURL(baseURL, "api/v1/data/Notes")
	if err != nil {
		return note, err
	}
	log.Info(reqURL)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(params)

	req, _ := http.NewRequest("POST", reqURL, buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	err = checkVonigoError(body)

	if err != nil {
		return note, err
	}

	err = json.Unmarshal(body, &noteResponse)
	if err != nil {
		return note, err
	}

	note = noteResponse.Note
	return note, nil
}
