// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CloudCostDTO.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The tenancy of an instance defines what hardware the instance is running on.
type Tenancy int32

const (
	// Instance runs on shared/default hardware.
	// This is typically the cheapest option.
	Tenancy_DEFAULT Tenancy = 1
	// Instance runs on single-tenant hardware.
	// That means your instance runs on a host that's separate from other customers,
	// but the host details are abstracted away, and you're not paying for the whole host.
	Tenancy_DEDICATED Tenancy = 2
	// Instance runs on a dedicated Host.
	// This means your instance runs on a specific host, and you are paying for the full host and
	// are responsible for managing it.
	Tenancy_HOST Tenancy = 3
)

var Tenancy_name = map[int32]string{
	1: "DEFAULT",
	2: "DEDICATED",
	3: "HOST",
}

var Tenancy_value = map[string]int32{
	"DEFAULT":   1,
	"DEDICATED": 2,
	"HOST":      3,
}

func (x Tenancy) Enum() *Tenancy {
	p := new(Tenancy)
	*p = x
	return p
}

func (x Tenancy) String() string {
	return proto.EnumName(Tenancy_name, int32(x))
}

func (x *Tenancy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Tenancy_value, data, "Tenancy")
	if err != nil {
		return err
	}
	*x = Tenancy(value)
	return nil
}

func (Tenancy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{0}
}

// The supported operating systems.
type OSType int32

const (
	OSType_UNKNOWN_OS OSType = 0
	// Unix OS.
	OSType_LINUX                     OSType = 2
	OSType_SUSE                      OSType = 3
	OSType_RHEL                      OSType = 4
	OSType_LINUX_WITH_SQL_ENTERPRISE OSType = 5
	OSType_LINUX_WITH_SQL_STANDARD   OSType = 6
	OSType_LINUX_WITH_SQL_WEB        OSType = 7
	// Windows OS.
	OSType_WINDOWS                     OSType = 20
	OSType_WINDOWS_WITH_SQL_STANDARD   OSType = 21
	OSType_WINDOWS_WITH_SQL_WEB        OSType = 22
	OSType_WINDOWS_WITH_SQL_ENTERPRISE OSType = 23
	OSType_WINDOWS_BYOL                OSType = 24
)

var OSType_name = map[int32]string{
	0:  "UNKNOWN_OS",
	2:  "LINUX",
	3:  "SUSE",
	4:  "RHEL",
	5:  "LINUX_WITH_SQL_ENTERPRISE",
	6:  "LINUX_WITH_SQL_STANDARD",
	7:  "LINUX_WITH_SQL_WEB",
	20: "WINDOWS",
	21: "WINDOWS_WITH_SQL_STANDARD",
	22: "WINDOWS_WITH_SQL_WEB",
	23: "WINDOWS_WITH_SQL_ENTERPRISE",
	24: "WINDOWS_BYOL",
}

var OSType_value = map[string]int32{
	"UNKNOWN_OS":                  0,
	"LINUX":                       2,
	"SUSE":                        3,
	"RHEL":                        4,
	"LINUX_WITH_SQL_ENTERPRISE":   5,
	"LINUX_WITH_SQL_STANDARD":     6,
	"LINUX_WITH_SQL_WEB":          7,
	"WINDOWS":                     20,
	"WINDOWS_WITH_SQL_STANDARD":   21,
	"WINDOWS_WITH_SQL_WEB":        22,
	"WINDOWS_WITH_SQL_ENTERPRISE": 23,
	"WINDOWS_BYOL":                24,
}

func (x OSType) Enum() *OSType {
	p := new(OSType)
	*p = x
	return p
}

func (x OSType) String() string {
	return proto.EnumName(OSType_name, int32(x))
}

func (x *OSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OSType_value, data, "OSType")
	if err != nil {
		return err
	}
	*x = OSType(value)
	return nil
}

func (OSType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{1}
}

// The engine for a database tier.
// This is an enum to save on space - and also because
// the list of supported engines across cloud providers is pretty small.
type DatabaseEngine int32

