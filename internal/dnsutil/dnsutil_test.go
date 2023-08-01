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
	_, err := GetWHOISFromSLD(tldDomain)
	assert.NotEqual(t, err, nil)
	invalidSLD := "asdasdqweqweqwewqe.com"
	_, err = GetWHOISFromSLD(invalidSLD)
	assert.NotEqual(t, err, nil)
	validSLD := "ads.yahoo.com"
	_, err = GetWHOISFromSLD(validSLD)
	assert.Equal(t, err, nil)
}
