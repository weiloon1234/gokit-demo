package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/weiloon1234/gokit-base-entity/ent/schema"

	"github.com/weiloon1234/gokit/utils"
)

// User holds the schema definition for the User entity.
type User struct {
	schema.User
}

// Fields of the User.
func (User) Fields() []ent.Field {
	baseFields := schema.User{}.Fields()

	// Insert password3 after password2
	modifiedFields := utils.InsertFieldsAfter(baseFields, "password2",
		field.String("password3").Optional().Comment("Third password for the user"),
	)

	return modifiedFields
}
