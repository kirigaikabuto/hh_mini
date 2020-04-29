package students

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Endpoints interface{
	Add() func(w http.ResponseWriter,r *http.Request)
	Get() func(w http.ResponseWriter,r *http.Request)
	Update(idParam string) func(w http.ResponseWriter,r *http.Request)
	GetById(idParam string) func(w http.ResponseWriter, r *http.Request)
	Delete(idParam string) func(w http.ResponseWriter, r *http.Request)
}
type endpointsFactory struct {
	Inter Internship
}
func NewEndpoints(inter Internship) Endpoints{
	return &endpointsFactory{
		Inter: inter,
	}
}
type customError struct {
	ContextInfo string `json:"error"`
}
func(ef *endpointsFactory) Add() func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		obj:=&Student{}
		if err:=json.Unmarshal(data,&obj);err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		newobj,err:=ef.Inter.Add(obj)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,&customError{
				ContextInfo: err.Error(),
			})
			return
		}
		respondJSON(w,http.StatusOK,newobj)
	}
}

func(ef *endpointsFactory) Get() func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		objects,err:=ef.Inter.Get()
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,objects)
	}
}
func(ef *endpointsFactory) Update(idParam string)  func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars:=mux.Vars(r)
		paramid,paramerr:=vars[idParam]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,&customError{
				"no parameter",
			})
			return
		}
		id,err:=strconv.ParseInt(paramid,10,10)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		obj,err:=ef.Inter.GetById(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,&customError{
				"no object by id",
			})
			return
		}
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		if err:=json.Unmarshal(data,&obj);err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		updatedobj,err:=ef.Inter.Update(obj)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,&customError{
				err.Error(),
			})
			return
		}
		respondJSON(w,http.StatusOK,updatedobj)

	}
}
func(ef *endpointsFactory) GetById(idParam string)  func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramid, paramerr := vars[idParam]
		if !paramerr {
			respondJSON(w, http.StatusBadRequest, &customError{
				"no parameter",
			})
			return
		}
		id, err := strconv.ParseInt(paramid, 10, 10)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		obj, err := ef.Inter.GetById(id)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{
				"no object by id",
			})
			return
		}
		respondJSON(w, http.StatusOK,obj)
	}
}
func(ef *endpointsFactory) Delete(idParam string) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramid, paramerr := vars[idParam]
		if !paramerr {
			respondJSON(w, http.StatusBadRequest, &customError{
				"no parameter",
			})
			return
		}
		id, err := strconv.ParseInt(paramid, 10, 10)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		err = ef.Inter.Delete(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,&customError{
				err.Error(),
			})
		}
		respondJSON(w,http.StatusOK,"Element was Deleted")
	}
}
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}