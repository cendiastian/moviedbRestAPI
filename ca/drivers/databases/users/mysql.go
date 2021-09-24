package users

import (
	"context"
	"project/ca/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Connect *gorm.DB
}

func NewMysqlUserRepository(connect *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Connect: connect,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.User, error) {
	var user Users
	result := rep.Connect.First(&user, "email = ? AND password = ?",
		email, password)
	if result.Error != nil {
		return users.User{}, result.Error
	}
	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetAll(ctx context.Context) ([]users.User, error) {
	var user []Users
	result := rep.Connect.Find(&user)
	if result.Error != nil {
		return []users.User{}, result.Error
	}
	return ToListDomain(user), nil
}

func (rep *MysqlUserRepository) UserDetail(ctx context.Context, id int) (users.User, error) {
	var user Users
	result := rep.Connect.First(&user, "id= ?", id)
	if result.Error != nil {
		return users.User{}, result.Error
	}
	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Delete(ctx context.Context, id int) (users.User, error) {
	var user Users
	result := rep.Connect.Delete(&user, "id= ?", id)

	if result.Error != nil {
		return users.User{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Update(ctx context.Context, id int, email string, password string) (users.User, error) {
	var user Users
	result := rep.Connect.Where("id = ?", id).Updates(&Users{Email: email, Password: password})

	if result.Error != nil {
		return users.User{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Register(ctx context.Context, name string, email string, password string) (users.User, error) {
	var user Users
	result := rep.Connect.Create(&user)

	if result.Error != nil {
		return users.User{}, result.Error
	}

	return user.ToDomain(), nil
}
