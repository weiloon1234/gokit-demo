package schema

import (
	"entgo.io/ent"
	"github.com/weiloon1234/gokit/ent/schema"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	schema.Admin
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return schema.Admin{}.Fields()
}
