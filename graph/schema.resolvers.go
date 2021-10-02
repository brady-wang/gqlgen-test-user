package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"test9/graph/generated"
	"test9/graph/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:34561)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var userInfo = &model.User{
		Name:   user.Name,
		Gender: user.Gender,
		Hobby:  user.Hobby,
	}
	db.Debug().Table("user").Create(&userInfo) // 创建user

	return userInfo, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, user model.UserInput) (*model.User, error) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:34561)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var userInfo = &model.User{
		Name:   user.Name,
		Gender: user.Gender,
		Hobby:  user.Hobby,
	}

	db.Debug().Table("user").Where("id=?", id).Update(userInfo)
	userInfo.ID, _ = strconv.Atoi(id)
	return userInfo, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:34561)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var uu model.User
	db.Debug().Table("user").Find(&uu, "id=?", id)
	return &uu, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:34561)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var users []*model.User
	db.Debug().Table("user").Find(&users)
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
