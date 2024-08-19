package user

type UserRepository interface {
    Save(user User) (User, error)
    GetByID(id uint) (User, error)
    GetAll() ([]User, error)
    Update(id uint, updates map[string]interface{}) error
    Delete(id uint) error
	GetUserByEmail(email string) (User, error)
}
