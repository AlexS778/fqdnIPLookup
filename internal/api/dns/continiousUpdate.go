package dnsapi

import (
	"errors"
	"log"
	"time"

	"github.com/AlexS778/fqdnIPLookup/internal/dnsutil"
	"github.com/AlexS778/fqdnIPLookup/internal/models"
	"github.com/AlexS778/fqdnIPLookup/internal/utils"
	"gorm.io/gorm"
)

func ContinuousUpdate(db *gorm.DB, waitTime time.Duration) {
	for {
		fqdns, err := getAllFQDNsFromDB(db)
		if err != nil {
			log.Println("Error fetching FQDNs from database:", err)
		}

		ipAdresses := make(map[string][]string)
		for fqdn := range fqdns {
			res, err := dnsutil.GetIPsByFQDN(fqdn)
			if err != nil {
				ipAdresses[fqdn] = []string{err.Error()}
			} else {
				ipAdresses[fqdn] = res
			}
		}

		res := utils.FindDifferentValues(fqdns, ipAdresses)
		saveFQDNAndIPsToDatabase(db, res)

		time.Sleep(waitTime)
	}
}

func getAllFQDNsFromDB(db *gorm.DB) (map[string][]string, error) {
	var fqdns []models.FQDN
	if err := db.Preload("IPs").Find(&fqdns).Error; err != nil {
		return nil, err
	}

	log.Printf("Found %d FQDN's in the database", len(fqdns))

	fqdnMap := make(map[string][]string, len(fqdns))
	for _, fqdn := range fqdns {
		fqdnMap[fqdn.FQDN] = []string{}
		for _, IPs := range fqdn.IPs {
			fqdnMap[fqdn.FQDN] = append(fqdnMap[fqdn.FQDN], IPs.Address)
		}
	}

	return fqdnMap, nil
}

func saveFQDNAndIPsToDatabase(db *gorm.DB, fqdnMap map[string][]string) {
	rowsAffectedCounter := 0
	for fqdn, ips := range fqdnMap {
		var fqdnRecord models.FQDN
		if err := db.Where("fqdn = ?", fqdn).First(&fqdnRecord).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("FQDN %s not found in the database", fqdn)
			}
		}

		tx := db.Begin()

		for _, ip := range ips {
			ipRecord := models.IP{
				FQDNID:  fqdnRecord.ID,
				Address: ip,
			}
			count := tx.Create(&ipRecord).RowsAffected
			rowsAffectedCounter += int(count)
		}

		if tx.Error != nil {
			tx.Rollback()
			log.Println(tx.Error.Error())
		}

		tx.Commit()
	}

	if rowsAffectedCounter == 0 {
		log.Printf("%d rows were affected, no new IP's found", rowsAffectedCounter)
	} else {
		log.Printf("%d rows were affected", rowsAffectedCounter)
	}
}
