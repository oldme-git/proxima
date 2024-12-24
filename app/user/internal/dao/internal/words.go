// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WordsDao is the data access object for the table words.
type WordsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns WordsColumns // columns contains all the column names of Table for convenient usage.
}

// WordsColumns defines and stores column names for the table words.
type WordsColumns struct {
	Id                 string //
	Uid                string //
	Word               string //
	Definition         string //
	ExampleSentence    string //
	ChineseTranslation string //
	Pronunciation      string //
	CreatedAt          string //
	UpdatedAt          string //
}

// wordsColumns holds the columns for the table words.
var wordsColumns = WordsColumns{
	Id:                 "id",
	Uid:                "uid",
	Word:               "word",
	Definition:         "definition",
	ExampleSentence:    "example_sentence",
	ChineseTranslation: "chinese_translation",
	Pronunciation:      "pronunciation",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewWordsDao creates and returns a new DAO object for table data access.
func NewWordsDao() *WordsDao {
	return &WordsDao{
		group:   "default",
		table:   "words",
		columns: wordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WordsDao) Columns() WordsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
