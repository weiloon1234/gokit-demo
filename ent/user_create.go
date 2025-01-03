// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gokit-demo/ent/bank"
	"gokit-demo/ent/country"
	"gokit-demo/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uc *UserCreate) SetNillableUsername(s *string) *UserCreate {
	if s != nil {
		uc.SetUsername(*s)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetEmailVerifiedAt sets the "email_verified_at" field.
func (uc *UserCreate) SetEmailVerifiedAt(t time.Time) *UserCreate {
	uc.mutation.SetEmailVerifiedAt(t)
	return uc
}

// SetNillableEmailVerifiedAt sets the "email_verified_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmailVerifiedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetEmailVerifiedAt(*t)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetPassword2 sets the "password2" field.
func (uc *UserCreate) SetPassword2(s string) *UserCreate {
	uc.mutation.SetPassword2(s)
	return uc
}

// SetCountryID sets the "country_id" field.
func (uc *UserCreate) SetCountryID(u uint64) *UserCreate {
	uc.mutation.SetCountryID(u)
	return uc
}

// SetNillableCountryID sets the "country_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableCountryID(u *uint64) *UserCreate {
	if u != nil {
		uc.SetCountryID(*u)
	}
	return uc
}

// SetContactCountryID sets the "contact_country_id" field.
func (uc *UserCreate) SetContactCountryID(u uint64) *UserCreate {
	uc.mutation.SetContactCountryID(u)
	return uc
}

// SetNillableContactCountryID sets the "contact_country_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableContactCountryID(u *uint64) *UserCreate {
	if u != nil {
		uc.SetContactCountryID(*u)
	}
	return uc
}

// SetContactNumber sets the "contact_number" field.
func (uc *UserCreate) SetContactNumber(s string) *UserCreate {
	uc.mutation.SetContactNumber(s)
	return uc
}

// SetNillableContactNumber sets the "contact_number" field if the given value is not nil.
func (uc *UserCreate) SetNillableContactNumber(s *string) *UserCreate {
	if s != nil {
		uc.SetContactNumber(*s)
	}
	return uc
}

// SetFullContactNumber sets the "full_contact_number" field.
func (uc *UserCreate) SetFullContactNumber(s string) *UserCreate {
	uc.mutation.SetFullContactNumber(s)
	return uc
}

// SetNillableFullContactNumber sets the "full_contact_number" field if the given value is not nil.
func (uc *UserCreate) SetNillableFullContactNumber(s *string) *UserCreate {
	if s != nil {
		uc.SetFullContactNumber(*s)
	}
	return uc
}

// SetIntroducerUserID sets the "introducer_user_id" field.
func (uc *UserCreate) SetIntroducerUserID(u uint64) *UserCreate {
	uc.mutation.SetIntroducerUserID(u)
	return uc
}

// SetNillableIntroducerUserID sets the "introducer_user_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableIntroducerUserID(u *uint64) *UserCreate {
	if u != nil {
		uc.SetIntroducerUserID(*u)
	}
	return uc
}

