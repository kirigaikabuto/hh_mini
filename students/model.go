package students

import "time"
type Repository interface {
	Add(obj *Student) (*Student,error)
	GetById(id int64)(*Student,error)
	Get()([]*Student,error)
	Delete(obj *Student) error
	Update(obj *Student) (*Student,error)
}
type Student struct {
	Id        int64  `json:"id" pg:"id,pk"`
	FirstName string `json:"first_name,omitempty" pg:"first_name"`
	LastName string `json:"last_name,omitempty" pg:"last_name"`
	Username string `json:"username,omitempty" pg:"username"`
	Password string `json:"password,omitempty" pg:"password"`
	Email string `json:"email,omitempty" pg:"email"`
	Phone string `json:"phone,omitempty" pg:"phone"`
	CreatedAt time.Time `json:"created_at,omitempty" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" pg:"updated_at"`
}