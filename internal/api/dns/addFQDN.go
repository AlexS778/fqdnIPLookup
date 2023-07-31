package dnsapi

import "github.com/AlexS778/fqdnIPLookup/internal/models"

func (h handler) saveFQDNAndIPsToDatabase(fqdn string, ips []string) {
	tx := h.DB.Begin()

	fqdnRecord := models.FQDN{
		FQDN: fqdn,
	}
	tx.Create(&fqdnRecord)

	for _, ip := range ips {
		ipRecord := models.IP{
			FQDNID:  fqdnRecord.ID,
			Address: ip,
		}
		tx.Create(&ipRecord)
	}

	tx.Commit()
}
