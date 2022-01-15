// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"time"

	"github.com/suhabe/genqlient/graphql"
)

// __convertTimezoneInput is used internally by genqlient
type __convertTimezoneInput struct {
	Dt time.Time `json:"dt"`
	Tz string    `json:"tz"`
}

// GetDt returns __convertTimezoneInput.Dt, and is useful for accessing the field via an interface.
func (v *__convertTimezoneInput) GetDt() time.Time { return v.Dt }

// GetTz returns __convertTimezoneInput.Tz, and is useful for accessing the field via an interface.
func (v *__convertTimezoneInput) GetTz() string { return v.Tz }

// convertTimezoneResponse is returned by convertTimezone on success.
type convertTimezoneResponse struct {
	Convert time.Time `json:"convert"`
}

// GetConvert returns convertTimezoneResponse.Convert, and is useful for accessing the field via an interface.
func (v *convertTimezoneResponse) GetConvert() time.Time { return v.Convert }

func convertTimezone(
	client graphql.Client,
	dt time.Time,
	tz string,
) (*convertTimezoneResponse, error) {
	__input := __convertTimezoneInput{
		Dt: dt,
		Tz: tz,
	}
	var err error

	var retval convertTimezoneResponse
	err = client.MakeRequest(
		nil,
		"convertTimezone",
		`
query convertTimezone ($dt: DateTime!, $tz: String) {
	convert(dt: $dt, tz: $tz)
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

