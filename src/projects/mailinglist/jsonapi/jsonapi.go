package jsonapi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mailinglist/mdb"
	"net/http"
)

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func fromJson[T any](body io.Reader, target T) {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(body)
	json.Unmarshal(buffer.Bytes(), &target)
}

func returnJson[T any](w http.ResponseWriter, withData func() (T, error)) {
	setJsonHeader(w)

	data, err := withData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		serverErrorJson, err := json.Marshal(&err)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(serverErrorJson)
		return
	}

	dataJson, err := json.Marshal(&data)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(dataJson)
}

func returnErr(w http.ResponseWriter, err error, code int) {
	returnJson(w, func() (interface{}, error) {
		errorMessage := struct {
			Err string
		}{
			Err: err.Error(),
		}
		w.WriteHeader(code)
		return errorMessage, nil
	})
}

func CreateEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			return
		}

		entry := mdb.EmailEntry{}
		fromJson(r.Body, &entry)

		if err := mdb.CreateEmail(db, entry.Email); err != nil {
			returnErr(w, err, http.StatusBadRequest)
			return
		}

		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON: Created email entry: %v\n", entry)
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

func GetEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			return
		}

		entry := mdb.EmailEntry{}
		fromJson(r.Body, &entry)

		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON GetEmaily: %v\n", entry)
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

func GetEmailBatch(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			return
		}

		opts := mdb.GetEmailBatchQueryParams{}
		fromJson(r.Body, &opts)

		if opts.Count <= 0 || opts.Page <= 0 {
			returnErr(w, errors.New("invalid count or page"), http.StatusBadRequest)
		}

		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON GetEmailBatch: %v\n", opts)
			return mdb.GetEmailBatch(db, opts)
		})

	})
}

func UpdateEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPut {
			return
		}

		entry := mdb.EmailEntry{}
		fromJson(r.Body, &entry)

		if err := mdb.UpdateEmail(db, entry); err != nil {
			returnErr(w, err, http.StatusBadRequest)
			return
		}

		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON: Updated email entry: %v\n", entry)
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

func DeleteEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			return
		}

		entry := mdb.EmailEntry{}
		fromJson(r.Body, &entry)

		if err := mdb.DeleteEmail(db, entry.Email); err != nil {
			returnErr(w, err, http.StatusBadRequest)
			return
		}

		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON: Deleted email entry: %v\n", entry)
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

func Serve(db *sql.DB, bind string) {
	http.Handle("/email/create", CreateEmail(db))
	http.Handle("/email/get", GetEmail(db))
	http.Handle("/email/get_batch", GetEmailBatch(db))
	http.Handle("/email/update", UpdateEmail(db))
	http.Handle("/email/delete", DeleteEmail(db))

	log.Printf("Listening on %s\n", bind)
	err := http.ListenAndServe(bind, nil)

	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
