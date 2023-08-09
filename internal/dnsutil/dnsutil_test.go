package dnsutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIPsBYFQDN(t *testing.T) {
	validFQDN := "mail.yahoo.com"
	_, err := GetIPsByFQDN(validFQDN)
	assert.Equal(t, err, nil)
	invalidFQDN := "aaa"
	_, err = GetIPsByFQDN(invalidFQDN)
	assert.NotEqual(t, err, nil)
}

func TestGetFQDNsByIP(t *testing.T) {
	invalidIP := "192.1111"
	_, err := GetFQDNsByIP(invalidIP)
	assert.NotEqual(t, err, nil)
	wrongIP := "0000:000:0000::0"
	_, err = GetFQDNsByIP(wrongIP)
	assert.NotEqual(t, err, nil)
	goodIp := "69.147.88.7"
	_, err = GetFQDNsByIP(goodIp)
	assert.Equal(t, err, nil)
}

func TestGetWHOISFromSLD(t *testing.T) {
	tldDomain := ""
	res := GetWHOISFromSLD(tldDomain)
	assert.NotNil(t, res.Error)
	invalidSLD := "asdasdqweqweqwewqeaasasdasd.com"
	res = GetWHOISFromSLD(invalidSLD)
	assert.NotNil(t, res.Error)
	validSLD := "ads.yahoo.com"
	res = GetWHOISFromSLD(validSLD)
	assert.Nil(t, res.Error)
}
