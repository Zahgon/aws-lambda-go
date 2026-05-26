// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package events

// DynamoDBAttributeValue provides convenient access for a value stored in DynamoDB.
// For more information,  please see http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeValue.html
type DynamoDBAttributeValue struct {
	value    anyValue
	dataType DynamoDBDataType
}

// This struct represents DynamoDBAttributeValue which doesn't
// implement fmt.Stringer interface and safely `fmt.Sprintf`able
type dynamoDbAttributeValue DynamoDBAttributeValue //nolint: staticcheck

// Binary provides access to an attribute of type Binary.
// Method panics if the attribute is not of type Binary.
func (av DynamoDBAttributeValue) Binary() []byte { _ = "STUB: not implemented"; return nil }

// Boolean provides access to an attribute of type Boolean.
// Method panics if the attribute is not of type Boolean.
func (av DynamoDBAttributeValue) Boolean() bool { _ = "STUB: not implemented"; return false }

// BinarySet provides access to an attribute of type Binary Set.
// Method panics if the attribute is not of type BinarySet.
func (av DynamoDBAttributeValue) BinarySet() [][]byte { _ = "STUB: not implemented"; return nil }

// List provides access to an attribute of type List. Each element
// of the list is an DynamoDBAttributeValue itself.
// Method panics if the attribute is not of type List.
func (av DynamoDBAttributeValue) List() []DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return nil
}

// Map provides access to an attribute of type Map. They Keys are strings
// and the values are DynamoDBAttributeValue instances.
// Method panics if the attribute is not of type Map.
func (av DynamoDBAttributeValue) Map() map[string]DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return nil
}

// Number provides access to an attribute of type Number.
// DynamoDB sends the values as strings. For convenience please see also
// the methods Integer() and Float().
// Method panics if the attribute is not of type Number.
func (av DynamoDBAttributeValue) Number() string { _ = "STUB: not implemented"; return "" }

// Int64 provides access to an attribute of type Number.
// DynamoDB sends the values as strings. For convenience this method
// provides conversion to int.
// Method panics if the attribute is not of type Number.
func (av DynamoDBAttributeValue) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Integer provides access to an attribute of type Number.
// DynamoDB sends the values as strings. For convenience this method
// provides conversion to int. If the value cannot be represented by
// a signed integer, err.Err = ErrRange and the returned value is the maximum magnitude integer
// of an int64 of the appropriate sign.
// Method panics if the attribute is not of type Number.
func (av DynamoDBAttributeValue) Integer() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Float provides access to an attribute of type Number.
// DynamoDB sends the values as strings. For convenience this method
// provides conversion to float64.
// The returned value is the nearest floating point number rounded using IEEE754 unbiased rounding.
// If the number is more than 1/2 ULP away from the largest floating point number of the given size,
// the value returned is ±Inf, err.Err = ErrRange.
// Method panics if the attribute is not of type Number.
func (av DynamoDBAttributeValue) Float() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// NumberSet provides access to an attribute of type Number Set.
// DynamoDB sends the numbers as strings.
// Method panics if the attribute is not of type Number.
func (av DynamoDBAttributeValue) NumberSet() []string { _ = "STUB: not implemented"; return nil }

// String provides access to an attribute of type String.
// Method panics if the attribute is not of type String.
func (av DynamoDBAttributeValue) String() string { _ = "STUB: not implemented"; return "" }

// If dataType is not DataTypeString during fmt.Sprintf("%#v", ...)
// compiler confuses with fmt.Stringer interface and panics
// instead of printing the struct.

// StringSet provides access to an attribute of type String Set.
// Method panics if the attribute is not of type String Set.
func (av DynamoDBAttributeValue) StringSet() []string { _ = "STUB: not implemented"; return nil }

// IsNull returns true if the attribute is of type Null.
func (av DynamoDBAttributeValue) IsNull() bool { _ = "STUB: not implemented"; return false }

// DataType provides access to the DynamoDB type of the attribute
func (av DynamoDBAttributeValue) DataType() DynamoDBDataType {
	_ = "STUB: not implemented"
	return *

	// NewBinaryAttribute creates an DynamoDBAttributeValue containing a Binary
	new(DynamoDBDataType)
}

func NewBinaryAttribute(value []byte) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewBooleanAttribute creates an DynamoDBAttributeValue containing a Boolean
func NewBooleanAttribute(value bool) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewBinarySetAttribute creates an DynamoDBAttributeValue containing a BinarySet
func NewBinarySetAttribute(value [][]byte) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewListAttribute creates an DynamoDBAttributeValue containing a List
func NewListAttribute(value []DynamoDBAttributeValue) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewMapAttribute creates an DynamoDBAttributeValue containing a Map
func NewMapAttribute(value map[string]DynamoDBAttributeValue) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewNumberAttribute creates an DynamoDBAttributeValue containing a Number
func NewNumberAttribute(value string) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewNumberSetAttribute creates an DynamoDBAttributeValue containing a NumberSet
func NewNumberSetAttribute(value []string) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewNullAttribute creates an DynamoDBAttributeValue containing a Null
func NewNullAttribute() DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewStringAttribute creates an DynamoDBAttributeValue containing a String
func NewStringAttribute(value string) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// NewStringSetAttribute creates an DynamoDBAttributeValue containing a StringSet
func NewStringSetAttribute(value []string) DynamoDBAttributeValue {
	_ = "STUB: not implemented"
	return *new(DynamoDBAttributeValue)
}

// DynamoDBDataType specifies the type supported natively by DynamoDB for an attribute
type DynamoDBDataType int

const (
	DataTypeBinary DynamoDBDataType = iota
	DataTypeBoolean
	DataTypeBinarySet
	DataTypeList
	DataTypeMap
	DataTypeNumber
	DataTypeNumberSet
	DataTypeNull
	DataTypeString
	DataTypeStringSet
)

type anyValue interface{}

// UnsupportedDynamoDBTypeError is the error returned when trying to unmarshal a DynamoDB Attribute type not recognized by this library
type UnsupportedDynamoDBTypeError struct {
	Type string
}

func (e UnsupportedDynamoDBTypeError) Error() string { _ = "STUB: not implemented"; return "" }

// IncompatibleDynamoDBTypeError is the error passed in a panic when calling an accessor for an incompatible type
type IncompatibleDynamoDBTypeError struct {
	Requested DynamoDBDataType
	Actual    DynamoDBDataType
}

func (e IncompatibleDynamoDBTypeError) Error() string { _ = "STUB: not implemented"; return "" }

func (av *DynamoDBAttributeValue) ensureType(expectedType DynamoDBDataType) {
	_ = "STUB: not implemented"
	return
}

// MarshalJSON implements custom marshaling to be used by the standard json/encoding package
func (av DynamoDBAttributeValue) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalNull(target *DynamoDBAttributeValue) error { _ = "STUB: not implemented"; return nil }

func unmarshalString(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalBinary(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalBoolean(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalBinarySet(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalList(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalMap(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalNumber(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalNumberSet(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalStringSet(target *DynamoDBAttributeValue, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalDynamoDBAttributeValue(target *DynamoDBAttributeValue, typeLabel string, jsonValue interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON unmarshals a JSON description of this DynamoDBAttributeValue
func (av *DynamoDBAttributeValue) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalDynamoDBAttributeValueMap(target *DynamoDBAttributeValue, m map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
