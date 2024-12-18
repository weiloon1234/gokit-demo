// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gokit-demo/ent/bank"
	"gokit-demo/ent/country"
	"gokit-demo/ent/countrylocation"
	"gokit-demo/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CountryCreate is the builder for creating a Country entity.
type CountryCreate struct {
	config
	mutation *CountryMutation
	hooks    []Hook
}

// SetIso2 sets the "iso2" field.
func (cc *CountryCreate) SetIso2(s string) *CountryCreate {
	cc.mutation.SetIso2(s)
	return cc
}

// SetIso3 sets the "iso3" field.
func (cc *CountryCreate) SetIso3(s string) *CountryCreate {
	cc.mutation.SetIso3(s)
	return cc
}

// SetName sets the "name" field.
func (cc *CountryCreate) SetName(s string) *CountryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetOfficialName sets the "official_name" field.
func (cc *CountryCreate) SetOfficialName(s string) *CountryCreate {
	cc.mutation.SetOfficialName(s)
	return cc
}

// SetNillableOfficialName sets the "official_name" field if the given value is not nil.
func (cc *CountryCreate) SetNillableOfficialName(s *string) *CountryCreate {
	if s != nil {
		cc.SetOfficialName(*s)
	}
	return cc
}

// SetNumericCode sets the "numeric_code" field.
func (cc *CountryCreate) SetNumericCode(s string) *CountryCreate {
	cc.mutation.SetNumericCode(s)
	return cc
}

// SetNillableNumericCode sets the "numeric_code" field if the given value is not nil.
func (cc *CountryCreate) SetNillableNumericCode(s *string) *CountryCreate {
	if s != nil {
		cc.SetNumericCode(*s)
	}
	return cc
}

// SetPhoneCode sets the "phone_code" field.
func (cc *CountryCreate) SetPhoneCode(s string) *CountryCreate {
	cc.mutation.SetPhoneCode(s)
	return cc
}

// SetNillablePhoneCode sets the "phone_code" field if the given value is not nil.
func (cc *CountryCreate) SetNillablePhoneCode(s *string) *CountryCreate {
	if s != nil {
		cc.SetPhoneCode(*s)
	}
	return cc
}

// SetCapital sets the "capital" field.
func (cc *CountryCreate) SetCapital(s string) *CountryCreate {
	cc.mutation.SetCapital(s)
	return cc
}

// SetNillableCapital sets the "capital" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCapital(s *string) *CountryCreate {
	if s != nil {
		cc.SetCapital(*s)
	}
	return cc
}

// SetCurrencyName sets the "currency_name" field.
func (cc *CountryCreate) SetCurrencyName(s string) *CountryCreate {
	cc.mutation.SetCurrencyName(s)
	return cc
}

// SetNillableCurrencyName sets the "currency_name" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCurrencyName(s *string) *CountryCreate {
	if s != nil {
		cc.SetCurrencyName(*s)
	}
	return cc
}

// SetCurrencyCode sets the "currency_code" field.
func (cc *CountryCreate) SetCurrencyCode(s string) *CountryCreate {
	cc.mutation.SetCurrencyCode(s)
	return cc
}

// SetNillableCurrencyCode sets the "currency_code" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCurrencyCode(s *string) *CountryCreate {
	if s != nil {
		cc.SetCurrencyCode(*s)
	}
	return cc
}

// SetCurrencySymbol sets the "currency_symbol" field.
func (cc *CountryCreate) SetCurrencySymbol(s string) *CountryCreate {
	cc.mutation.SetCurrencySymbol(s)
	return cc
}

// SetNillableCurrencySymbol sets the "currency_symbol" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCurrencySymbol(s *string) *CountryCreate {
	if s != nil {
		cc.SetCurrencySymbol(*s)
	}
	return cc
}

// SetConversionRate sets the "conversion_rate" field.
func (cc *CountryCreate) SetConversionRate(f float64) *CountryCreate {
	cc.mutation.SetConversionRate(f)
	return cc
}

// SetNillableConversionRate sets the "conversion_rate" field if the given value is not nil.
func (cc *CountryCreate) SetNillableConversionRate(f *float64) *CountryCreate {
	if f != nil {
		cc.SetConversionRate(*f)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CountryCreate) SetStatus(u uint8) *CountryCreate {
	cc.mutation.SetStatus(u)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CountryCreate) SetNillableStatus(u *uint8) *CountryCreate {
	if u != nil {
		cc.SetStatus(*u)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CountryCreate) SetCreatedAt(t time.Time) *CountryCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCreatedAt(t *time.Time) *CountryCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CountryCreate) SetUpdatedAt(t time.Time) *CountryCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableUpdatedAt(t *time.Time) *CountryCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CountryCreate) SetID(u uint64) *CountryCreate {
	cc.mutation.SetID(u)
	return cc
}

// AddLocationIDs adds the "locations" edge to the CountryLocation entity by IDs.
func (cc *CountryCreate) AddLocationIDs(ids ...uint64) *CountryCreate {
	cc.mutation.AddLocationIDs(ids...)
	return cc
}

// AddLocations adds the "locations" edges to the CountryLocation entity.
func (cc *CountryCreate) AddLocations(c ...*CountryLocation) *CountryCreate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddLocationIDs(ids...)
}

