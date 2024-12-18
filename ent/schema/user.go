package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/weiloon1234/gokit/ent/schema"
)

// User holds the schema definition for the User entity.
type User struct {
	schema.User
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return append(schema.User{}.Fields(),
		field.String("password3").Optional(), // Add a new field
	)
}
