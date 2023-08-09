package dnsapi

import (
	"database/sql"
	"fmt"
	"log"
)

func (h handler) saveFQDNAndIPsToDatabase(fqdn string, ips []string) {
	insertFQDNStmt, err := h.DB.Prepare("INSERT INTO fqdns(fqdn) VALUES ($1) ON CONFLICT (fqdn) DO NOTHING RETURNING id;")
	if err != nil {
		log.Println(err)
	}
	defer insertFQDNStmt.Close()

	query := `
		WITH ins AS (
			INSERT INTO fqdns(fqdn) VALUES ($1) ON CONFLICT (fqdn) DO NOTHING RETURNING id
		)
		SELECT id FROM ins
		UNION ALL
		SELECT id FROM fqdns WHERE fqdn = $1;
	`

	var id int
	err = h.DB.QueryRow(query, fqdn).Scan(&id)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Inserted or existing row ID:", id)

	insertIPsstmt, err := h.DB.Prepare("INSERT INTO ips(fqdn_id, address) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING ID;")
	if err != nil {
		log.Fatalln(err)
	}
	defer insertIPsstmt.Close()

	for _, v := range ips {
		var newId int
		err = insertIPsstmt.QueryRow(id, v).Scan(&newId)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Println(err)
			}

			log.Printf("Row already exists with this ip: %s", v)
		}
		if newId != 0 {
			fmt.Printf("Inserted row ID %d for IP: %s\n", newId, v)
		}
	}
}