// AddBankIDs adds the "banks" edge to the Bank entity by IDs.
func (cc *CountryCreate) AddBankIDs(ids ...uint64) *CountryCreate {
	cc.mutation.AddBankIDs(ids...)
	return cc
}

// AddBanks adds the "banks" edges to the Bank entity.
func (cc *CountryCreate) AddBanks(b ...*Bank) *CountryCreate {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cc.AddBankIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cc *CountryCreate) AddUserIDs(ids ...uint64) *CountryCreate {
	cc.mutation.AddUserIDs(ids...)
	return cc
}

// AddUsers adds the "users" edges to the User entity.
func (cc *CountryCreate) AddUsers(u ...*User) *CountryCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddUserIDs(ids...)
}

// AddContactUserIDs adds the "contact_users" edge to the User entity by IDs.
func (cc *CountryCreate) AddContactUserIDs(ids ...uint64) *CountryCreate {
	cc.mutation.AddContactUserIDs(ids...)
	return cc
}

// AddContactUsers adds the "contact_users" edges to the User entity.
func (cc *CountryCreate) AddContactUsers(u ...*User) *CountryCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddContactUserIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cc *CountryCreate) Mutation() *CountryMutation {
	return cc.mutation
}

// Save creates the Country in the database.
func (cc *CountryCreate) Save(ctx context.Context) (*Country, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CountryCreate) SaveX(ctx context.Context) *Country {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CountryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CountryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CountryCreate) defaults() {
	if _, ok := cc.mutation.ConversionRate(); !ok {
		v := country.DefaultConversionRate
		cc.mutation.SetConversionRate(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := country.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := country.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := country.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CountryCreate) check() error {
	if _, ok := cc.mutation.Iso2(); !ok {
		return &ValidationError{Name: "iso2", err: errors.New(`ent: missing required field "Country.iso2"`)}
	}
	if v, ok := cc.mutation.Iso2(); ok {
		if err := country.Iso2Validator(v); err != nil {
			return &ValidationError{Name: "iso2", err: fmt.Errorf(`ent: validator failed for field "Country.iso2": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Iso3(); !ok {
		return &ValidationError{Name: "iso3", err: errors.New(`ent: missing required field "Country.iso3"`)}
	}
	if v, ok := cc.mutation.Iso3(); ok {
		if err := country.Iso3Validator(v); err != nil {
			return &ValidationError{Name: "iso3", err: fmt.Errorf(`ent: validator failed for field "Country.iso3": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Country.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := country.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Country.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.ConversionRate(); !ok {
		return &ValidationError{Name: "conversion_rate", err: errors.New(`ent: missing required field "Country.conversion_rate"`)}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Country.status"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Country.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Country.updated_at"`)}
	}
	return nil
}

func (cc *CountryCreate) sqlSave(ctx context.Context) (*Country, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CountryCreate) createSpec() (*Country, *sqlgraph.CreateSpec) {
	var (
		_node = &Country{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(country.Table, sqlgraph.NewFieldSpec(country.FieldID, field.TypeUint64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Iso2(); ok {
		_spec.SetField(country.FieldIso2, field.TypeString, value)
		_node.Iso2 = value
	}
	if value, ok := cc.mutation.Iso3(); ok {
		_spec.SetField(country.FieldIso3, field.TypeString, value)
		_node.Iso3 = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.OfficialName(); ok {
		_spec.SetField(country.FieldOfficialName, field.TypeString, value)
		_node.OfficialName = &value
	}
	if value, ok := cc.mutation.NumericCode(); ok {
		_spec.SetField(country.FieldNumericCode, field.TypeString, value)
		_node.NumericCode = &value
	}
	if value, ok := cc.mutation.PhoneCode(); ok {
		_spec.SetField(country.FieldPhoneCode, field.TypeString, value)
		_node.PhoneCode = &value
	}
	if value, ok := cc.mutation.Capital(); ok {
		_spec.SetField(country.FieldCapital, field.TypeString, value)
		_node.Capital = &value
	}
	if value, ok := cc.mutation.CurrencyName(); ok {
		_spec.SetField(country.FieldCurrencyName, field.TypeString, value)
		_node.CurrencyName = &value
	}
	if value, ok := cc.mutation.CurrencyCode(); ok {
		_spec.SetField(country.FieldCurrencyCode, field.TypeString, value)
		_node.CurrencyCode = &value
	}
	if value, ok := cc.mutation.CurrencySymbol(); ok {
		_spec.SetField(country.FieldCurrencySymbol, field.TypeString, value)
		_node.CurrencySymbol = &value
	}
	if value, ok := cc.mutation.ConversionRate(); ok {
		_spec.SetField(country.FieldConversionRate, field.TypeFloat64, value)
		_node.ConversionRate = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(country.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(country.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(country.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.LocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LocationsTable,
			Columns: []string{country.LocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.BanksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.BanksTable,
			Columns: []string{country.BanksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bank.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.UsersTable,
			Columns: []string{country.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ContactUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.ContactUsersTable,
			Columns: []string{country.ContactUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CountryCreateBulk is the builder for creating many Country entities in bulk.
type CountryCreateBulk struct {
	config
	err      error
	builders []*CountryCreate
}

// Save creates the Country entities in the database.
func (ccb *CountryCreateBulk) Save(ctx context.Context) ([]*Country, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Country, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CountryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CountryCreateBulk) SaveX(ctx context.Context) []*Country {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CountryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CountryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
