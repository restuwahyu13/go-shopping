package repo

import (
	rinf "restuwahyu13/shopping-cart/internal/domain/interface/repository"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type userRepository struct {
	db    *bun.DB
	model *model.UsersModel
}

func NewUsersRepository(db *bun.DB) rinf.IUserRepository {
	return userRepository{db: db, model: new(model.UsersModel)}
}

func (r userRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r userRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r userRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r userRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}