// SetLang sets the "lang" field.
func (uc *UserCreate) SetLang(s string) *UserCreate {
	uc.mutation.SetLang(s)
	return uc
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (uc *UserCreate) SetNillableLang(s *string) *UserCreate {
	if s != nil {
		uc.SetLang(*s)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetCredit1 sets the "credit_1" field.
func (uc *UserCreate) SetCredit1(f float64) *UserCreate {
	uc.mutation.SetCredit1(f)
	return uc
}

// SetNillableCredit1 sets the "credit_1" field if the given value is not nil.
func (uc *UserCreate) SetNillableCredit1(f *float64) *UserCreate {
	if f != nil {
		uc.SetCredit1(*f)
	}
	return uc
}

// SetCredit2 sets the "credit_2" field.
func (uc *UserCreate) SetCredit2(f float64) *UserCreate {
	uc.mutation.SetCredit2(f)
	return uc
}

// SetNillableCredit2 sets the "credit_2" field if the given value is not nil.
func (uc *UserCreate) SetNillableCredit2(f *float64) *UserCreate {
	if f != nil {
		uc.SetCredit2(*f)
	}
	return uc
}

// SetCredit3 sets the "credit_3" field.
func (uc *UserCreate) SetCredit3(f float64) *UserCreate {
	uc.mutation.SetCredit3(f)
	return uc
}

// SetNillableCredit3 sets the "credit_3" field if the given value is not nil.
func (uc *UserCreate) SetNillableCredit3(f *float64) *UserCreate {
	if f != nil {
		uc.SetCredit3(*f)
	}
	return uc
}

// SetCredit4 sets the "credit_4" field.
func (uc *UserCreate) SetCredit4(f float64) *UserCreate {
	uc.mutation.SetCredit4(f)
	return uc
}

// SetNillableCredit4 sets the "credit_4" field if the given value is not nil.
func (uc *UserCreate) SetNillableCredit4(f *float64) *UserCreate {
	if f != nil {
		uc.SetCredit4(*f)
	}
	return uc
}

// SetCredit5 sets the "credit_5" field.
func (uc *UserCreate) SetCredit5(f float64) *UserCreate {
	uc.mutation.SetCredit5(f)
	return uc
}

// SetNillableCredit5 sets the "credit_5" field if the given value is not nil.
func (uc *UserCreate) SetNillableCredit5(f *float64) *UserCreate {
	if f != nil {
		uc.SetCredit5(*f)
	}
	return uc
}

// SetBankID sets the "bank_id" field.
func (uc *UserCreate) SetBankID(u uint64) *UserCreate {
	uc.mutation.SetBankID(u)
	return uc
}

// SetNillableBankID sets the "bank_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableBankID(u *uint64) *UserCreate {
	if u != nil {
		uc.SetBankID(*u)
	}
	return uc
}

// SetBankAccountName sets the "bank_account_name" field.
func (uc *UserCreate) SetBankAccountName(s string) *UserCreate {
	uc.mutation.SetBankAccountName(s)
	return uc
}

// SetNillableBankAccountName sets the "bank_account_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableBankAccountName(s *string) *UserCreate {
	if s != nil {
		uc.SetBankAccountName(*s)
	}
	return uc
}

// SetBankAccountNumber sets the "bank_account_number" field.
func (uc *UserCreate) SetBankAccountNumber(s string) *UserCreate {
	uc.mutation.SetBankAccountNumber(s)
	return uc
}

// SetNillableBankAccountNumber sets the "bank_account_number" field if the given value is not nil.
func (uc *UserCreate) SetNillableBankAccountNumber(s *string) *UserCreate {
	if s != nil {
		uc.SetBankAccountNumber(*s)
	}
	return uc
}

// SetNationalID sets the "national_id" field.
func (uc *UserCreate) SetNationalID(s string) *UserCreate {
	uc.mutation.SetNationalID(s)
	return uc
}

// SetNillableNationalID sets the "national_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableNationalID(s *string) *UserCreate {
	if s != nil {
		uc.SetNationalID(*s)
	}
	return uc
}

// SetFirstLogin sets the "first_login" field.
func (uc *UserCreate) SetFirstLogin(b bool) *UserCreate {
	uc.mutation.SetFirstLogin(b)
	return uc
}

// SetNillableFirstLogin sets the "first_login" field if the given value is not nil.
func (uc *UserCreate) SetNillableFirstLogin(b *bool) *UserCreate {
	if b != nil {
		uc.SetFirstLogin(*b)
	}
	return uc
}

// SetBanUntil sets the "ban_until" field.
func (uc *UserCreate) SetBanUntil(t time.Time) *UserCreate {
	uc.mutation.SetBanUntil(t)
	return uc
}

// SetNillableBanUntil sets the "ban_until" field if the given value is not nil.
func (uc *UserCreate) SetNillableBanUntil(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetBanUntil(*t)
	}
	return uc
}

// SetNewLoginAt sets the "new_login_at" field.
func (uc *UserCreate) SetNewLoginAt(t time.Time) *UserCreate {
	uc.mutation.SetNewLoginAt(t)
	return uc
}

// SetNillableNewLoginAt sets the "new_login_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableNewLoginAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetNewLoginAt(*t)
	}
	return uc
}

// SetLastLoginAt sets the "last_login_at" field.
func (uc *UserCreate) SetLastLoginAt(t time.Time) *UserCreate {
	uc.mutation.SetLastLoginAt(t)
	return uc
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastLoginAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastLoginAt(*t)
	}
	return uc
}

