package users

import (
	"context"
	"fmt"
	"project/business/users"

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

func (rep *MysqlUserRepository) Login(ctx context.Context, domain users.User) (users.User, error) {
	var user Users
	result := rep.Connect.Preload("Premium").Preload("Transaction").First(&user, "email = ? ", domain.Email)
	if result.Error != nil {
		fmt.Println(result.Error)
		return users.User{}, result.Error
	}
	return user.ToDomainUser(), nil
}

func (rep *MysqlUserRepository) GetAll(ctx context.Context) ([]users.User, error) {
	var user []Users
	result := rep.Connect.Preload("Premium").Find(&user)
	if result.Error != nil {
		return []users.User{}, result.Error
	}
	return ToListDomain(user), nil
}

func (rep *MysqlUserRepository) UserDetail(ctx context.Context, id int) (users.User, error) {
	var user Users
	result := rep.Connect.Preload("Transaction.Payment_method").Preload("Transaction.Subscription_Plan").Preload("Premium").First(&user, "id= ?", id)
	if result.Error != nil {
		return users.User{}, result.Error
	}
	return user.ToDomainUser(), nil
}

func (rep *MysqlUserRepository) Delete(ctx context.Context, id int) (users.User, error) {
	var user Users
	result := rep.Connect.Delete(&user, "id= ?", id)

	if result.Error != nil {
		return users.User{}, result.Error
	}

	return user.ToDomainUser(), nil
}

func (rep *MysqlUserRepository) Update(ctx context.Context, domain users.User) (users.User, error) {
	user := FromDomain(domain)
	result := rep.Connect.Where("id = ?", user.Id).Updates(&Users{Email: user.Email, Password: user.Password})

	if result.Error != nil {
		return users.User{}, result.Error
	}

	return user.ToDomainUser(), nil
}

func (rep *MysqlUserRepository) Register(ctx context.Context, domain users.User) (users.User, error) {
	user := FromDomain(domain)

	// hashedPassword, err := encrypt.Hash(domain.Password)
	// if err != nil {
	// 	return users.User{}, err
	// }

	// user.Password = hashedPassword

	result := rep.Connect.Create(&user)

	if result.Error != nil {
		return users.User{}, result.Error
	}

	return user.ToDomainUser(), nil
}
