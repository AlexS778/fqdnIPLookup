package dnsapi

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/AlexS778/fqdnIPLookup/internal/dnsutil"
)

func ContinuousUpdate(ctx context.Context, db *sql.DB, waitTime time.Duration) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Printf("ContinuousUpdate err: %s\n", err)
			}
			log.Println("ContinuousUpdate: finished")
			return
		default:
			fqdns := getAllFQDNsFromDB(db)

			ipAdresses := make(map[string][]string)
			for _, fqdn := range fqdns {
				res, err := dnsutil.GetIPsByFQDN(fqdn)
				if err != nil {
					ipAdresses[fqdn] = []string{err.Error()}
				} else {
					ipAdresses[fqdn] = res
					saveFQDNAndIPsToDatabase(db, fqdn, ipAdresses[fqdn])
				}
			}

			time.Sleep(waitTime)
		}
	}
}

func getAllFQDNsFromDB(db *sql.DB) []string {
	var fqdn string
	var result []string
	rows, err := db.Query("SELECT fqdn FROM fqdns")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&fqdn)
		if err != nil {
			log.Println(err)
		}
		result = append(result, fqdn)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}

	return result
}

func saveFQDNAndIPsToDatabase(db *sql.DB, fqdn string, ips []string) {
	insertFQDNStmt, err := db.Prepare("INSERT INTO fqdns(fqdn) VALUES ($1) ON CONFLICT (fqdn) DO NOTHING RETURNING id;")
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
	err = db.QueryRow(query, fqdn).Scan(&id)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Inserted or existing row ID:", id)

	insertIPsstmt, err := db.Prepare("INSERT INTO ips(fqdn_id, address) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING ID;")
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
			fmt.Printf("Inserted row ID %d for IP: %s\n", newId, fqdn)
		}
	}
}
