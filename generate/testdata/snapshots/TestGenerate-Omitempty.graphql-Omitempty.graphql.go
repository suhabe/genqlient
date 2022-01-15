// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/suhabe/genqlient/graphql"
	"github.com/suhabe/genqlient/internal/testutil"
)

// OmitEmptyQueryResponse is returned by OmitEmptyQuery on success.
type OmitEmptyQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User         OmitEmptyQueryUser        `json:"user"`
	Users        []OmitEmptyQueryUsersUser `json:"users"`
	MaybeConvert time.Time                 `json:"maybeConvert"`
	Convert2     time.Time                 `json:"convert2"`
}

// GetUser returns OmitEmptyQueryResponse.User, and is useful for accessing the field via an interface.
func (v *OmitEmptyQueryResponse) GetUser() OmitEmptyQueryUser { return v.User }

// GetUsers returns OmitEmptyQueryResponse.Users, and is useful for accessing the field via an interface.
func (v *OmitEmptyQueryResponse) GetUsers() []OmitEmptyQueryUsersUser { return v.Users }

// GetMaybeConvert returns OmitEmptyQueryResponse.MaybeConvert, and is useful for accessing the field via an interface.
func (v *OmitEmptyQueryResponse) GetMaybeConvert() time.Time { return v.MaybeConvert }

// GetConvert2 returns OmitEmptyQueryResponse.Convert2, and is useful for accessing the field via an interface.
func (v *OmitEmptyQueryResponse) GetConvert2() time.Time { return v.Convert2 }

// OmitEmptyQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type OmitEmptyQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns OmitEmptyQueryUser.Id, and is useful for accessing the field via an interface.
func (v *OmitEmptyQueryUser) GetId() testutil.ID { return v.Id }

// OmitEmptyQueryUsersUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type OmitEmptyQueryUsersUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns OmitEmptyQueryUsersUser.Id, and is useful for accessing the field via an interface.
func (v *OmitEmptyQueryUsersUser) GetId() testutil.ID { return v.Id }

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type UserQueryInput struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id         testutil.ID      `json:"id"`
	Role       Role             `json:"role,omitempty"`
	Names      []string         `json:"names,omitempty"`
	HasPokemon testutil.Pokemon `json:"hasPokemon,omitempty"`
	Birthdate  time.Time        `json:"-"`
}

// GetEmail returns UserQueryInput.Email, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetEmail() string { return v.Email }

// GetName returns UserQueryInput.Name, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetName() string { return v.Name }

// GetId returns UserQueryInput.Id, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetId() testutil.ID { return v.Id }

// GetRole returns UserQueryInput.Role, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetRole() Role { return v.Role }

// GetNames returns UserQueryInput.Names, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetNames() []string { return v.Names }

// GetHasPokemon returns UserQueryInput.HasPokemon, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetHasPokemon() testutil.Pokemon { return v.HasPokemon }

// GetBirthdate returns UserQueryInput.Birthdate, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetBirthdate() time.Time { return v.Birthdate }

func (v *UserQueryInput) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*UserQueryInput
		Birthdate json.RawMessage `json:"birthdate"`
		graphql.NoUnmarshalJSON
	}
	firstPass.UserQueryInput = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Birthdate
		src := firstPass.Birthdate
		if len(src) != 0 && string(src) != "null" {
			err = testutil.UnmarshalDate(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal UserQueryInput.Birthdate: %w", err)
			}
		}
	}
	return nil
}

type __premarshalUserQueryInput struct {
	Email string `json:"email"`

	Name string `json:"name"`

	Id testutil.ID `json:"id"`

	Role Role `json:"role"`

	Names []string `json:"names"`

	HasPokemon testutil.Pokemon `json:"hasPokemon"`

	Birthdate json.RawMessage `json:"birthdate"`
}

func (v *UserQueryInput) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *UserQueryInput) __premarshalJSON() (*__premarshalUserQueryInput, error) {
	var retval __premarshalUserQueryInput

	retval.Email = v.Email
	retval.Name = v.Name
	retval.Id = v.Id
	retval.Role = v.Role
	retval.Names = v.Names
	retval.HasPokemon = v.HasPokemon
	{

		dst := &retval.Birthdate
		src := v.Birthdate
		var err error
		*dst, err = testutil.MarshalDate(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"Unable to marshal UserQueryInput.Birthdate: %w", err)
		}
	}
	return &retval, nil
}

// __OmitEmptyQueryInput is used internally by genqlient
type __OmitEmptyQueryInput struct {
	Query         UserQueryInput   `json:"query,omitempty"`
	Queries       []UserQueryInput `json:"queries,omitempty"`
	Dt            time.Time        `json:"dt,omitempty"`
	Tz            string           `json:"tz,omitempty"`
	TzNoOmitEmpty string           `json:"tzNoOmitEmpty"`
}

// GetQuery returns __OmitEmptyQueryInput.Query, and is useful for accessing the field via an interface.
func (v *__OmitEmptyQueryInput) GetQuery() UserQueryInput { return v.Query }

// GetQueries returns __OmitEmptyQueryInput.Queries, and is useful for accessing the field via an interface.
func (v *__OmitEmptyQueryInput) GetQueries() []UserQueryInput { return v.Queries }

// GetDt returns __OmitEmptyQueryInput.Dt, and is useful for accessing the field via an interface.
func (v *__OmitEmptyQueryInput) GetDt() time.Time { return v.Dt }

// GetTz returns __OmitEmptyQueryInput.Tz, and is useful for accessing the field via an interface.
func (v *__OmitEmptyQueryInput) GetTz() string { return v.Tz }

// GetTzNoOmitEmpty returns __OmitEmptyQueryInput.TzNoOmitEmpty, and is useful for accessing the field via an interface.
func (v *__OmitEmptyQueryInput) GetTzNoOmitEmpty() string { return v.TzNoOmitEmpty }

func OmitEmptyQuery(
	client graphql.Client,
	query UserQueryInput,
	queries []UserQueryInput,
	dt time.Time,
	tz string,
	tzNoOmitEmpty string,
) (*OmitEmptyQueryResponse, error) {
	__input := __OmitEmptyQueryInput{
		Query:         query,
		Queries:       queries,
		Dt:            dt,
		Tz:            tz,
		TzNoOmitEmpty: tzNoOmitEmpty,
	}
	var err error

	var retval OmitEmptyQueryResponse
	err = client.MakeRequest(
		nil,
		"OmitEmptyQuery",
		`
query OmitEmptyQuery ($query: UserQueryInput, $queries: [UserQueryInput], $dt: DateTime, $tz: String, $tzNoOmitEmpty: String) {
	user(query: $query) {
		id
	}
	users(query: $queries) {
		id
	}
	maybeConvert(dt: $dt, tz: $tz)
	convert2: maybeConvert(dt: $dt, tz: $tzNoOmitEmpty)
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

