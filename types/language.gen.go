//  Source: lemmy/crates/db_schema/src/source/language.rs
// Code generated by go.elara.ws/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type Language struct {
	ID   int    `json:"id" url:"id,omitempty"`
	Code string `json:"code" url:"code,omitempty"`
	Name string `json:"name" url:"name,omitempty"`
}
