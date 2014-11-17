// Code generated by protoc-gen-go.
// source: person.proto
// DO NOT EDIT!

/*
Package person is a generated protocol buffer package.

It is generated from these files:
	person.proto

It has these top-level messages:
	Person
	Phone
*/
package person

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Phone_PHONE_TYPE int32

const (
	Phone_MOBILE Phone_PHONE_TYPE = 1
	Phone_HOME   Phone_PHONE_TYPE = 2
)

var Phone_PHONE_TYPE_name = map[int32]string{
	1: "MOBILE",
	2: "HOME",
}
var Phone_PHONE_TYPE_value = map[string]int32{
	"MOBILE": 1,
	"HOME":   2,
}

func (x Phone_PHONE_TYPE) Enum() *Phone_PHONE_TYPE {
	p := new(Phone_PHONE_TYPE)
	*p = x
	return p
}
func (x Phone_PHONE_TYPE) String() string {
	return proto.EnumName(Phone_PHONE_TYPE_name, int32(x))
}
func (x *Phone_PHONE_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Phone_PHONE_TYPE_value, data, "Phone_PHONE_TYPE")
	if err != nil {
		return err
	}
	*x = Phone_PHONE_TYPE(value)
	return nil
}

type Person struct {
	Id               *int32                    `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string                   `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Email            *string                   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}

var extRange_Person = []proto.ExtensionRange{
	{10, 536870911},
}

func (*Person) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Person
}
func (m *Person) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Person) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

type Phone struct {
	Num              *string           `protobuf:"bytes,1,opt,name=num" json:"num,omitempty"`
	Type             *Phone_PHONE_TYPE `protobuf:"varint,2,opt,name=type,enum=Phone_PHONE_TYPE" json:"type,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}

func (m *Phone) GetNum() string {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return ""
}

func (m *Phone) GetType() Phone_PHONE_TYPE {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Phone_MOBILE
}

var E_Phone_Phones = &proto.ExtensionDesc{
	ExtendedType:  (*Person)(nil),
	ExtensionType: ([]*Phone)(nil),
	Field:         10,
	Name:          "Phone.phones",
	Tag:           "bytes,10,rep,name=phones",
}

func init() {
	proto.RegisterEnum("Phone_PHONE_TYPE", Phone_PHONE_TYPE_name, Phone_PHONE_TYPE_value)
	proto.RegisterExtension(E_Phone_Phones)
}
