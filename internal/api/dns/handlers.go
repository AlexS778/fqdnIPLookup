package dnsapi

import (
	"log"

	"github.com/AlexS778/fqdnIPLookup/internal/dnsutil"
	"github.com/AlexS778/fqdnIPLookup/models"
	"github.com/gin-gonic/gin"
)

func (h handler) FQDNToIPHanlder(c *gin.Context) {
	var arrayOfFQDNs []string
	if err := c.ShouldBindJSON(&arrayOfFQDNs); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		log.Printf("Error binding json %s", err)
		return
	}

	ipAdresses := make(map[string][]string)
	for _, fqdn := range arrayOfFQDNs {
		res, err := dnsutil.GetIPsByFQDN(fqdn)
		if err != nil {
			ipAdresses[fqdn] = []string{err.Error()}
		} else {
			ipAdresses[fqdn] = res
			h.saveFQDNAndIPsToDatabase(fqdn, ipAdresses[fqdn])
		}

	}

	c.JSON(200, ipAdresses)
}

func (h handler) IPToFQDN(c *gin.Context) {
	var arrayOfIPs []string
	if err := c.ShouldBindJSON(&arrayOfIPs); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		log.Printf("Error binding json %s", err)
		return
	}

	fqdns := make(map[string][]string)
	for _, ip := range arrayOfIPs {
		res, err := dnsutil.GetFQDNsByIP(ip)
		if err != nil {
			fqdns[ip] = []string{err.Error()}
		} else {
			fqdns[ip] = res
		}
	}

	c.JSON(200, fqdns)
}

func (h handler) GetWhoIsData(c *gin.Context) {
	var slds []string
	if err := c.ShouldBindJSON(&slds); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		log.Printf("Error binding json %s", err)
		return
	}

	whoisData := make(map[string]models.WhoISDTO)
	for _, sld := range slds {
		res := dnsutil.GetWHOISFromSLD(sld)
		whoisData[sld] = res
	}

	c.JSON(200, whoisData)
}
