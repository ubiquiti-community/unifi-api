// Package fields holds the field-type tokens shared by the generator. These
// mirror go-unifi's internal/fields constants so the ported FieldProcessor
// logic stays identical; only String/Int/Bool/Number affect OpenAPI output
// (the "types.*" tokens are inert metadata for this repo).
package fields

const (
	Int    = "int64"
	Bool   = "bool"
	String = "string"
	Number = "types.Number"
)
