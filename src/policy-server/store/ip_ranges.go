package store

import (
	"policy-server/db"
)

//go:generate counterfeiter -o fakes/ip_ranges_repo.go --fake-name IPRangesRepo . IPRangesRepo
type IPRangesRepo interface {
	Create(db.Transaction, int, string, string) error
}

type IPRangesTable struct {
}

func (ip *IPRangesTable) Create(tx db.Transaction, groupId int, startIP, endIP string) error {
	dualStatement := ""
	if tx.DriverName() == "mysql" {
		dualStatement = " FROM DUAL "
	}

	_, err := tx.Exec(tx.Rebind(`
		INSERT INTO ip_ranges (group_id, start_ip, end_ip)
		SELECT ?, ?, ? `+dualStatement+`
		WHERE
		NOT EXISTS (
			SELECT *
			FROM policies
			WHERE group_id = ? AND start_ip = ? AND end_ip = ?
		)`),
		groupId,
		startIP,
		endIP,
		groupId,
		startIP,
		endIP,
	)
	return err
}