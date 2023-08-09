package dnsutil

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/AlexS778/fqdnIPLookup/models"
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

func GetWHOISFromSLD(url string) models.WhoISDTO {
	whoISDTONew := models.WhoISDTO{}
	sld := domainutil.DomainPrefix(url)
	if sld == "" {
		whoISDTONew.Error = fmt.Sprintf("no second-leve domain found in url %s", url)
		return whoISDTONew

	}

	whoIsRaw, err := whois.Whois(sld)
	if err != nil {
		whoISDTONew.Error = err.Error()
		return whoISDTONew
	}

	result, err := whoisparser.Parse(whoIsRaw)
	if err != nil {
		whoISDTONew.Error = err.Error()
		return whoISDTONew
	}

	whoISDTONew.Domain = *result.Domain
	if result.Registrar != nil {
		whoISDTONew.Registrar = *result.Registrar
	}
	if result.Registrant != nil {
		whoISDTONew.Registrant = *result.Registrant
	}
	if result.Administrative != nil {
		whoISDTONew.Administrative = *result.Administrative
	}
	if result.Technical != nil {
		whoISDTONew.Technical = *result.Technical
	}
	if result.Billing != nil {
		whoISDTONew.Billing = *result.Billing
	}

	return whoISDTONew
}