const (
	DatabaseEngine_UNKNOWN    DatabaseEngine = 0
	DatabaseEngine_MYSQL      DatabaseEngine = 1
	DatabaseEngine_MARIADB    DatabaseEngine = 2
	DatabaseEngine_POSTGRESQL DatabaseEngine = 3
	DatabaseEngine_ORACLE     DatabaseEngine = 4
	DatabaseEngine_SQL_SERVER DatabaseEngine = 5
	DatabaseEngine_AURORA     DatabaseEngine = 6
)

var DatabaseEngine_name = map[int32]string{
	0: "UNKNOWN",
	1: "MYSQL",
	2: "MARIADB",
	3: "POSTGRESQL",
	4: "ORACLE",
	5: "SQL_SERVER",
	6: "AURORA",
}

var DatabaseEngine_value = map[string]int32{
	"UNKNOWN":    0,
	"MYSQL":      1,
	"MARIADB":    2,
	"POSTGRESQL": 3,
	"ORACLE":     4,
	"SQL_SERVER": 5,
	"AURORA":     6,
}

func (x DatabaseEngine) Enum() *DatabaseEngine {
	p := new(DatabaseEngine)
	*p = x
	return p
}

func (x DatabaseEngine) String() string {
	return proto.EnumName(DatabaseEngine_name, int32(x))
}

func (x *DatabaseEngine) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DatabaseEngine_value, data, "DatabaseEngine")
	if err != nil {
		return err
	}
	*x = DatabaseEngine(value)
	return nil
}

func (DatabaseEngine) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{2}
}

// LicenseModel describes all supported license models in cloud projects.
type LicenseModel int32

const (
	LicenseModel_BRING_YOUR_OWN_LICENSE LicenseModel = 0
	LicenseModel_LICENSE_INCLUDED       LicenseModel = 1
	LicenseModel_NO_LICENSE_REQUIRED    LicenseModel = 2
)

var LicenseModel_name = map[int32]string{
	0: "BRING_YOUR_OWN_LICENSE",
	1: "LICENSE_INCLUDED",
	2: "NO_LICENSE_REQUIRED",
}

var LicenseModel_value = map[string]int32{
	"BRING_YOUR_OWN_LICENSE": 0,
	"LICENSE_INCLUDED":       1,
	"NO_LICENSE_REQUIRED":    2,
}

func (x LicenseModel) Enum() *LicenseModel {
	p := new(LicenseModel)
	*p = x
	return p
}

func (x LicenseModel) String() string {
	return proto.EnumName(LicenseModel_name, int32(x))
}

func (x *LicenseModel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LicenseModel_value, data, "LicenseModel")
	if err != nil {
		return err
	}
	*x = LicenseModel(value)
	return nil
}

func (LicenseModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{3}
}

// DeploymentType describes all supported deployment types by cloud probes.
type DeploymentType int32

const (
	DeploymentType_SINGLE_AZ DeploymentType = 0
	DeploymentType_MULTI_AZ  DeploymentType = 1
)

var DeploymentType_name = map[int32]string{
	0: "SINGLE_AZ",
	1: "MULTI_AZ",
}

var DeploymentType_value = map[string]int32{
	"SINGLE_AZ": 0,
	"MULTI_AZ":  1,
}

func (x DeploymentType) Enum() *DeploymentType {
	p := new(DeploymentType)
	*p = x
	return p
}

func (x DeploymentType) String() string {
	return proto.EnumName(DeploymentType_name, int32(x))
}

func (x *DeploymentType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeploymentType_value, data, "DeploymentType")
	if err != nil {
		return err
	}
	*x = DeploymentType(value)
	return nil
}

func (DeploymentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{4}
}

// The edition of a database engine.
// The edition enum is closely related to the DatabaseEngine enum, and in the future it may be
// worth it to have a separate "database identifier" message that forbids illegal
// engine-edition combinations. For now, there are only two database engines with editions,
// so this seems manageable.
type DatabaseEdition int32

