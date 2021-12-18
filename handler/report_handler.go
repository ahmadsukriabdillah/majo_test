package handler

import (
	"majo_test/services"
	"majo_test/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	userService   services.UserService
	reportService services.ReportService
}

func NewReportHandler(userService services.UserService, reportService services.ReportService) *ReportHandler {
	return &ReportHandler{userService, reportService}
}

type ReportMontlyRequest struct {
	MerchantId uint   `json:"merchant_id" binding:"required"`
	OutletId   uint   `json:"outlet_id" binding:"required"`
	Month      string `json:"month" binding:"required"`
	Year       string `json:"year" binding:"required"`
}

func (s *ReportHandler) MonthlyReport(ctx *gin.Context) {
	var req ReportMontlyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.NewResponseError(ctx, err)
		return
	}

	var pagination utils.Pagination
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		utils.NewResponseError(ctx, err)
		return
	}

	payload := ctx.MustGet(utils.AuthorizationPayloadKey).(*utils.Payload)
	err := s.userService.IsAllowAccess(payload.ID, strconv.FormatUint(uint64(req.MerchantId), 10), strconv.FormatUint(uint64(req.OutletId), 10))
	if err != nil {
		utils.NewResponseError(ctx, err)
		return
	}
	datas, err := s.reportService.GetReportMonthly(req.MerchantId, req.OutletId, req.Month, req.Year, pagination)
	if err != nil {
		utils.NewResponseError(ctx, err)
		return
	}
	utils.NewResponseOk(ctx, datas)
}
