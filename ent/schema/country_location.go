package schema

import (
	"entgo.io/ent"
	"github.com/weiloon1234/gokit/ent/schema"
)

// CountryLocation holds the schema definition for the CountryLocation entity.
type CountryLocation struct {
	schema.CountryLocation
}

// Fields of the CountryLocation.
func (CountryLocation) Fields() []ent.Field {
	return schema.CountryLocation{}.Fields()
}