const (
	DatabaseEdition_NONE DatabaseEdition = 0
	// The available database editions when DatabaseEngine is ORACLE.
	// These values are meaningless with any other engine.
	DatabaseEdition_ORACLE_ENTERPRISE DatabaseEdition = 1
	DatabaseEdition_ORACLE_STANDARD   DatabaseEdition = 2
	DatabaseEdition_ORACLE_STANDARD_1 DatabaseEdition = 3
	DatabaseEdition_ORACLE_STANDARD_2 DatabaseEdition = 4
	// The available database editions when DatabaseEngine is SQL_SERVER.
	// These values are meaningless with any other engine.
	DatabaseEdition_SQL_SERVER_ENTERPRISE DatabaseEdition = 10
	DatabaseEdition_SQL_SERVER_STANDARD   DatabaseEdition = 11
	DatabaseEdition_SQL_SERVER_WEB        DatabaseEdition = 12
	DatabaseEdition_SQL_SERVER_EXPRESS    DatabaseEdition = 13
)

var DatabaseEdition_name = map[int32]string{
	0:  "NONE",
	1:  "ORACLE_ENTERPRISE",
	2:  "ORACLE_STANDARD",
	3:  "ORACLE_STANDARD_1",
	4:  "ORACLE_STANDARD_2",
	10: "SQL_SERVER_ENTERPRISE",
	11: "SQL_SERVER_STANDARD",
	12: "SQL_SERVER_WEB",
	13: "SQL_SERVER_EXPRESS",
}

var DatabaseEdition_value = map[string]int32{
	"NONE":                  0,
	"ORACLE_ENTERPRISE":     1,
	"ORACLE_STANDARD":       2,
	"ORACLE_STANDARD_1":     3,
	"ORACLE_STANDARD_2":     4,
	"SQL_SERVER_ENTERPRISE": 10,
	"SQL_SERVER_STANDARD":   11,
	"SQL_SERVER_WEB":        12,
	"SQL_SERVER_EXPRESS":    13,
}

func (x DatabaseEdition) Enum() *DatabaseEdition {
	p := new(DatabaseEdition)
	*p = x
	return p
}

func (x DatabaseEdition) String() string {
	return proto.EnumName(DatabaseEdition_name, int32(x))
}

func (x *DatabaseEdition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DatabaseEdition_value, data, "DatabaseEdition")
	if err != nil {
		return err
	}
	*x = DatabaseEdition(value)
	return nil
}

func (DatabaseEdition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{5}
}

type ReservedInstanceType_OfferingClass int32

const (
	// Most of the attributes of a standard reserved instance are "fixed" at the time it's
	// bought.
	ReservedInstanceType_STANDARD ReservedInstanceType_OfferingClass = 1
	// A convertible reserved instance can be exchanged for a different
	// instance type. Not all service providers offer convertible reserved
	// instances.
	ReservedInstanceType_CONVERTIBLE ReservedInstanceType_OfferingClass = 2
)

var ReservedInstanceType_OfferingClass_name = map[int32]string{
	1: "STANDARD",
	2: "CONVERTIBLE",
}

var ReservedInstanceType_OfferingClass_value = map[string]int32{
	"STANDARD":    1,
	"CONVERTIBLE": 2,
}

func (x ReservedInstanceType_OfferingClass) Enum() *ReservedInstanceType_OfferingClass {
	p := new(ReservedInstanceType_OfferingClass)
	*p = x
	return p
}

func (x ReservedInstanceType_OfferingClass) String() string {
	return proto.EnumName(ReservedInstanceType_OfferingClass_name, int32(x))
}

func (x *ReservedInstanceType_OfferingClass) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReservedInstanceType_OfferingClass_value, data, "ReservedInstanceType_OfferingClass")
	if err != nil {
		return err
	}
	*x = ReservedInstanceType_OfferingClass(value)
	return nil
}

func (ReservedInstanceType_OfferingClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{0, 0}
}

type ReservedInstanceType_PaymentOption int32

const (
	// The user must pay the entire price of this instance upfront. There is no recurring
	// cost.
	// (e.g. $10000.00 upfront for the year)
	ReservedInstanceType_ALL_UPFRONT ReservedInstanceType_PaymentOption = 1
	// The user must pay some part of the instance price upfront, and the rest over time.
	// (e.g. $1000.00 upfront, and $0.5 per instance-hour afterwards).
	ReservedInstanceType_PARTIAL_UPFRONT ReservedInstanceType_PaymentOption = 2
	// The entire price of the instance is recurring
	// (e.g. $0.7 per instance-hour)
	ReservedInstanceType_NO_UPFRONT ReservedInstanceType_PaymentOption = 3
)

