package bu_const

type IPType string

const (
	IPTypeIPv6Only IPType = "ipv6-only"
	IPTypeIPv4Only IPType = "ipv4-only"
	IPTypeIPv6     IPType = "ipv6"
	IPTypeIPv4     IPType = "ipv4"
)
