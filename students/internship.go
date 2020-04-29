package students

import (
	"errors"
	"time"
)

type Internship interface {
	Add(obj *Student) (*Student,error)
	Get()([]*Student,error)
	GetById(id int64)(*Student,error)
	Update(obj *Student)(*Student,error)
	Delete(id int64) error
}
type internshipStruct struct {
	Repo Repository
}
func NewInternship(repo Repository) Internship{
	return &internshipStruct{
		Repo: repo,
	}
}
func(inter *internshipStruct) checkForEmpty(obj *Student) error{
	if obj.FirstName=="" {
		return errors.New("first Name is empty")
	}else if obj.LastName==""{
		return errors.New("last Name is empty")
	}else if obj.Username==""{
		return errors.New("username is empty")
	}else if obj.Password==""{
		return errors.New("password is empty")
	}else if obj.Email==""{
		return errors.New("email is empty")
	}else if obj.Phone==""{
		return errors.New("phone is empty")
	}

	return nil
}
func(inter *internshipStruct) checkForUsername(obj *Student) error{
	students,err:=inter.Repo.Get()
	if err!=nil{
		return err
	}
	for _,v := range students{
		if v.Username == obj.Username && v.Id!=obj.Id{
			return errors.New("user with that username already exist")
		}
	}
	return nil
}
func(inter *internshipStruct) checkForEmail(obj *Student) error {
	students,err:=inter.Repo.Get()
	if err!=nil{
		return err
	}
	for _,v := range students{
		if v.Email == obj.Email{
			return errors.New("user with that email already exist")
		}
	}
	return nil
}
func(inter *internshipStruct) Add(obj *Student) (*Student,error){
	err:=inter.checkForEmpty(obj)
	if err!=nil{
		return nil, err
	}
	err=inter.checkForUsername(obj)
	if err!=nil{
		return nil, err
	}
	err=inter.checkForEmail(obj)
	if err!=nil{
		return nil, err
	}
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()
	newobj,err:=inter.Repo.Add(obj)
	return newobj,err
}
func(inter *internshipStruct) Get()([]*Student,error){
	objects,err:=inter.Repo.Get()
	if err!=nil{
		return nil, err
	}
	return objects,err
}
func(inter *internshipStruct) Update(obj *Student)(*Student,error){
	err:=inter.checkForUsername(obj)
	if err!=nil{
		return nil, err
	}
	obj,err=inter.Repo.Update(obj)
	if err!=nil{
		return nil,err
	}
	return obj,err
}
func(inter *internshipStruct) GetById(id int64)(*Student,error){
	obj,err:=inter.Repo.GetById(id)
	if err!=nil{
		return nil,err
	}
	return obj,err
}
func(inter *internshipStruct) Delete(id int64) error{
	obj,err:=inter.Repo.GetById(id)
	if err!=nil{
		return err
	}
	err = inter.Repo.Delete(obj)
	if err!=nil{
		return err
	}
	return nil
}