var ReservedInstanceType_PaymentOption_name = map[int32]string{
	1: "ALL_UPFRONT",
	2: "PARTIAL_UPFRONT",
	3: "NO_UPFRONT",
}

var ReservedInstanceType_PaymentOption_value = map[string]int32{
	"ALL_UPFRONT":     1,
	"PARTIAL_UPFRONT": 2,
	"NO_UPFRONT":      3,
}

func (x ReservedInstanceType_PaymentOption) Enum() *ReservedInstanceType_PaymentOption {
	p := new(ReservedInstanceType_PaymentOption)
	*p = x
	return p
}

func (x ReservedInstanceType_PaymentOption) String() string {
	return proto.EnumName(ReservedInstanceType_PaymentOption_name, int32(x))
}

func (x *ReservedInstanceType_PaymentOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReservedInstanceType_PaymentOption_value, data, "ReservedInstanceType_PaymentOption")
	if err != nil {
		return err
	}
	*x = ReservedInstanceType_PaymentOption(value)
	return nil
}

func (ReservedInstanceType_PaymentOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{0, 1}
}

// Identifies the type of "reservation" for the instance, and the
// payment conditions.
type ReservedInstanceType struct {
	// The type of offering.
	OfferingClass *ReservedInstanceType_OfferingClass `protobuf:"varint,1,opt,name=offering_class,json=offeringClass,enum=common_dto.ReservedInstanceType_OfferingClass" json:"offering_class,omitempty"`
	// The payment option for this reserved instance.
	PaymentOption *ReservedInstanceType_PaymentOption `protobuf:"varint,2,opt,name=payment_option,json=paymentOption,enum=common_dto.ReservedInstanceType_PaymentOption" json:"payment_option,omitempty"`
	// The term, in years.
	TermYears            *uint32  `protobuf:"varint,3,opt,name=term_years,json=termYears" json:"term_years,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReservedInstanceType) Reset()         { *m = ReservedInstanceType{} }
func (m *ReservedInstanceType) String() string { return proto.CompactTextString(m) }
func (*ReservedInstanceType) ProtoMessage()    {}
func (*ReservedInstanceType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{0}
}

func (m *ReservedInstanceType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReservedInstanceType.Unmarshal(m, b)
}
func (m *ReservedInstanceType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReservedInstanceType.Marshal(b, m, deterministic)
}
func (m *ReservedInstanceType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReservedInstanceType.Merge(m, src)
}
func (m *ReservedInstanceType) XXX_Size() int {
	return xxx_messageInfo_ReservedInstanceType.Size(m)
}
func (m *ReservedInstanceType) XXX_DiscardUnknown() {
	xxx_messageInfo_ReservedInstanceType.DiscardUnknown(m)
}

var xxx_messageInfo_ReservedInstanceType proto.InternalMessageInfo

func (m *ReservedInstanceType) GetOfferingClass() ReservedInstanceType_OfferingClass {
	if m != nil && m.OfferingClass != nil {
		return *m.OfferingClass
	}
	return ReservedInstanceType_STANDARD
}

func (m *ReservedInstanceType) GetPaymentOption() ReservedInstanceType_PaymentOption {
	if m != nil && m.PaymentOption != nil {
		return *m.PaymentOption
	}
	return ReservedInstanceType_ALL_UPFRONT
}

func (m *ReservedInstanceType) GetTermYears() uint32 {
	if m != nil && m.TermYears != nil {
		return *m.TermYears
	}
	return 0
}

// An amount of money, expressed in some currency.
type CurrencyAmount struct {
	// The currency in which the amount is expressed.
	// This is the ISO 4217 numeric code.
	// The default (840) is the USD currency code.
	//
	// We use the ISO 4217 standard so that in the future it would be easier to integrate
	// with JSR 354: Money and Currency API.
	Currency *int32 `protobuf:"varint,1,opt,name=currency,def=840" json:"currency,omitempty"`
	// The value, in the currency.
	// This should be non-negative, with 0 representing "free".
	Amount               *float64 `protobuf:"fixed64,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrencyAmount) Reset()         { *m = CurrencyAmount{} }
func (m *CurrencyAmount) String() string { return proto.CompactTextString(m) }
func (*CurrencyAmount) ProtoMessage()    {}
func (*CurrencyAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{1}
}

func (m *CurrencyAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyAmount.Unmarshal(m, b)
}
func (m *CurrencyAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyAmount.Marshal(b, m, deterministic)
}
func (m *CurrencyAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyAmount.Merge(m, src)
}
func (m *CurrencyAmount) XXX_Size() int {
	return xxx_messageInfo_CurrencyAmount.Size(m)
}
func (m *CurrencyAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyAmount.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyAmount proto.InternalMessageInfo

const Default_CurrencyAmount_Currency int32 = 840

func (m *CurrencyAmount) GetCurrency() int32 {
	if m != nil && m.Currency != nil {
		return *m.Currency
	}
	return Default_CurrencyAmount_Currency
}

func (m *CurrencyAmount) GetAmount() float64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

// The ReservedInstanceSpec describes a solution offered by the cloud provider that allows the customer
// to buy in advance a number of compute instances, for a discounted price. Usually those solutions
// have long terms like 1 or 3 years.
type ReservedInstanceSpec struct {
	// The type of the reserved instance.
	Type *ReservedInstanceType `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// The tenancy of the reserved instance.
	Tenancy *Tenancy `protobuf:"varint,2,opt,name=tenancy,enum=common_dto.Tenancy" json:"tenancy,omitempty"`
	// The operating system of the reserved instance.
	Os *OSType `protobuf:"varint,3,opt,name=os,enum=common_dto.OSType" json:"os,omitempty"`
	// The entity profile of the reserved instance is using, such as t2.large.
	Tier *EntityDTO `protobuf:"bytes,4,opt,name=tier" json:"tier,omitempty"`
	// The region of the reserved instance.
	Region               *EntityDTO `protobuf:"bytes,5,opt,name=region" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReservedInstanceSpec) Reset()         { *m = ReservedInstanceSpec{} }
func (m *ReservedInstanceSpec) String() string { return proto.CompactTextString(m) }
func (*ReservedInstanceSpec) ProtoMessage()    {}
func (*ReservedInstanceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e61221a66d0df59, []int{2}
}

func (m *ReservedInstanceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReservedInstanceSpec.Unmarshal(m, b)
}
func (m *ReservedInstanceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReservedInstanceSpec.Marshal(b, m, deterministic)
}
func (m *ReservedInstanceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReservedInstanceSpec.Merge(m, src)
}
func (m *ReservedInstanceSpec) XXX_Size() int {
	return xxx_messageInfo_ReservedInstanceSpec.Size(m)
}
func (m *ReservedInstanceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ReservedInstanceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ReservedInstanceSpec proto.InternalMessageInfo

func (m *ReservedInstanceSpec) GetType() *ReservedInstanceType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ReservedInstanceSpec) GetTenancy() Tenancy {
	if m != nil && m.Tenancy != nil {
		return *m.Tenancy
	}
	return Tenancy_DEFAULT
}

func (m *ReservedInstanceSpec) GetOs() OSType {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return OSType_UNKNOWN_OS
}

func (m *ReservedInstanceSpec) GetTier() *EntityDTO {
	if m != nil {
		return m.Tier
	}
	return nil
}

func (m *ReservedInstanceSpec) GetRegion() *EntityDTO {
	if m != nil {
		return m.Region
	}
	return nil
}

func init() {
	proto.RegisterEnum("common_dto.Tenancy", Tenancy_name, Tenancy_value)
	proto.RegisterEnum("common_dto.OSType", OSType_name, OSType_value)
	proto.RegisterEnum("common_dto.DatabaseEngine", DatabaseEngine_name, DatabaseEngine_value)
	proto.RegisterEnum("common_dto.LicenseModel", LicenseModel_name, LicenseModel_value)
	proto.RegisterEnum("common_dto.DeploymentType", DeploymentType_name, DeploymentType_value)
	proto.RegisterEnum("common_dto.DatabaseEdition", DatabaseEdition_name, DatabaseEdition_value)
	proto.RegisterEnum("common_dto.ReservedInstanceType_OfferingClass", ReservedInstanceType_OfferingClass_name, ReservedInstanceType_OfferingClass_value)
	proto.RegisterEnum("common_dto.ReservedInstanceType_PaymentOption", ReservedInstanceType_PaymentOption_name, ReservedInstanceType_PaymentOption_value)
	proto.RegisterType((*ReservedInstanceType)(nil), "common_dto.ReservedInstanceType")
	proto.RegisterType((*CurrencyAmount)(nil), "common_dto.CurrencyAmount")
	proto.RegisterType((*ReservedInstanceSpec)(nil), "common_dto.ReservedInstanceSpec")
}

func init() { proto.RegisterFile("CloudCostDTO.proto", fileDescriptor_6e61221a66d0df59) }

var fileDescriptor_6e61221a66d0df59 = []byte{
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6f, 0xdb, 0xc6,
	0x13, 0x35, 0xa9, 0x3f, 0xb6, 0xc7, 0x96, 0xbc, 0xbf, 0xf5, 0x3f, 0x25, 0x41, 0x60, 0xff, 0x74,
	0x4a, 0x05, 0x98, 0x6d, 0x8d, 0x1c, 0x8a, 0xde, 0x28, 0x72, 0x63, 0x13, 0xa5, 0x77, 0xe5, 0x25,
	0x19, 0x45, 0xbd, 0x10, 0x8c, 0xb4, 0x36, 0x84, 0x8a, 0x5c, 0x82, 0xa4, 0x03, 0xe8, 0xde, 0xaf,
	0xd8, 0x4b, 0xbf, 0x49, 0x6f, 0xc5, 0x92, 0xb4, 0xcc, 0xc4, 0x46, 0xdb, 0xdb, 0xce, 0xbc, 0x37,
	0xb3, 0xb3, 0xef, 0xed, 0x00, 0xb6, 0x56, 0xf2, 0x61, 0x61, 0xc9, 0xbc, 0xb0, 0x7d, 0x66, 0xa4,
	0x99, 0x2c, 0x24, 0x86, 0xb9, 0x8c, 0x63, 0x99, 0x84, 0x8b, 0x42, 0xbe, 0x3e, 0xb0, 0xca, 0xf3,
	0x06, 0x1c, 0xfe, 0xa9, 0xc3, 0x11, 0x17, 0xb9, 0xc8, 0xbe, 0x88, 0x85, 0x93, 0xe4, 0x45, 0x94,
	0xcc, 0x85, 0xbf, 0x4e, 0x05, 0x0e, 0xa0, 0x2f, 0xef, 0xee, 0x44, 0xb6, 0x4c, 0xee, 0xc3, 0xf9,
	0x2a, 0xca, 0xf3, 0x81, 0x76, 0xae, 0xbd, 0xeb, 0x5f, 0x1a, 0xc6, 0x53, 0x3b, 0xe3, 0xa5, 0x4a,
	0x83, 0xd5, 0x65, 0x96, 0xaa, 0xe2, 0x3d, 0xd9, 0x0c, 0x55, 0xdb, 0x34, 0x5a, 0xc7, 0x22, 0x29,
	0x42, 0x99, 0x16, 0x4b, 0x99, 0x0c, 0xf4, 0xff, 0xd8, 0x76, 0x52, 0x95, 0xb1, 0xb2, 0x8a, 0xf7,
	0xd2, 0x66, 0x88, 0xdf, 0x02, 0x14, 0x22, 0x8b, 0xc3, 0xb5, 0x88, 0xb2, 0x7c, 0xd0, 0x3a, 0xd7,
	0xde, 0xf5, 0xf8, 0xae, 0xca, 0xcc, 0x54, 0x62, 0x68, 0x40, 0xef, 0xab, 0xa9, 0xf0, 0x3e, 0xec,
	0x78, 0xbe, 0x49, 0x6d, 0x93, 0xdb, 0x48, 0xc3, 0x07, 0xb0, 0x67, 0x31, 0xfa, 0x91, 0x70, 0xdf,
	0x19, 0xbb, 0x04, 0xe9, 0x43, 0x02, 0xbd, 0xaf, 0xae, 0x53, 0x0c, 0xd3, 0x75, 0xc3, 0x60, 0xf2,
	0x81, 0x33, 0xea, 0x23, 0x0d, 0x1f, 0xc2, 0xc1, 0xc4, 0xe4, 0xbe, 0x63, 0x3e, 0x25, 0x75, 0xdc,
	0x07, 0xa0, 0x6c, 0x13, 0xb7, 0x86, 0x0e, 0xf4, 0xad, 0x87, 0x2c, 0x13, 0xc9, 0x7c, 0x6d, 0xc6,
	0xf2, 0x21, 0x29, 0xf0, 0x19, 0xec, 0xcc, 0xeb, 0x4c, 0xa9, 0x67, 0xe7, 0xe7, 0xd6, 0x4f, 0xef,
	0x7f, 0xe0, 0x9b, 0x24, 0x3e, 0x81, 0x6e, 0x54, 0x52, 0x4b, 0x5d, 0x34, 0x5e, 0x47, 0xc3, 0xbf,
	0xb4, 0xe7, 0x3e, 0x79, 0xa9, 0x98, 0xe3, 0xf7, 0xd0, 0x2e, 0xd6, 0xa9, 0x28, 0xbb, 0xed, 0x5d,
	0x9e, 0xff, 0x9b, 0x8c, 0xbc, 0x64, 0xe3, 0x0b, 0xd8, 0x2e, 0x44, 0x12, 0xa9, 0x31, 0x2a, 0xfd,
	0x0f, 0x9b, 0x85, 0x7e, 0x05, 0xf1, 0x47, 0x0e, 0x1e, 0x82, 0x2e, 0x2b, 0x59, 0xfb, 0x97, 0xb8,
	0xc9, 0x64, 0x5e, 0xd9, 0x54, 0x97, 0x39, 0xfe, 0x0e, 0xda, 0xc5, 0x52, 0x64, 0x83, 0x76, 0x39,
	0xc8, 0x71, 0x93, 0x45, 0x92, 0x62, 0x59, 0xac, 0x6d, 0x9f, 0xf1, 0x92, 0x82, 0x2f, 0xa0, 0x9b,
	0x89, 0x7b, 0x65, 0x7e, 0xe7, 0x9f, 0xc8, 0x35, 0x69, 0xf4, 0x3d, 0x6c, 0xd7, 0x13, 0xe1, 0x3d,
	0xd8, 0xb6, 0xc9, 0x07, 0x33, 0x70, 0x95, 0x07, 0x3d, 0xd8, 0xb5, 0x89, 0xed, 0x58, 0xa6, 0x4f,
	0x6c, 0xa4, 0xe3, 0x1d, 0x68, 0x5f, 0x33, 0xcf, 0x47, 0xad, 0xd1, 0xef, 0x3a, 0x74, 0xab, 0xc9,
	0x94, 0x25, 0x01, 0xfd, 0x85, 0xb2, 0x29, 0x0d, 0x99, 0x87, 0xb6, 0xf0, 0x2e, 0x74, 0x5c, 0x87,
	0x06, 0x9f, 0x2a, 0xbe, 0x17, 0x78, 0x04, 0xb5, 0xd4, 0x89, 0x5f, 0x13, 0x17, 0xb5, 0xf1, 0x5b,
	0x78, 0x55, 0xc2, 0xe1, 0xd4, 0xf1, 0xaf, 0x43, 0xef, 0xd6, 0x0d, 0x09, 0xf5, 0x09, 0x9f, 0x70,
	0xc7, 0x23, 0xa8, 0x83, 0xdf, 0xc0, 0xe9, 0x37, 0xf0, 0xe6, 0x17, 0x75, 0xf1, 0x09, 0xe0, 0x6f,
	0xc0, 0x29, 0x19, 0xa3, 0x6d, 0x35, 0xf3, 0xd4, 0xa1, 0x36, 0x9b, 0x7a, 0xe8, 0x48, 0x5d, 0x50,
	0x07, 0x2f, 0xf4, 0x38, 0xc6, 0x03, 0x38, 0x7a, 0x06, 0xab, 0x2e, 0x27, 0xf8, 0x0c, 0xde, 0x3c,
	0x43, 0x1a, 0xb3, 0x9d, 0x62, 0x04, 0xfb, 0x8f, 0x84, 0xf1, 0x8c, 0xb9, 0x68, 0x30, 0x8a, 0xa1,
	0x6f, 0x47, 0x45, 0xf4, 0x39, 0xca, 0x05, 0x49, 0xee, 0x97, 0x89, 0x50, 0xa3, 0xd4, 0x6a, 0x54,
	0x52, 0xdc, 0xcc, 0xbc, 0x5b, 0x17, 0x69, 0x2a, 0x7f, 0x63, 0x72, 0xc7, 0xb4, 0xc7, 0xd5, 0x2f,
	0x9e, 0x30, 0xcf, 0xbf, 0xe2, 0x44, 0x81, 0x2d, 0x0c, 0xd0, 0x65, 0xdc, 0xb4, 0x5c, 0x82, 0xda,
	0x0a, 0x2b, 0x27, 0x26, 0xfc, 0x23, 0xe1, 0xa8, 0xa3, 0x30, 0x33, 0xe0, 0x8c, 0x9b, 0xa8, 0x3b,
	0x9a, 0xc1, 0xbe, 0xbb, 0x9c, 0x8b, 0x24, 0x17, 0x37, 0x72, 0x21, 0x56, 0xf8, 0x35, 0x9c, 0x8c,
	0xb9, 0x43, 0xaf, 0xc2, 0x19, 0x0b, 0x78, 0xa8, 0x1c, 0x70, 0x1d, 0x8b, 0x50, 0x8f, 0xa0, 0x2d,
	0x7c, 0x04, 0xa8, 0x0e, 0x42, 0x87, 0x5a, 0x6e, 0x60, 0x13, 0xb5, 0x87, 0xa7, 0x70, 0x48, 0xd9,
	0x23, 0x2b, 0xe4, 0xe4, 0x36, 0x70, 0xb8, 0xb2, 0x76, 0x74, 0x01, 0x7d, 0x5b, 0xa4, 0x2b, 0x59,
	0xae, 0x64, 0xe9, 0x6b, 0x0f, 0x76, 0x3d, 0x87, 0x5e, 0xb9, 0x24, 0x34, 0x7f, 0x45, 0x5b, 0x6a,
	0x9f, 0x6f, 0x02, 0xd7, 0x77, 0x54, 0xa4, 0x8d, 0xfe, 0xd0, 0xe0, 0x60, 0xf3, 0xf2, 0xc5, 0xb2,
	0xdc, 0xe0, 0x1d, 0x68, 0x53, 0x46, 0xd5, 0xdd, 0xc7, 0xf0, 0xbf, 0xea, 0x3d, 0x4d, 0xfd, 0xca,
	0x8d, 0xae, 0xd3, 0x1b, 0x3f, 0xf4, 0x06, 0xf7, 0x31, 0x19, 0xfe, 0x88, 0x5a, 0x2f, 0xa5, 0x2f,
	0x51, 0x1b, 0xbf, 0x82, 0xe3, 0x27, 0x75, 0x9a, 0xdd, 0x41, 0x3d, 0xad, 0x01, 0x6d, 0x6e, 0xd8,
	0xc3, 0x18, 0xfa, 0x0d, 0x40, 0x79, 0xbd, 0xaf, 0x7e, 0x52, 0xb3, 0xcf, 0xa7, 0x09, 0x27, 0x9e,
	0x87, 0x7a, 0xe3, 0xff, 0xc3, 0xd9, 0x5c, 0xc6, 0xc6, 0x97, 0xb8, 0x78, 0xc8, 0x3e, 0x4b, 0x23,
	0x5d, 0x45, 0xc5, 0x9d, 0xcc, 0x62, 0x23, 0x5f, 0xfc, 0x56, 0x6f, 0xd0, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x42, 0x8f, 0x43, 0xc7, 0x01, 0x06, 0x00, 0x00,
}
