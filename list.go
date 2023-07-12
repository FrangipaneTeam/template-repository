package supertypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ListTypable = ListType{}

type ListType struct {
	basetypes.ListType
	// ... potentially other fields ...
}

func (t ListType) Equal(o attr.Type) bool {
	other, ok := o.(basetypes.ListType)

	if !ok {
		return false
	}

	return t.ListType.Equal(other)
}

func (t ListType) String() string {
	return "types.ListType[" + t.ElementType().String() + "]"
}

func (t ListType) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	// CustomListValue defined in the value type section
	value := ListValue{
		ListValue: in,
	}

	return value, nil
}

func (t ListType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	ListValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromList(ctx, ListValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return stringValuable, nil
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the AttributeType to something Terraform
// can understand.
func (t ListType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.List{
		ElementType: t.ElementType().TerraformType(ctx),
	}
}

func (t ListType) ElementType() attr.Type {
	if t.ListType.ElemType == nil {
		return missingType{}
	}

	return t.ListType.ElemType
}

func (t ListType) ValueType(ctx context.Context) attr.Value {
	// CustomListValue defined in the value type section
	return ListValue{
		ListValue: t.ListType.ValueType(ctx).(basetypes.ListValue),
	}
}

// * ---------
// * VALUE

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ListValuable = ListValue{}

type ListValue struct {
	basetypes.ListValue
}

func (v ListValue) Equal(o attr.Value) bool {
	other, ok := o.(ListValue)

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v ListValue) Type(ctx context.Context) attr.Type {
	// CustomListType defined in the schema type section
	return v.ListValue.Type(ctx)
}

func (v ListValue) ToListValue(ctx context.Context) (basetypes.ListValue, diag.Diagnostics) {
	return v.ListValue, nil
}

func ListNull(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListNull(elementType),
	}
}

func ListUnknown(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListUnknown(elementType),
	}
}

// * CustomFunc

func (v *ListValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.ListValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *ListValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ListValue, d = types.ListValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *ListValue) SetNull(ctx context.Context) {
	v.ListValue = basetypes.NewListNull(v.ElementType(ctx))
}

func (v *ListValue) SetUnknown(ctx context.Context) {
	v.ListValue = basetypes.NewListUnknown(v.ElementType(ctx))
}

func (v ListValue) IsKnown() bool {
	return !v.ListValue.IsNull() && !v.ListValue.IsUnknown()
}
