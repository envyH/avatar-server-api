package service

import (
	"avatar/global"
)

func GetLaiBuonHelp() ([]int, error) {
	return global.LaiBuonHelp[:], nil
}

func GetDailyReport() ([]int, error) {
	return global.DailyReport[:], nil
}

func GetQuickBuyProduct() ([]int, error) {
	return global.QuickBuyProduct[:], nil
}

func GetFarmData() (map[string]any, error) {
	laiBuonHelp, err := GetLaiBuonHelp()
	if err != nil {
		return nil, err
	}
	dailyReport, err := GetDailyReport()
	if err != nil {
		return nil, err
	}
	quickBuyProduct, err := GetQuickBuyProduct()
	if err != nil {
		return nil, err
	}

	response := map[string]any{
		"boss-farm-id":      global.BOSS_FARM_ID,
		"lai-buon-help":     laiBuonHelp,
		"quick-buy-product": quickBuyProduct,
		"daily-report":      dailyReport,
	}

	return response, nil
}
