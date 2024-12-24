// Package words implements the business logic for word management
package words

import (
	"context"

	"proxima/app/word/internal/dao"
	"proxima/app/word/internal/model/do"
	"proxima/app/word/internal/model/entity"
)

type Words struct{}

func New() *Words {
	return &Words{}
}

type CreateInput struct {
	Uid        uint32
	Word       string
	Definition string
}

func (*Words) Create(ctx context.Context, in CreateInput) (id int64, err error) {
	return dao.Words.Ctx(ctx).Data(do.Words{
		Uid:        in.Uid,
		Word:       in.Word,
		Definition: in.Definition,
	}).InsertAndGetId()
}

func (*Words) Get(ctx context.Context, id uint32) (word *entity.Words, err error) {
	err = dao.Words.Ctx(ctx).WherePri(id).Scan(&word)
	return
}
