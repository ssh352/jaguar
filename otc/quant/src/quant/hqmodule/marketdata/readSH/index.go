package readSH

import (
	"quant/hqmodule/hqbase"
)

func getAmountIdx(b []byte) (amount1 int64) {
	amount1 = 0
	for i := 0; i < 15; i++ {
		if i == 14 || i == 13 {
			continue
		}
		if b[33+i] == 49 {
			amount1 += powb(10, 14-i)
		} else if b[33+i] == 50 {
			amount1 += 2 * powb(10, 14-i)
		} else if b[33+i] == 51 {
			amount1 += 3 * powb(10, 14-i)
		} else if b[33+i] == 52 {
			amount1 += 4 * powb(10, 14-i)
		} else if b[33+i] == 53 {
			amount1 += 5 * powb(10, 14-i)
		} else if b[33+i] == 54 {
			amount1 += 6 * powb(10, 14-i)
		} else if b[33+i] == 55 {
			amount1 += 7 * powb(10, 14-i)
		} else if b[33+i] == 56 {
			amount1 += 8 * powb(10, 14-i)
		} else if b[33+i] == 57 {
			amount1 += 9 * powb(10, 14-i)
		} else {
		}
	}

	for i := 0; i < 2; i++ {
		if b[47+i] == 49 {
			amount1 += powb(10, 1-i)
		} else if b[47+i] == 50 {
			amount1 += 2 * powb(10, 1-i)
		} else if b[47+i] == 51 {
			amount1 += 3 * powb(10, 1-i)
		} else if b[47+i] == 52 {
			amount1 += 4 * powb(10, 1-i)
		} else if b[47+i] == 53 {
			amount1 += 5 * powb(10, 1-i)
		} else if b[47+i] == 54 {
			amount1 += 6 * powb(10, 1-i)
		} else if b[47+i] == 55 {
			amount1 += 7 * powb(10, 1-i)
		} else if b[47+i] == 56 {
			amount1 += 8 * powb(10, 1-i)
		} else if b[47+i] == 57 {
			amount1 += 9 * powb(10, 1-i)
		} else {
		}
	}
	return
}

func getLastpriceIdx(b []byte) (lastprice1 int64) {
	lastprice1 = 0
	for i := 0; i < 10; i++ {
		if i == 6 || i == 7 || i == 8 || i == 9 {
			continue
		}
		if b[50+i] == 49 {
			lastprice1 += powb(10, 9-i)
		} else if b[50+i] == 50 {
			lastprice1 += 2 * powb(10, 9-i)
		} else if b[50+i] == 51 {
			lastprice1 += 3 * powb(10, 9-i)
		} else if b[50+i] == 52 {
			lastprice1 += 4 * powb(10, 9-i)
		} else if b[50+i] == 53 {
			lastprice1 += 5 * powb(10, 9-i)
		} else if b[50+i] == 54 {
			lastprice1 += 6 * powb(10, 9-i)
		} else if b[50+i] == 55 {
			lastprice1 += 7 * powb(10, 9-i)
		} else if b[50+i] == 56 {
			lastprice1 += 8 * powb(10, 9-i)
		} else if b[50+i] == 57 {
			lastprice1 += 9 * powb(10, 9-i)
		} else {
		}
	}
	for i := 0; i < 4; i++ {
		if b[57+i] == 49 {
			lastprice1 += powb(10, 3-i)
		} else if b[57+i] == 50 {
			lastprice1 += 2 * powb(10, 3-i)
		} else if b[57+i] == 51 {
			lastprice1 += 3 * powb(10, 3-i)
		} else if b[57+i] == 52 {
			lastprice1 += 4 * powb(10, 3-i)
		} else if b[57+i] == 53 {
			lastprice1 += 5 * powb(10, 3-i)
		} else if b[57+i] == 54 {
			lastprice1 += 6 * powb(10, 3-i)
		} else if b[57+i] == 55 {
			lastprice1 += 7 * powb(10, 3-i)
		} else if b[57+i] == 56 {
			lastprice1 += 8 * powb(10, 3-i)
		} else if b[57+i] == 57 {
			lastprice1 += 9 * powb(10, 3-i)
		} else {
		}
	}
	return
}

