package schema

import (
	"entgo.io/ent"
	"github.com/weiloon1234/gokit/ent/schema"
)

// User holds the schema definition for the User entity.
type Country struct {
	schema.Country
}

// Fields of the User.
func (Country) Fields() []ent.Field {
	return schema.Country{}.Fields()
}
