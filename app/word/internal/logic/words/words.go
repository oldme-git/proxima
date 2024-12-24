// Package words implements the business logic for word management
package words

import (
    "context"

    "proxima/app/word/internal/dao"
    "proxima/app/word/internal/model/do"
    "proxima/app/word/internal/model/entity"
)

// Words implements the word management logic layer
type Words struct{}

// New creates and returns a new Words instance
func New() *Words {
    return &Words{}
}

// CreateInput defines the required fields for creating a new word
type CreateInput struct {
    Uid        uint32 // User ID who created the word
    Word       string // The word to be stored
    Definition string // Definition or meaning of the word
}

// Create stores a new word with its associated information
// Returns the new word's ID and any error encountered during creation
func (*Words) Create(ctx context.Context, in CreateInput) (id int64, err error) {
    return dao.Words.Ctx(ctx).Data(do.Words{
        Uid:        in.Uid,
        Word:       in.Word,
        Definition: in.Definition,
    }).InsertAndGetId()
}

// Get retrieves a word's complete information by its ID
// Returns the word entity if found, or error if not found or on database errors
func (*Words) Get(ctx context.Context, id uint32) (word *entity.Words, err error) {
    err = dao.Words.Ctx(ctx).WherePri(id).Scan(&word)
    return
}
