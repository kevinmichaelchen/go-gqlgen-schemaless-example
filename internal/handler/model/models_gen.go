// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateFooInput struct {
	// A random grab-bag of attributes.
	// These aren't explicitly modeled for a variety of reasons:
	//   - they change often; they're ephemeral; they'll be deprecated soon
	//   - they're too bespoke and not really the concern of this API
	Attributes map[string]interface{} `json:"attributes"`
}

type Foo struct {
	Name string `json:"name"`
}
