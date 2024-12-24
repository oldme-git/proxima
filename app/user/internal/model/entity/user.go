// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       uint        `json:"id"       orm:"id"        description:"User ID"`            // User ID
	Passport string      `json:"passport" orm:"passport"  description:"User Passport"`      // User Passport
	Password string      `json:"password" orm:"password"  description:"User Password"`      // User Password
	Email    string      `json:"email"    orm:"Email"     description:"User Email Address"` // User Email Address
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"Created Time"`       // Created Time
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"Updated Time"`       // Updated Time
}
