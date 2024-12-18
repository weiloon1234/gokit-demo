package schema

import (
	"entgo.io/ent"
	"github.com/weiloon1234/gokit/ent/schema"
)

// Bank holds the schema definition for the Bank entity.
type Bank struct {
	schema.Bank
}

// Fields of the Bank.
func (Bank) Fields() []ent.Field {
	return schema.Bank{}.Fields()
}
