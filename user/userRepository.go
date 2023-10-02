package user

import "fmt"

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// method to get data by id
func (u *UserRepo) GetById(id int, chanResult chan string) {
	result := fmt.Sprintf("get data with id %v from database", id)
	chanResult <- result
}