// SetUnilevel sets the "unilevel" field.
func (uc *UserCreate) SetUnilevel(u uint64) *UserCreate {
	uc.mutation.SetUnilevel(u)
	return uc
}

// SetNillableUnilevel sets the "unilevel" field if the given value is not nil.
func (uc *UserCreate) SetNillableUnilevel(u *uint64) *UserCreate {
	if u != nil {
		uc.SetUnilevel(*u)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetDeletedAt sets the "deleted_at" field.
func (uc *UserCreate) SetDeletedAt(t time.Time) *UserCreate {
	uc.mutation.SetDeletedAt(t)
	return uc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetDeletedAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uint64) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetCountry sets the "country" edge to the Country entity.
func (uc *UserCreate) SetCountry(c *Country) *UserCreate {
	return uc.SetCountryID(c.ID)
}

// SetContactCountry sets the "contact_country" edge to the Country entity.
func (uc *UserCreate) SetContactCountry(c *Country) *UserCreate {
	return uc.SetContactCountryID(c.ID)
}

// SetIntroducerID sets the "introducer" edge to the User entity by ID.
func (uc *UserCreate) SetIntroducerID(id uint64) *UserCreate {
	uc.mutation.SetIntroducerID(id)
	return uc
}

// SetNillableIntroducerID sets the "introducer" edge to the User entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableIntroducerID(id *uint64) *UserCreate {
	if id != nil {
		uc = uc.SetIntroducerID(*id)
	}
	return uc
}

// SetIntroducer sets the "introducer" edge to the User entity.
func (uc *UserCreate) SetIntroducer(u *User) *UserCreate {
	return uc.SetIntroducerID(u.ID)
}

// SetBank sets the "bank" edge to the Bank entity.
func (uc *UserCreate) SetBank(b *Bank) *UserCreate {
	return uc.SetBankID(b.ID)
}

// AddIntroducedUserIDs adds the "introduced_users" edge to the User entity by IDs.
func (uc *UserCreate) AddIntroducedUserIDs(ids ...uint64) *UserCreate {
	uc.mutation.AddIntroducedUserIDs(ids...)
	return uc
}

// AddIntroducedUsers adds the "introduced_users" edges to the User entity.
func (uc *UserCreate) AddIntroducedUsers(u ...*User) *UserCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddIntroducedUserIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Lang(); !ok {
		v := user.DefaultLang
		uc.mutation.SetLang(v)
	}
	if _, ok := uc.mutation.Credit1(); !ok {
		v := user.DefaultCredit1
		uc.mutation.SetCredit1(v)
	}
	if _, ok := uc.mutation.Credit2(); !ok {
		v := user.DefaultCredit2
		uc.mutation.SetCredit2(v)
	}
	if _, ok := uc.mutation.Credit3(); !ok {
		v := user.DefaultCredit3
		uc.mutation.SetCredit3(v)
	}
	if _, ok := uc.mutation.Credit4(); !ok {
		v := user.DefaultCredit4
		uc.mutation.SetCredit4(v)
	}
	if _, ok := uc.mutation.Credit5(); !ok {
		v := user.DefaultCredit5
		uc.mutation.SetCredit5(v)
	}
	if _, ok := uc.mutation.FirstLogin(); !ok {
		v := user.DefaultFirstLogin
		uc.mutation.SetFirstLogin(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Password2(); !ok {
		return &ValidationError{Name: "password2", err: errors.New(`ent: missing required field "User.password2"`)}
	}
	if v, ok := uc.mutation.Password2(); ok {
		if err := user.Password2Validator(v); err != nil {
			return &ValidationError{Name: "password2", err: fmt.Errorf(`ent: validator failed for field "User.password2": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Lang(); !ok {
		return &ValidationError{Name: "lang", err: errors.New(`ent: missing required field "User.lang"`)}
	}
	if _, ok := uc.mutation.Credit1(); !ok {
		return &ValidationError{Name: "credit_1", err: errors.New(`ent: missing required field "User.credit_1"`)}
	}
	if _, ok := uc.mutation.Credit2(); !ok {
		return &ValidationError{Name: "credit_2", err: errors.New(`ent: missing required field "User.credit_2"`)}
	}
	if _, ok := uc.mutation.Credit3(); !ok {
		return &ValidationError{Name: "credit_3", err: errors.New(`ent: missing required field "User.credit_3"`)}
	}
	if _, ok := uc.mutation.Credit4(); !ok {
		return &ValidationError{Name: "credit_4", err: errors.New(`ent: missing required field "User.credit_4"`)}
	}
	if _, ok := uc.mutation.Credit5(); !ok {
		return &ValidationError{Name: "credit_5", err: errors.New(`ent: missing required field "User.credit_5"`)}
	}
	if _, ok := uc.mutation.FirstLogin(); !ok {
		return &ValidationError{Name: "first_login", err: errors.New(`ent: missing required field "User.first_login"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.EmailVerifiedAt(); ok {
		_spec.SetField(user.FieldEmailVerifiedAt, field.TypeTime, value)
		_node.EmailVerifiedAt = &value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Password2(); ok {
		_spec.SetField(user.FieldPassword2, field.TypeString, value)
		_node.Password2 = value
	}
	if value, ok := uc.mutation.ContactNumber(); ok {
		_spec.SetField(user.FieldContactNumber, field.TypeString, value)
		_node.ContactNumber = &value
	}
	if value, ok := uc.mutation.FullContactNumber(); ok {
		_spec.SetField(user.FieldFullContactNumber, field.TypeString, value)
		_node.FullContactNumber = &value
	}
	if value, ok := uc.mutation.Lang(); ok {
		_spec.SetField(user.FieldLang, field.TypeString, value)
		_node.Lang = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = &value
	}
	if value, ok := uc.mutation.Credit1(); ok {
		_spec.SetField(user.FieldCredit1, field.TypeFloat64, value)
		_node.Credit1 = value
	}
	if value, ok := uc.mutation.Credit2(); ok {
		_spec.SetField(user.FieldCredit2, field.TypeFloat64, value)
		_node.Credit2 = value
	}
	if value, ok := uc.mutation.Credit3(); ok {
		_spec.SetField(user.FieldCredit3, field.TypeFloat64, value)
		_node.Credit3 = value
	}
	if value, ok := uc.mutation.Credit4(); ok {
		_spec.SetField(user.FieldCredit4, field.TypeFloat64, value)
		_node.Credit4 = value
	}
	if value, ok := uc.mutation.Credit5(); ok {
		_spec.SetField(user.FieldCredit5, field.TypeFloat64, value)
		_node.Credit5 = value
	}
	if value, ok := uc.mutation.BankAccountName(); ok {
		_spec.SetField(user.FieldBankAccountName, field.TypeString, value)
		_node.BankAccountName = &value
	}
	if value, ok := uc.mutation.BankAccountNumber(); ok {
		_spec.SetField(user.FieldBankAccountNumber, field.TypeString, value)
		_node.BankAccountNumber = &value
	}
	if value, ok := uc.mutation.NationalID(); ok {
		_spec.SetField(user.FieldNationalID, field.TypeString, value)
		_node.NationalID = &value
	}
	if value, ok := uc.mutation.FirstLogin(); ok {
		_spec.SetField(user.FieldFirstLogin, field.TypeBool, value)
		_node.FirstLogin = value
	}
	if value, ok := uc.mutation.BanUntil(); ok {
		_spec.SetField(user.FieldBanUntil, field.TypeTime, value)
		_node.BanUntil = &value
	}
	if value, ok := uc.mutation.NewLoginAt(); ok {
		_spec.SetField(user.FieldNewLoginAt, field.TypeTime, value)
		_node.NewLoginAt = &value
	}
	if value, ok := uc.mutation.LastLoginAt(); ok {
		_spec.SetField(user.FieldLastLoginAt, field.TypeTime, value)
		_node.LastLoginAt = &value
	}
	if value, ok := uc.mutation.Unilevel(); ok {
		_spec.SetField(user.FieldUnilevel, field.TypeUint64, value)
		_node.Unilevel = &value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := uc.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CountryTable,
			Columns: []string{user.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CountryID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ContactCountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ContactCountryTable,
			Columns: []string{user.ContactCountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ContactCountryID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.IntroducerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.IntroducerTable,
			Columns: []string{user.IntroducerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.IntroducerUserID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.BankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.BankTable,
			Columns: []string{user.BankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bank.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BankID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.IntroducedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.IntroducedUsersTable,
			Columns: []string{user.IntroducedUsersColumn},
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

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