func getOpenIdx(b []byte) (open1 int64) {
	open1 = 0
	for i := 0; i < 10; i++ {
		if i == 6 || i == 7 || i == 8 || i == 9 {
			continue
		}
		if b[62+i] == 49 {
			open1 += powb(10, 9-i)
		} else if b[62+i] == 50 {
			open1 += 2 * powb(10, 9-i)
		} else if b[62+i] == 51 {
			open1 += 3 * powb(10, 9-i)
		} else if b[62+i] == 52 {
			open1 += 4 * powb(10, 9-i)
		} else if b[62+i] == 53 {
			open1 += 5 * powb(10, 9-i)
		} else if b[62+i] == 54 {
			open1 += 6 * powb(10, 9-i)
		} else if b[62+i] == 55 {
			open1 += 7 * powb(10, 9-i)
		} else if b[62+i] == 56 {
			open1 += 8 * powb(10, 9-i)
		} else if b[62+i] == 57 {
			open1 += 9 * powb(10, 9-i)
		} else {
		}
	}
	for i := 0; i < 4; i++ {
		if b[69+i] == 49 {
			open1 += powb(10, 3-i)
		} else if b[69+i] == 50 {
			open1 += 2 * powb(10, 3-i)
		} else if b[69+i] == 51 {
			open1 += 3 * powb(10, 3-i)
		} else if b[69+i] == 52 {
			open1 += 4 * powb(10, 3-i)
		} else if b[69+i] == 53 {
			open1 += 5 * powb(10, 3-i)
		} else if b[69+i] == 54 {
			open1 += 6 * powb(10, 3-i)
		} else if b[69+i] == 55 {
			open1 += 7 * powb(10, 3-i)
		} else if b[69+i] == 56 {
			open1 += 8 * powb(10, 3-i)
		} else if b[69+i] == 57 {
			open1 += 9 * powb(10, 3-i)
		} else {
		}
	}
	return
}

func getHighIdx(b []byte) (high1 int64) {
	high1 = 0
	for i := 0; i < 10; i++ {
		if i == 6 || i == 7 || i == 8 || i == 9 {
			continue
		}
		if b[74+i] == 49 {
			high1 += powb(10, 9-i)
		} else if b[74+i] == 50 {
			high1 += 2 * powb(10, 9-i)
		} else if b[74+i] == 51 {
			high1 += 3 * powb(10, 9-i)
		} else if b[74+i] == 52 {
			high1 += 4 * powb(10, 9-i)
		} else if b[74+i] == 53 {
			high1 += 5 * powb(10, 9-i)
		} else if b[74+i] == 54 {
			high1 += 6 * powb(10, 9-i)
		} else if b[74+i] == 55 {
			high1 += 7 * powb(10, 9-i)
		} else if b[74+i] == 56 {
			high1 += 8 * powb(10, 9-i)
		} else if b[74+i] == 57 {
			high1 += 9 * powb(10, 9-i)
		} else {
		}
	}

	for i := 0; i < 4; i++ {
		if b[81+i] == 49 {
			high1 += powb(10, 3-i)
		} else if b[81+i] == 50 {
			high1 += 2 * powb(10, 3-i)
		} else if b[81+i] == 51 {
			high1 += 3 * powb(10, 3-i)
		} else if b[81+i] == 52 {
			high1 += 4 * powb(10, 3-i)
		} else if b[81+i] == 53 {
			high1 += 5 * powb(10, 3-i)
		} else if b[81+i] == 54 {
			high1 += 6 * powb(10, 3-i)
		} else if b[81+i] == 55 {
			high1 += 7 * powb(10, 3-i)
		} else if b[81+i] == 56 {
			high1 += 8 * powb(10, 3-i)
		} else if b[81+i] == 57 {
			high1 += 9 * powb(10, 3-i)
		} else {
		}
	}
	return
}

func getLowIdx(b []byte) (low1 int64) {
	low1 = 0
	for i := 0; i < 10; i++ {
		if i == 6 || i == 7 || i == 8 || i == 9 {
			continue
		}
		if b[86+i] == 49 {
			low1 += powb(10, 9-i)
		} else if b[86+i] == 50 {
			low1 += 2 * powb(10, 9-i)
		} else if b[86+i] == 51 {
			low1 += 3 * powb(10, 9-i)
		} else if b[86+i] == 52 {
			low1 += 4 * powb(10, 9-i)
		} else if b[86+i] == 53 {
			low1 += 5 * powb(10, 9-i)
		} else if b[86+i] == 54 {
			low1 += 6 * powb(10, 9-i)
		} else if b[86+i] == 55 {
			low1 += 7 * powb(10, 9-i)
		} else if b[86+i] == 56 {
			low1 += 8 * powb(10, 9-i)
		} else if b[86+i] == 57 {
			low1 += 9 * powb(10, 9-i)
		} else {
		}
	}
	for i := 0; i < 4; i++ {
		if b[93+i] == 49 {
			low1 += powb(10, 3-i)
		} else if b[93+i] == 50 {
			low1 += 2 * powb(10, 3-i)
		} else if b[93+i] == 51 {
			low1 += 3 * powb(10, 3-i)
		} else if b[93+i] == 52 {
			low1 += 4 * powb(10, 3-i)
		} else if b[93+i] == 53 {
			low1 += 5 * powb(10, 3-i)
		} else if b[93+i] == 54 {
			low1 += 6 * powb(10, 3-i)
		} else if b[93+i] == 55 {
			low1 += 7 * powb(10, 3-i)
		} else if b[93+i] == 56 {
			low1 += 8 * powb(10, 3-i)
		} else if b[93+i] == 57 {
			low1 += 9 * powb(10, 3-i)
		} else {
		}
	}

	return
}

