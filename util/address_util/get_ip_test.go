package address_util

import (
	"github.com/cloudflare-fans/cf-ddns/bu_const"
	"testing"
)

func TestGetIPDNSType(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name        string
		args        args
		wantDnsType bu_const.DNSType
		wantErr     bool
	}{
		// test cases.
		{
			"local_addr_4",
			args{ipStr: "127.0.0.1"},
			bu_const.DNSTypeIPv4,
			false,
		},
		{
			"lan_addr_4",
			args{ipStr: "192.168.0.1"},
			bu_const.DNSTypeIPv4,
			false,
		},
		{
			"inet_addr_4",
			args{ipStr: "222.210.66.134"},
			bu_const.DNSTypeIPv4,
			false,
		},
		{
			"local_addr_6",
			args{ipStr: "::1"},
			bu_const.DNSTypeIPv6,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDnsType, err := GetIPDNSType(tt.args.ipStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIPDNSType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDnsType != tt.wantDnsType {
				t.Errorf("GetIPDNSType() gotDnsType = %v, want %v", gotDnsType, tt.wantDnsType)
			}
		})
	}
}
