package user

type UserService struct {
	UserRepo *UserRepo
}

// create function provider
func NewUserService(repo *UserRepo) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}
