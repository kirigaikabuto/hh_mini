package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"hh/config"
	"hh/students"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	PATH string  = ""
	conf config.PostgreConfig
)
var flags []cli.Flag=[]cli.Flag{
	&cli.StringFlag{
		Name: "config,c",
		Usage:"Load configuration from `FILE`",
		Destination: &PATH,
	},
	&cli.StringFlag{
		Name:"Host",
		Usage:"Postgre host",
		Destination: &conf.Host,
	},
	&cli.StringFlag{
		Name:"Port",
		Usage:"Postgre port",
		Destination: &conf.Port,
	},
	&cli.StringFlag{
		Name:"User",
		Usage:"Postgre user",
		Destination: &conf.User,
	},
	&cli.StringFlag{
		Name:"Password",
		Usage: "Postgre password",
		Destination: &conf.Password,
	},
	&cli.StringFlag{
		Name:"Database",
		Usage:"Postgre database",
		Destination: &conf.Database,
	},
}
func main(){
	app :=cli.NewApp()
	app.Name = "HH Rest Api"
	app.Flags = flags
	app.Action = runRestApi
	fmt.Println(app.Run(os.Args))
}


func extractConfigPostgre(path string, conf *config.PostgreConfig) error{
	file, _ := ioutil.ReadFile(path)
	err:=json.Unmarshal(file, &conf)
	if err!=nil{
		return err
	}
	return nil
}
func runRestApi(*cli.Context) error{
	if PATH!=""{
		err:=extractConfigPostgre(PATH,&conf)
		if err!=nil{
			return err
		}
	} else if conf.Host=="" {
		return errors.New("Nothing found in host")
	} else if conf.Port==""{
		return errors.New("Nothing found in port")
	} else if conf.User==""{
		return errors.New("Nothing found in user")
	} else if conf.Host ==""{
		return errors.New("Nothing found in host")
	} else if conf.Database ==""{
		return errors.New("Nothing found in database")
	}
	router:=mux.NewRouter()
	studentsrepo,err:=students.NewPostgreStore(conf)
	if err!=nil{
		return err
	}
	studentsinter:=students.NewInternship(studentsrepo)
	studentsendpoints:=students.NewEndpoints(studentsinter)
	router.Methods("POST").Path("/students/").HandlerFunc(studentsendpoints.Add())
	router.Methods("GET").Path("/students/").HandlerFunc(studentsendpoints.Get())
	router.Methods("GET").Path("/students/{id}").HandlerFunc(studentsendpoints.GetById("id"))
	router.Methods("PUT").Path("/students/{id}").HandlerFunc(studentsendpoints.Update("id"))
	router.Methods("DELETE").Path("/students/{id}").HandlerFunc(studentsendpoints.Delete("id"))
	fmt.Println("Server is running")
	err = http.ListenAndServe(":8000",router)
	if err!=nil{
		return err
	}
	return nil
}