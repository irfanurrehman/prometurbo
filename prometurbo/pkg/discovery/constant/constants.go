package constant

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

const (
	// MetricType
	Used = "used"

	// Internal matching property
	// The default namespace of entity property
	DefaultPropertyNamespace = "DEFAULT"

	// The attribute used for stitching with other probes (e.g., prometurbo) with app and vapp
	StitchingAttr string = "IP"

	// External matching property
	// The attribute used for stitching with other probes (e.g., prometurbo) with vm
	SupplyChainConstantIpAddress          = "ipAddress"
	SupplyChainConstantVirtualMachineData = "virtual_machine_data"

	VAppPrefix   = "vApp-"
	BizAppPrefix = "businessApp-"
)

// In most cases, the capacity for a commodity should be provided from the input JSON.
// For some commodities, the capacity may have a fixed default value, for example, the
// capacity for garbage collection time and db cache hit rate are all 100% by default.
type DefaultValue struct {
	Capacity float64
}

type CommodityTypeMap map[proto.CommodityDTO_CommodityType]DefaultValue

var EntityTypeMap = map[proto.EntityDTO_EntityType]CommodityTypeMap{
	proto.EntityDTO_APPLICATION:     SupportedAppCommodities,
	proto.EntityDTO_DATABASE_SERVER: SupportedDBCommodities,
}

var SupportedVMCommodities = CommodityTypeMap{
	proto.CommodityDTO_VCPU: {},
	proto.CommodityDTO_VMEM: {},
}

var SupportedAppCommodities = CommodityTypeMap{
	proto.CommodityDTO_TRANSACTION:     {},
	proto.CommodityDTO_RESPONSE_TIME:   {},
	proto.CommodityDTO_VCPU:            {},
	proto.CommodityDTO_VMEM:            {},
	proto.CommodityDTO_COLLECTION_TIME: {Capacity: 100.0},
	proto.CommodityDTO_HEAP:            {},
	proto.CommodityDTO_THREADS:         {},
}

var SupportedDBCommodities = CommodityTypeMap{
	proto.CommodityDTO_DB_CACHE_HIT_RATE: {Capacity: 100.0},
	proto.CommodityDTO_DB_MEM:            {},
	proto.CommodityDTO_CONNECTION:        {},
	proto.CommodityDTO_TRANSACTION:       {},
}
