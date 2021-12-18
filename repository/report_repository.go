package repository

import (
	"database/sql"
	"log"
	"majo_test/utils"
)

type TransactionReport struct {
	Tanggal      string
	MerchantName string
	OutletName   string
	Total        uint64
}

type ReportRepository interface {
	GetMonthlyReport(merchant_id uint, outlet_id uint, month string, year string, pagination utils.Pagination) ([]TransactionReport, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *reportRepository {
	return &reportRepository{db: db}
}

func (u reportRepository) GetMonthlyReport(merchant_id uint, outlet_id uint, month string, year string, pagination utils.Pagination) ([]TransactionReport, error) {
	offset := (pagination.Page - 1) * pagination.Limit

	query := `select a.tanggal, a.merchant_name, a.outlet_name, COALESCE(sum(t.bill_total),0) as total
		from 
		(
		select '2021-01-01' + INTERVAL (a.a + (10 * b.a) + (100 * c.a)) DAY as tanggal,
			(select merchant_name from merchants where id = ?) as merchant_name, 	
			(select outlet_name from outlets where id = ?) as outlet_name
		from (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as a
		cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as b
		cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as c
		) a 
		left join transactions t on (a.tanggal = date(t.created_at) and t.outlet_id = ? and t.merchant_id = ?)
		where MONTH(a.tanggal) = ? and YEAR(a.tanggal) = ?
		group by a.tanggal order by a.tanggal asc limit ? offset ?`

	datas := []TransactionReport{}
	rows, err := u.db.Query(query, merchant_id, outlet_id, merchant_id, outlet_id, month, year, pagination.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := new(TransactionReport)
		err := rows.Scan(&item.Tanggal, &item.MerchantName, &item.OutletName, &item.Total)
		if err != nil {
			log.Fatal(err)
		}
		datas = append(datas, *item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return datas, nil
}