func getTradePriceIdx(b []byte) (tradeprice1 int64) {
	tradeprice1 = 0
	for i := 0; i < 10; i++ {
		if i == 6 || i == 7 || i == 8 || i == 9 {
			continue
		}
		if b[98+i] == 49 {
			tradeprice1 += powb(10, 9-i)
		} else if b[98+i] == 50 {
			tradeprice1 += 2 * powb(10, 9-i)
		} else if b[98+i] == 51 {
			tradeprice1 += 3 * powb(10, 9-i)
		} else if b[98+i] == 52 {
			tradeprice1 += 4 * powb(10, 9-i)
		} else if b[98+i] == 53 {
			tradeprice1 += 5 * powb(10, 9-i)
		} else if b[98+i] == 54 {
			tradeprice1 += 6 * powb(10, 9-i)
		} else if b[98+i] == 55 {
			tradeprice1 += 7 * powb(10, 9-i)
		} else if b[98+i] == 56 {
			tradeprice1 += 8 * powb(10, 9-i)
		} else if b[98+i] == 57 {
			tradeprice1 += 9 * powb(10, 9-i)
		} else {
		}
	}
	for i := 0; i < 4; i++ {
		if b[105+i] == 49 {
			tradeprice1 += powb(10, 3-i)
		} else if b[105+i] == 50 {
			tradeprice1 += 2 * powb(10, 3-i)
		} else if b[105+i] == 51 {
			tradeprice1 += 3 * powb(10, 3-i)
		} else if b[105+i] == 52 {
			tradeprice1 += 4 * powb(10, 3-i)
		} else if b[105+i] == 53 {
			tradeprice1 += 5 * powb(10, 3-i)
		} else if b[105+i] == 54 {
			tradeprice1 += 6 * powb(10, 3-i)
		} else if b[105+i] == 55 {
			tradeprice1 += 7 * powb(10, 3-i)
		} else if b[105+i] == 56 {
			tradeprice1 += 8 * powb(10, 3-i)
		} else if b[105+i] == 57 {
			tradeprice1 += 9 * powb(10, 3-i)
		} else {
		}
	}
	return
}

func getCloseIdx(b []byte) (closepx1 int64) {
	closepx1 = 0
	for i := 0; i < 10; i++ {
		if i == 6 || i == 7 || i == 8 || i == 9 {
			continue
		}
		if b[110+i] == 49 {
			closepx1 += powb(10, 9-i)
		} else if b[110+i] == 50 {
			closepx1 += 2 * powb(10, 9-i)
		} else if b[110+i] == 51 {
			closepx1 += 3 * powb(10, 9-i)
		} else if b[110+i] == 52 {
			closepx1 += 4 * powb(10, 9-i)
		} else if b[110+i] == 53 {
			closepx1 += 5 * powb(10, 9-i)
		} else if b[110+i] == 54 {
			closepx1 += 6 * powb(10, 9-i)
		} else if b[110+i] == 55 {
			closepx1 += 7 * powb(10, 9-i)
		} else if b[110+i] == 56 {
			closepx1 += 8 * powb(10, 9-i)
		} else if b[110+i] == 57 {
			closepx1 += 9 * powb(10, 9-i)
		} else {
		}
	}

	for i := 0; i < 4; i++ {
		if b[117+i] == 49 {
			closepx1 += powb(10, 3-i)
		} else if b[117+i] == 50 {
			closepx1 += 2 * powb(10, 3-i)
		} else if b[117+i] == 51 {
			closepx1 += 3 * powb(10, 3-i)
		} else if b[117+i] == 52 {
			closepx1 += 4 * powb(10, 3-i)
		} else if b[117+i] == 53 {
			closepx1 += 5 * powb(10, 3-i)
		} else if b[117+i] == 54 {
			closepx1 += 6 * powb(10, 3-i)
		} else if b[117+i] == 55 {
			closepx1 += 7 * powb(10, 3-i)
		} else if b[117+i] == 56 {
			closepx1 += 8 * powb(10, 3-i)
		} else if b[117+i] == 57 {
			closepx1 += 9 * powb(10, 3-i)
		} else {
		}
	}
	return
}

func getIndexQuote(msg interface{}) hqbase.Marketdata {
	b := msg.([]byte)
	return hqbase.Marketdata{Code: string(b[0:6]) + ".SH",
		Time:     getTimeStamp(b, "INDEX"),
		PreClose: float64(getLastpriceIdx(b)) / 10000.0,

		Open:  float64(getOpenIdx(b)) / 10000.0,
		High:  float64(getHighIdx(b)) / 10000.0,
		Low:   float64(getLowIdx(b)) / 10000.0,
		Close: float64(getCloseIdx(b)) / 10000.0,

		Match:    float64(getTradePriceIdx(b)) / 10000.0,
		Volume:   getVol(b),
		Turnover: getAmountIdx(b) / 100}
}
