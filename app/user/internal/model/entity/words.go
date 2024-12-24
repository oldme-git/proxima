// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure for table words.
type Words struct {
	Id                 uint        `json:"id"                 orm:"id"                  description:""` //
	Uid                uint        `json:"uid"                orm:"uid"                 description:""` //
	Word               string      `json:"word"               orm:"word"                description:""` //
	Definition         string      `json:"definition"         orm:"definition"          description:""` //
	ExampleSentence    string      `json:"exampleSentence"    orm:"example_sentence"    description:""` //
	ChineseTranslation string      `json:"chineseTranslation" orm:"chinese_translation" description:""` //
	Pronunciation      string      `json:"pronunciation"      orm:"pronunciation"       description:""` //
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          description:""` //
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          description:""` //
}
