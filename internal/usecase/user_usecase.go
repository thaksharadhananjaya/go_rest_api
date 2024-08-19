package usecase

import (
    "restapi/pkg/errors"
    "restapi/internal/domain/user"
    "gorm.io/gorm"
)

type UserUseCase struct {
    userRepository user.UserRepository
}

func NewUserUseCase(repo user.UserRepository) *UserUseCase {
    return &UserUseCase{
        userRepository: repo,
    }
}

func (u *UserUseCase) CheckIfEmailExists(email string) error {
    _, err := u.userRepository.GetUserByEmail(email)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil
        }
        return errors.ErrInternalServerError
    }
    return errors.ErrBadRequest // Email already exists
}

func (u *UserUseCase) CreateUser(newUser user.User) (user.User, error) {
    
    if err := u.CheckIfEmailExists(newUser.Email); err != nil {
        return user.User{}, err
    }

    createdUser, err := u.userRepository.Save(newUser)
    if err != nil {
        return user.User{}, err
    }

    return createdUser, nil
}

func (u *UserUseCase) GetUserByID(id uint) (user.User, error) {
    usr, err := u.userRepository.GetByID(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return user.User{}, errors.ErrNotFound
        }
        return user.User{}, errors.ErrInternalServerError
    }
    return usr, nil
}

func (u *UserUseCase) GetAllUsers() ([]user.User, error) {
    users, err := u.userRepository.GetAll()
    if err != nil {
        return nil, errors.ErrInternalServerError
    }
    return users, nil
}

func (u *UserUseCase) UpdateUser(id uint, updates map[string]interface{}) (user.User, error) {
    _, err := u.userRepository.GetByID(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return user.User{},errors.ErrNotFound
        }
        return user.User{},errors.ErrInternalServerError
    }

    err = u.userRepository.Update(id, updates)
    if err != nil {
        return user.User{}, err
    }
    
    usr, err := u.userRepository.GetByID(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return user.User{}, errors.ErrNotFound
        }
        return user.User{}, errors.ErrInternalServerError
    }
    return usr, nil
}

func (u *UserUseCase) DeleteUser(id uint) error {
    if err := u.userRepository.Delete(id); err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.ErrNotFound
        }
        return errors.ErrInternalServerError
    }
    return nil
}
