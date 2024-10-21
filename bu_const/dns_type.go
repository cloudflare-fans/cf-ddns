package bu_const

type DNSType string

const (
	DNSTypeIPv4    DNSType = "A"
	DNSTypeIPv6    DNSType = "AAAA"
	DNSTypeInvalid DNSType = ""
)
