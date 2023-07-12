package supertypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var TypeToto = StringType{}

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.StringTypable = StringType{}

type StringType struct {
	basetypes.StringType
	// ... potentially other fields ...
}

func (t StringType) Equal(o attr.Type) bool {
	other, ok := o.(StringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t StringType) String() string {
	return "SuperTypesStringType"
}

func (t StringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// CustomStringValue defined in the value type section
	value := StringValue{
		StringValue: in,
	}

	return value, nil
}

func (t StringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t StringType) ValueType(ctx context.Context) attr.Value {
	// CustomStringValue defined in the value type section
	return StringValue{}
}

// * ---------
// * VALUE

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.StringValuable = StringValue{}

type StringValue struct {
	basetypes.StringValue
	// ... potentially other fields ...
}

func (v StringValue) Equal(o attr.Value) bool {
	other, ok := o.(StringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v StringValue) Type(ctx context.Context) attr.Type {
	// CustomStringType defined in the schema type section
	return StringType{}
}

func StringNull() StringValue {
	return StringValue{
		StringValue: basetypes.NewStringNull(),
	}
}

func StringUnknown() StringValue {
	return StringValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// * CustomFunc

func (v *StringValue) Get() string {
	return v.StringValue.ValueString()
}

func (v *StringValue) Set(s string) {
	v.StringValue = basetypes.NewStringValue(s)
}

func (v *StringValue) SetNull() {
	v.StringValue = basetypes.NewStringNull()
}

func (v *StringValue) SetUnknown() {
	v.StringValue = basetypes.NewStringUnknown()
}

func (v StringValue) IsKnown() bool {
	return !v.StringValue.IsNull() && !v.StringValue.IsUnknown()
}
