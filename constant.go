package main

const FILTER_SIZE_SM = 2

const FILTER_SIZE_SM_FLOAT float64 = float64(FILTER_SIZE_SM)

const FILTER_SIZE_MD = 4

const FILTER_SIZE_MD_FLOAT float64 = float64(FILTER_SIZE_MD)

const FILTER_SIZE_LG = 10

const FILTER_SIZE_LG_FLOAT float64 = float64(FILTER_SIZE_LG)

var djiFileName = "_DJI_20150319-20170203"

var djiNames = []string{"AAPL", "AXP", "BA", "CAT", "CSCO",
	"CVX", "KO", "DIS", "DD", "GE",
	"GS", "HD", "IBM", "INTC", "JNJ",
	"JPM", "MCD", "MMM", "MRK", "MSFT",
	"NKE", "PFE", "PG", "TRV", "UNH",
	"UTX", "V", "VZ", "WMT", "XOM",
}

var headers = []string{"Date", "Open", "Close", "Diff0", "DiffHighLow", "DiffSM", "DiffMD", "DiffLG", "Avg0", "AvgSM", "AvgMD", "AvgLG"}
