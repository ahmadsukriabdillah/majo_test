package services

import (
	"majo_test/repository"
	"majo_test/utils"
)

type ReportService interface {
	GetReportMonthly(merchant_id uint, outlet_id uint, month string, year string, pagination utils.Pagination) ([]repository.TransactionReport, error)
}

type reportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) *reportService {
	return &reportService{repo: repo}
}

func (r *reportService) GetReportMonthly(merchant_id uint, outlet_id uint, month string, year string, pagination utils.Pagination) ([]repository.TransactionReport, error) {
	data, err := r.repo.GetMonthlyReport(merchant_id, outlet_id, month, year, pagination)
	if err != nil {
		return nil, err
	}
	return data, nil
}
