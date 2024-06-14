package repo

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// User repo -> ur
func (ur *UserRepo) GetInfoUser() string {
	return "Success!"
}
