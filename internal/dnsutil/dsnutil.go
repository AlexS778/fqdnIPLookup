package dnsutil

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func GetIPsByFQDN(FQDN string) ([]string, error) {
	var result []string

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	ips, err := resolver.LookupIP(context.Background(), "ip", FQDN)
	if err != nil {
		return nil, err
	}

	for _, ip := range ips {
		result = append(result, ip.String())
	}
	return result, nil
}

func GetFQDNsByIP(ipAdress string) ([]string, error) {
	ip := net.ParseIP(ipAdress)
	if ip == nil {
		return nil, fmt.Errorf("error parsing IP adress for %s", ipAdress)
	}

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	names, err := resolver.LookupAddr(context.Background(), ip.String())
	if err != nil {
		return nil, err
	}

	if len(names) == 0 {
		return nil, fmt.Errorf("no FQDNs found for IP %s", ipAdress)
	}

	return names, nil
}

func GetWHOISFromSLD(url string) (whoisparser.Domain, error) {
	sld := domainutil.DomainPrefix(url)
	if sld == "" {
		return whoisparser.Domain{}, fmt.Errorf("no second-leve domain found in url %s", url)
	}

	whoIsRaw, err := whois.Whois(sld)
	if err != nil {
		return whoisparser.Domain{}, err
	}

	result, err := whoisparser.Parse(whoIsRaw)
	if err != nil {
		return whoisparser.Domain{}, err
	}
	return *result.Domain, nil
}
