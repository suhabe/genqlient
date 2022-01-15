// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package queries

import (
	"context"

	"github.com/suhabe/genqlient/internal/testutil"
)

// SimpleQueryResponse is returned by SimpleQuery on success.
type SimpleQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleQueryUser `json:"user"`
}

// GetUser returns SimpleQueryResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryResponse) GetUser() SimpleQueryUser { return v.User }

// SimpleQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id string `json:"id"`
}

// GetId returns SimpleQueryUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryUser) GetId() string { return v.Id }

func SimpleQuery(
	ctx context.Context,
) (*SimpleQueryResponse, error) {
	var err error
	client, err := testutil.GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var retval SimpleQueryResponse
	err = client.MakeRequest(
		ctx,
		"SimpleQuery",
		`
query SimpleQuery {
	user {
		id
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

