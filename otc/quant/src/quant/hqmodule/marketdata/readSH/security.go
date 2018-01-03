package readSH

import (
	// log "github.com/thinkboy/log4go"
	"quant/hqmodule/hqbase"
	"strconv"
	// "time"
)

func powb(x, n int) int64 {
	if n == 0 {
		return 1
	}
	for {
		if (n & 1) != 0 {
			break
		}
		n >>= 1
		x *= x
	}
	result := x
	n >>= 1
	for {
		if n == 0 {
			break
		}
		x *= x
		if (n & 1) != 0 {
			result *= x
		}
		n >>= 1
	}
	return int64(result)
}

func powbint(x, n int) int {
	if n == 0 {
		return 1
	}
	for {
		if (n & 1) != 0 {
			break
		}
		n >>= 1
		x *= x
	}
	result := x
	n >>= 1
	for {
		if n == 0 {
			break
		}
		x *= x
		if (n & 1) != 0 {
			result *= x
		}
		n >>= 1
	}
	return result
}

func getVol(b []byte) (volume int64) {
	volume = 0
	for i := 0; i < 16; i++ {
		if b[16+i] == 49 {
			volume += powb(10, 15-i)
		} else if b[16+i] == 50 {
			volume += 2 * powb(10, 15-i)
		} else if b[16+i] == 51 {
			volume += 3 * powb(10, 15-i)
		} else if b[16+i] == 52 {
			volume += 4 * powb(10, 15-i)
		} else if b[16+i] == 53 {
			volume += 5 * powb(10, 15-i)
		} else if b[16+i] == 54 {
			volume += 6 * powb(10, 15-i)
		} else if b[16+i] == 55 {
			volume += 7 * powb(10, 15-i)
		} else if b[16+i] == 56 {
			volume += 8 * powb(10, 15-i)
		} else if b[16+i] == 57 {
			volume += 9 * powb(10, 15-i)
		} else {
		}
	}
	return
}

func getAmount(b []byte) (amount int64) {
	amount = 0
	for i := 0; i < 16; i++ {
		n := 0
		if i == 13 {
			continue
		} else if i > 13 {
			n = i - 1
		} else {
			n = i
		}

		if b[33+i] == 49 {
			amount += powb(10, 14-n)
		} else if b[33+i] == 50 {
			amount += 2 * powb(10, 14-n)
		} else if b[33+i] == 51 {
			amount += 3 * powb(10, 14-n)
		} else if b[33+i] == 52 {
			amount += 4 * powb(10, 14-n)
		} else if b[33+i] == 53 {
			amount += 5 * powb(10, 14-n)
		} else if b[33+i] == 54 {
			amount += 6 * powb(10, 14-n)
		} else if b[33+i] == 55 {
			amount += 7 * powb(10, 14-n)
		} else if b[33+i] == 56 {
			amount += 8 * powb(10, 14-n)
		} else if b[33+i] == 57 {
			amount += 9 * powb(10, 14-n)
		} else {
		}
	}
	return
}

func getLastprice(b []byte) (lastprice int) {
	lastprice = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}
		if b[50+i] == 49 {
			lastprice += powbint(10, 9-n)
		} else if b[50+i] == 50 {
			lastprice += 2 * powbint(10, 9-n)
		} else if b[50+i] == 51 {
			lastprice += 3 * powbint(10, 9-n)
		} else if b[50+i] == 52 {
			lastprice += 4 * powbint(10, 9-n)
		} else if b[50+i] == 53 {
			lastprice += 5 * powbint(10, 9-n)
		} else if b[50+i] == 54 {
			lastprice += 6 * powbint(10, 9-n)
		} else if b[50+i] == 55 {
			lastprice += 7 * powbint(10, 9-n)
		} else if b[50+i] == 56 {
			lastprice += 8 * powbint(10, 9-n)
		} else if b[50+i] == 57 {
			lastprice += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getOpen(b []byte) (open int) {
	open = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[62+i] == 49 {
			open += powbint(10, 9-n)
		} else if b[62+i] == 50 {
			open += 2 * powbint(10, 9-n)
		} else if b[62+i] == 51 {
			open += 3 * powbint(10, 9-n)
		} else if b[62+i] == 52 {
			open += 4 * powbint(10, 9-n)
		} else if b[62+i] == 53 {
			open += 5 * powbint(10, 9-n)
		} else if b[62+i] == 54 {
			open += 6 * powbint(10, 9-n)
		} else if b[62+i] == 55 {
			open += 7 * powbint(10, 9-n)
		} else if b[62+i] == 56 {
			open += 8 * powbint(10, 9-n)
		} else if b[62+i] == 57 {
			open += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getHigh(b []byte) (high int) {
	high = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[74+i] == 49 {
			high += powbint(10, 9-n)
		} else if b[74+i] == 50 {
			high += 2 * powbint(10, 9-n)
		} else if b[74+i] == 51 {
			high += 3 * powbint(10, 9-n)
		} else if b[74+i] == 52 {
			high += 4 * powbint(10, 9-n)
		} else if b[74+i] == 53 {
			high += 5 * powbint(10, 9-n)
		} else if b[74+i] == 54 {
			high += 6 * powbint(10, 9-n)
		} else if b[74+i] == 55 {
			high += 7 * powbint(10, 9-n)
		} else if b[74+i] == 56 {
			high += 8 * powbint(10, 9-n)
		} else if b[74+i] == 57 {
			high += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getLow(b []byte) (low int) {
	low = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[86+i] == 49 {
			low += powbint(10, 9-n)
		} else if b[86+i] == 50 {
			low += 2 * powbint(10, 9-n)
		} else if b[86+i] == 51 {
			low += 3 * powbint(10, 9-n)
		} else if b[86+i] == 52 {
			low += 4 * powbint(10, 9-n)
		} else if b[86+i] == 53 {
			low += 5 * powbint(10, 9-n)
		} else if b[86+i] == 54 {
			low += 6 * powbint(10, 9-n)
		} else if b[86+i] == 55 {
			low += 7 * powbint(10, 9-n)
		} else if b[86+i] == 56 {
			low += 8 * powbint(10, 9-n)
		} else if b[86+i] == 57 {
			low += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getTradePrice(b []byte) (tradeprice int) {
	tradeprice = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[98+i] == 49 {
			tradeprice += powbint(10, 9-n)
		} else if b[98+i] == 50 {
			tradeprice += 2 * powbint(10, 9-n)
		} else if b[98+i] == 51 {
			tradeprice += 3 * powbint(10, 9-n)
		} else if b[98+i] == 52 {
			tradeprice += 4 * powbint(10, 9-n)
		} else if b[98+i] == 53 {
			tradeprice += 5 * powbint(10, 9-n)
		} else if b[98+i] == 54 {
			tradeprice += 6 * powbint(10, 9-n)
		} else if b[98+i] == 55 {
			tradeprice += 7 * powbint(10, 9-n)
		} else if b[98+i] == 56 {
			tradeprice += 8 * powbint(10, 9-n)
		} else if b[98+i] == 57 {
			tradeprice += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getClosepx(b []byte) (closepx int) {
	closepx = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[110+i] == 49 {
			closepx += powbint(10, 9-n)
		} else if b[110+i] == 50 {
			closepx += 2 * powbint(10, 9-n)
		} else if b[110+i] == 51 {
			closepx += 3 * powbint(10, 9-n)
		} else if b[110+i] == 52 {
			closepx += 4 * powbint(10, 9-n)
		} else if b[110+i] == 53 {
			closepx += 5 * powbint(10, 9-n)
		} else if b[110+i] == 54 {
			closepx += 6 * powbint(10, 9-n)
		} else if b[110+i] == 55 {
			closepx += 7 * powbint(10, 9-n)
		} else if b[110+i] == 56 {
			closepx += 8 * powbint(10, 9-n)
		} else if b[110+i] == 57 {
			closepx += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getbp1(b []byte) (bp1 int) {
	bp1 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[122+i] == 49 {
			bp1 += powbint(10, 9-n)
		} else if b[122+i] == 50 {
			bp1 += 2 * powbint(10, 9-n)
		} else if b[122+i] == 51 {
			bp1 += 3 * powbint(10, 9-n)
		} else if b[122+i] == 52 {
			bp1 += 4 * powbint(10, 9-n)
		} else if b[122+i] == 53 {
			bp1 += 5 * powbint(10, 9-n)
		} else if b[122+i] == 54 {
			bp1 += 6 * powbint(10, 9-n)
		} else if b[122+i] == 55 {
			bp1 += 7 * powbint(10, 9-n)
		} else if b[122+i] == 56 {
			bp1 += 8 * powbint(10, 9-n)
		} else if b[122+i] == 57 {
			bp1 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getbv1(b []byte) (bv1 int64) {
	bv1 = 0
	for i := 0; i < 12; i++ {
		if b[134+i] == 49 {
			bv1 += powb(10, 11-i)
		} else if b[134+i] == 50 {
			bv1 += 2 * powb(10, 11-i)
		} else if b[134+i] == 51 {
			bv1 += 3 * powb(10, 11-i)
		} else if b[134+i] == 52 {
			bv1 += 4 * powb(10, 11-i)
		} else if b[134+i] == 53 {
			bv1 += 5 * powb(10, 11-i)
		} else if b[134+i] == 54 {
			bv1 += 6 * powb(10, 11-i)
		} else if b[134+i] == 55 {
			bv1 += 7 * powb(10, 11-i)
		} else if b[134+i] == 56 {
			bv1 += 8 * powb(10, 11-i)
		} else if b[134+i] == 57 {
			bv1 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getsp1(b []byte) (sp1 int) {
	sp1 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[147+i] == 49 {
			sp1 += powbint(10, 9-n)
		} else if b[147+i] == 50 {
			sp1 += 2 * powbint(10, 9-n)
		} else if b[147+i] == 51 {
			sp1 += 3 * powbint(10, 9-n)
		} else if b[147+i] == 52 {
			sp1 += 4 * powbint(10, 9-n)
		} else if b[147+i] == 53 {
			sp1 += 5 * powbint(10, 9-n)
		} else if b[147+i] == 54 {
			sp1 += 6 * powbint(10, 9-n)
		} else if b[147+i] == 55 {
			sp1 += 7 * powbint(10, 9-n)
		} else if b[147+i] == 56 {
			sp1 += 8 * powbint(10, 9-n)
		} else if b[147+i] == 57 {
			sp1 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getsv1(b []byte) (sv1 int64) {
	sv1 = 0
	for i := 0; i < 12; i++ {
		if b[159+i] == 49 {
			sv1 += powb(10, 11-i)
		} else if b[159+i] == 50 {
			sv1 += 2 * powb(10, 11-i)
		} else if b[159+i] == 51 {
			sv1 += 3 * powb(10, 11-i)
		} else if b[159+i] == 52 {
			sv1 += 4 * powb(10, 11-i)
		} else if b[159+i] == 53 {
			sv1 += 5 * powb(10, 11-i)
		} else if b[159+i] == 54 {
			sv1 += 6 * powb(10, 11-i)
		} else if b[159+i] == 55 {
			sv1 += 7 * powb(10, 11-i)
		} else if b[159+i] == 56 {
			sv1 += 8 * powb(10, 11-i)
		} else if b[159+i] == 57 {
			sv1 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getbp2(b []byte) (bp2 int) {
	bp2 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[172+i] == 49 {
			bp2 += powbint(10, 9-n)
		} else if b[172+i] == 50 {
			bp2 += 2 * powbint(10, 9-n)
		} else if b[172+i] == 51 {
			bp2 += 3 * powbint(10, 9-n)
		} else if b[172+i] == 52 {
			bp2 += 4 * powbint(10, 9-n)
		} else if b[172+i] == 53 {
			bp2 += 5 * powbint(10, 9-n)
		} else if b[172+i] == 54 {
			bp2 += 6 * powbint(10, 9-n)
		} else if b[172+i] == 55 {
			bp2 += 7 * powbint(10, 9-n)
		} else if b[172+i] == 56 {
			bp2 += 8 * powbint(10, 9-n)
		} else if b[172+i] == 57 {
			bp2 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getbv2(b []byte) (bv2 int64) {
	bv2 = 0
	for i := 0; i < 12; i++ {
		if b[184+i] == 49 {
			bv2 += powb(10, 11-i)
		} else if b[184+i] == 50 {
			bv2 += 2 * powb(10, 11-i)
		} else if b[184+i] == 51 {
			bv2 += 3 * powb(10, 11-i)
		} else if b[184+i] == 52 {
			bv2 += 4 * powb(10, 11-i)
		} else if b[184+i] == 53 {
			bv2 += 5 * powb(10, 11-i)
		} else if b[184+i] == 54 {
			bv2 += 6 * powb(10, 11-i)
		} else if b[184+i] == 55 {
			bv2 += 7 * powb(10, 11-i)
		} else if b[184+i] == 56 {
			bv2 += 8 * powb(10, 11-i)
		} else if b[184+i] == 57 {
			bv2 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getsp2(b []byte) (sp2 int) {
	sp2 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[197+i] == 49 {
			sp2 += powbint(10, 9-n)
		} else if b[197+i] == 50 {
			sp2 += 2 * powbint(10, 9-n)
		} else if b[197+i] == 51 {
			sp2 += 3 * powbint(10, 9-n)
		} else if b[197+i] == 52 {
			sp2 += 4 * powbint(10, 9-n)
		} else if b[197+i] == 53 {
			sp2 += 5 * powbint(10, 9-n)
		} else if b[197+i] == 54 {
			sp2 += 6 * powbint(10, 9-n)
		} else if b[197+i] == 55 {
			sp2 += 7 * powbint(10, 9-n)
		} else if b[197+i] == 56 {
			sp2 += 8 * powbint(10, 9-n)
		} else if b[197+i] == 57 {
			sp2 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getsv2(b []byte) (sv2 int64) {
	sv2 = 0
	for i := 0; i < 12; i++ {
		if b[209+i] == 49 {
			sv2 += powb(10, 11-i)
		} else if b[209+i] == 50 {
			sv2 += 2 * powb(10, 11-i)
		} else if b[209+i] == 51 {
			sv2 += 3 * powb(10, 11-i)
		} else if b[209+i] == 52 {
			sv2 += 4 * powb(10, 11-i)
		} else if b[209+i] == 53 {
			sv2 += 5 * powb(10, 11-i)
		} else if b[209+i] == 54 {
			sv2 += 6 * powb(10, 11-i)
		} else if b[209+i] == 55 {
			sv2 += 7 * powb(10, 11-i)
		} else if b[209+i] == 56 {
			sv2 += 8 * powb(10, 11-i)
		} else if b[209+i] == 57 {
			sv2 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getbp3(b []byte) (bp3 int) {
	bp3 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[222+i] == 49 {
			bp3 += powbint(10, 9-n)
		} else if b[222+i] == 50 {
			bp3 += 2 * powbint(10, 9-n)
		} else if b[222+i] == 51 {
			bp3 += 3 * powbint(10, 9-n)
		} else if b[222+i] == 52 {
			bp3 += 4 * powbint(10, 9-n)
		} else if b[222+i] == 53 {
			bp3 += 5 * powbint(10, 9-n)
		} else if b[222+i] == 54 {
			bp3 += 6 * powbint(10, 9-n)
		} else if b[222+i] == 55 {
			bp3 += 7 * powbint(10, 9-n)
		} else if b[222+i] == 56 {
			bp3 += 8 * powbint(10, 9-n)
		} else if b[222+i] == 57 {
			bp3 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getbv3(b []byte) (bv3 int64) {
	bv3 = 0
	for i := 0; i < 12; i++ {
		if b[234+i] == 49 {
			bv3 += powb(10, 11-i)
		} else if b[234+i] == 50 {
			bv3 += 2 * powb(10, 11-i)
		} else if b[234+i] == 51 {
			bv3 += 3 * powb(10, 11-i)
		} else if b[234+i] == 52 {
			bv3 += 4 * powb(10, 11-i)
		} else if b[234+i] == 53 {
			bv3 += 5 * powb(10, 11-i)
		} else if b[234+i] == 54 {
			bv3 += 6 * powb(10, 11-i)
		} else if b[234+i] == 55 {
			bv3 += 7 * powb(10, 11-i)
		} else if b[234+i] == 56 {
			bv3 += 8 * powb(10, 11-i)
		} else if b[234+i] == 57 {
			bv3 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getsp3(b []byte) (sp3 int) {
	sp3 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[247+i] == 49 {
			sp3 += powbint(10, 9-n)
		} else if b[247+i] == 50 {
			sp3 += 2 * powbint(10, 9-n)
		} else if b[247+i] == 51 {
			sp3 += 3 * powbint(10, 9-n)
		} else if b[247+i] == 52 {
			sp3 += 4 * powbint(10, 9-n)
		} else if b[247+i] == 53 {
			sp3 += 5 * powbint(10, 9-n)
		} else if b[247+i] == 54 {
			sp3 += 6 * powbint(10, 9-n)
		} else if b[247+i] == 55 {
			sp3 += 7 * powbint(10, 9-n)
		} else if b[247+i] == 56 {
			sp3 += 8 * powbint(10, 9-n)
		} else if b[247+i] == 57 {
			sp3 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getsv3(b []byte) (sv3 int64) {
	sv3 = 0
	for i := 0; i < 12; i++ {
		if b[259+i] == 49 {
			sv3 += powb(10, 11-i)
		} else if b[259+i] == 50 {
			sv3 += 2 * powb(10, 11-i)
		} else if b[259+i] == 51 {
			sv3 += 3 * powb(10, 11-i)
		} else if b[259+i] == 52 {
			sv3 += 4 * powb(10, 11-i)
		} else if b[259+i] == 53 {
			sv3 += 5 * powb(10, 11-i)
		} else if b[259+i] == 54 {
			sv3 += 6 * powb(10, 11-i)
		} else if b[259+i] == 55 {
			sv3 += 7 * powb(10, 11-i)
		} else if b[259+i] == 56 {
			sv3 += 8 * powb(10, 11-i)
		} else if b[259+i] == 57 {
			sv3 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getbp4(b []byte) (bp4 int) {
	bp4 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[272+i] == 49 {
			bp4 += powbint(10, 9-n)
		} else if b[272+i] == 50 {
			bp4 += 2 * powbint(10, 9-n)
		} else if b[272+i] == 51 {
			bp4 += 3 * powbint(10, 9-n)
		} else if b[272+i] == 52 {
			bp4 += 4 * powbint(10, 9-n)
		} else if b[272+i] == 53 {
			bp4 += 5 * powbint(10, 9-n)
		} else if b[272+i] == 54 {
			bp4 += 6 * powbint(10, 9-n)
		} else if b[272+i] == 55 {
			bp4 += 7 * powbint(10, 9-n)
		} else if b[272+i] == 56 {
			bp4 += 8 * powbint(10, 9-n)
		} else if b[272+i] == 57 {
			bp4 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getbv4(b []byte) (bv4 int64) {
	bv4 = 0
	for i := 0; i < 12; i++ {
		if b[284+i] == 49 {
			bv4 += powb(10, 11-i)
		} else if b[284+i] == 50 {
			bv4 += 2 * powb(10, 11-i)
		} else if b[284+i] == 51 {
			bv4 += 3 * powb(10, 11-i)
		} else if b[284+i] == 52 {
			bv4 += 4 * powb(10, 11-i)
		} else if b[284+i] == 53 {
			bv4 += 5 * powb(10, 11-i)
		} else if b[284+i] == 54 {
			bv4 += 6 * powb(10, 11-i)
		} else if b[284+i] == 55 {
			bv4 += 7 * powb(10, 11-i)
		} else if b[284+i] == 56 {
			bv4 += 8 * powb(10, 11-i)
		} else if b[284+i] == 57 {
			bv4 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getsp4(b []byte) (sp4 int) {
	sp4 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[297+i] == 49 {
			sp4 += powbint(10, 9-n)
		} else if b[297+i] == 50 {
			sp4 += 2 * powbint(10, 9-n)
		} else if b[297+i] == 51 {
			sp4 += 3 * powbint(10, 9-n)
		} else if b[297+i] == 52 {
			sp4 += 4 * powbint(10, 9-n)
		} else if b[297+i] == 53 {
			sp4 += 5 * powbint(10, 9-n)
		} else if b[297+i] == 54 {
			sp4 += 6 * powbint(10, 9-n)
		} else if b[297+i] == 55 {
			sp4 += 7 * powbint(10, 9-n)
		} else if b[297+i] == 56 {
			sp4 += 8 * powbint(10, 9-n)
		} else if b[297+i] == 57 {
			sp4 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getsv4(b []byte) (sv4 int64) {
	sv4 = 0
	for i := 0; i < 12; i++ {
		if b[309+i] == 49 {
			sv4 += powb(10, 11-i)
		} else if b[309+i] == 50 {
			sv4 += 2 * powb(10, 11-i)
		} else if b[309+i] == 51 {
			sv4 += 3 * powb(10, 11-i)
		} else if b[309+i] == 52 {
			sv4 += 4 * powb(10, 11-i)
		} else if b[309+i] == 53 {
			sv4 += 5 * powb(10, 11-i)
		} else if b[309+i] == 54 {
			sv4 += 6 * powb(10, 11-i)
		} else if b[309+i] == 55 {
			sv4 += 7 * powb(10, 11-i)
		} else if b[309+i] == 56 {
			sv4 += 8 * powb(10, 11-i)
		} else if b[309+i] == 57 {
			sv4 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getbp5(b []byte) (bp5 int) {
	bp5 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[322+i] == 49 {
			bp5 += powbint(10, 9-n)
		} else if b[322+i] == 50 {
			bp5 += 2 * powbint(10, 9-n)
		} else if b[322+i] == 51 {
			bp5 += 3 * powbint(10, 9-n)
		} else if b[322+i] == 52 {
			bp5 += 4 * powbint(10, 9-n)
		} else if b[322+i] == 53 {
			bp5 += 5 * powbint(10, 9-n)
		} else if b[322+i] == 54 {
			bp5 += 6 * powbint(10, 9-n)
		} else if b[322+i] == 55 {
			bp5 += 7 * powbint(10, 9-n)
		} else if b[322+i] == 56 {
			bp5 += 8 * powbint(10, 9-n)
		} else if b[322+i] == 57 {
			bp5 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getbv5(b []byte) (bv5 int64) {
	bv5 = 0
	for i := 0; i < 12; i++ {
		if b[334+i] == 49 {
			bv5 += powb(10, 11-i)
		} else if b[334+i] == 50 {
			bv5 += 2 * powb(10, 11-i)
		} else if b[334+i] == 51 {
			bv5 += 3 * powb(10, 11-i)
		} else if b[334+i] == 52 {
			bv5 += 4 * powb(10, 11-i)
		} else if b[334+i] == 53 {
			bv5 += 5 * powb(10, 11-i)
		} else if b[334+i] == 54 {
			bv5 += 6 * powb(10, 11-i)
		} else if b[334+i] == 55 {
			bv5 += 7 * powb(10, 11-i)
		} else if b[334+i] == 56 {
			bv5 += 8 * powb(10, 11-i)
		} else if b[334+i] == 57 {
			bv5 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getsp5(b []byte) (sp5 int) {
	sp5 = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[347+i] == 49 {
			sp5 += powbint(10, 9-n)
		} else if b[347+i] == 50 {
			sp5 += 2 * powbint(10, 9-n)
		} else if b[347+i] == 51 {
			sp5 += 3 * powbint(10, 9-n)
		} else if b[347+i] == 52 {
			sp5 += 4 * powbint(10, 9-n)
		} else if b[347+i] == 53 {
			sp5 += 5 * powbint(10, 9-n)
		} else if b[347+i] == 54 {
			sp5 += 6 * powbint(10, 9-n)
		} else if b[347+i] == 55 {
			sp5 += 7 * powbint(10, 9-n)
		} else if b[347+i] == 56 {
			sp5 += 8 * powbint(10, 9-n)
		} else if b[347+i] == 57 {
			sp5 += 9 * powbint(10, 9-n)
		} else {
		}
	}
	return
}

func getsv5(b []byte) (sv5 int64) {
	sv5 = 0
	for i := 0; i < 12; i++ {
		if b[359+i] == 49 {
			sv5 += powb(10, 11-i)
		} else if b[359+i] == 50 {
			sv5 += 2 * powb(10, 11-i)
		} else if b[359+i] == 51 {
			sv5 += 3 * powb(10, 11-i)
		} else if b[359+i] == 52 {
			sv5 += 4 * powb(10, 11-i)
		} else if b[359+i] == 53 {
			sv5 += 5 * powb(10, 11-i)
		} else if b[359+i] == 54 {
			sv5 += 6 * powb(10, 11-i)
		} else if b[359+i] == 55 {
			sv5 += 7 * powb(10, 11-i)
		} else if b[359+i] == 56 {
			sv5 += 8 * powb(10, 11-i)
		} else if b[359+i] == 57 {
			sv5 += 9 * powb(10, 11-i)
		} else {
		}
	}
	return
}

func getPrecloseiopv(b []byte) (precloseiopv int64) {
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[372+i] == 49 {
			precloseiopv += powb(10, 9-n)
		} else if b[372+i] == 50 {
			precloseiopv += 2 * powb(10, 9-n)
		} else if b[372+i] == 51 {
			precloseiopv += 3 * powb(10, 9-n)
		} else if b[372+i] == 52 {
			precloseiopv += 4 * powb(10, 9-n)
		} else if b[372+i] == 53 {
			precloseiopv += 5 * powb(10, 9-n)
		} else if b[372+i] == 54 {
			precloseiopv += 6 * powb(10, 9-n)
		} else if b[372+i] == 55 {
			precloseiopv += 7 * powb(10, 9-n)
		} else if b[372+i] == 56 {
			precloseiopv += 8 * powb(10, 9-n)
		} else if b[372+i] == 57 {
			precloseiopv += 9 * powb(10, 9-n)
		} else {
		}
	}
	return
}

func getIOPV(b []byte) (iopv int64) {
	iopv = 0
	for i := 0; i < 11; i++ {
		n := 0
		if i == 7 {
			continue
		} else if i > 7 {
			n = i - 1
		} else {
			n = i
		}

		if b[384+i] == 49 {
			iopv += powb(10, 9-n)
		} else if b[384+i] == 50 {
			iopv += 2 * powb(10, 9-n)
		} else if b[384+i] == 51 {
			iopv += 3 * powb(10, 9-n)
		} else if b[384+i] == 52 {
			iopv += 4 * powb(10, 9-n)
		} else if b[384+i] == 53 {
			iopv += 5 * powb(10, 9-n)
		} else if b[384+i] == 54 {
			iopv += 6 * powb(10, 9-n)
		} else if b[384+i] == 55 {
			iopv += 7 * powb(10, 9-n)
		} else if b[384+i] == 56 {
			iopv += 8 * powb(10, 9-n)
		} else if b[384+i] == 57 {
			iopv += 9 * powb(10, 9-n)
		} else {
		}
	}
	return
}

func getTimeStamp(b []byte, capitaltype string) (_time int32) {
	timestamp := make([]byte, 9)
	if capitaltype == "STOCK" {
		copy(timestamp[0:2], b[381:383])
		copy(timestamp[2:4], b[384:386])
		copy(timestamp[4:6], b[387:389])
		copy(timestamp[6:9], b[390:393])
	} else if capitaltype == "FUND" {
		copy(timestamp[0:2], b[405:407])
		copy(timestamp[2:4], b[408:410])
		copy(timestamp[4:6], b[411:413])
		copy(timestamp[6:9], b[414:417])
	} else if capitaltype == "INDEX" {
		copy(timestamp[0:2], b[131:133])
		copy(timestamp[2:4], b[134:136])
		copy(timestamp[4:6], b[137:139])
		copy(timestamp[6:9], b[140:143])
	}
	timestampbuf, _ := strconv.ParseInt(string(timestamp), 10, 32)
	_time = int32(timestampbuf)
	return
}

func getMkdata(msg interface{}, capitaltype string) hqbase.Marketdata {
	b := msg.([]byte)

	var _status string
	if b[372] == 0x43 { //C => I
		_status = "I"
	} else if b[372] == 0x54 { //T => O
		_status = "O"
	} else if b[372] == 0x50 { //P => B
		_status = "B"
	}

	var iopv int64
	if capitaltype == "FUND" {
		iopv = getIOPV(b)
	}

	var _askPrice = [10]float64{float64(getsp1(b)) / 1000.0,
		float64(getsp2(b)) / 1000.0,
		float64(getsp3(b)) / 1000.0,
		float64(getsp4(b)) / 1000.0,
		float64(getsp5(b)) / 1000.0}
	var _askVol = [10]int32{int32(getsv1(b)),
		int32(getsv2(b)),
		int32(getsv3(b)),
		int32(getsv4(b)),
		int32(getsv5(b))}
	var _bidPrice = [10]float64{float64(getbp1(b)) / 1000.0,
		float64(getbp2(b)) / 1000.0,
		float64(getbp3(b)) / 1000.0,
		float64(getbp4(b)) / 1000.0,
		float64(getbp5(b)) / 1000.0}
	var _bidVol = [10]int32{int32(getbv1(b)),
		int32(getbv2(b)),
		int32(getbv3(b)),
		int32(getbv4(b)),
		int32(getbv5(b))}

	return hqbase.Marketdata{
		Code:     string(b[0:6]) + ".SH",
		Time:     getTimeStamp(b, capitaltype),
		Status:   _status,
		PreClose: float64(getLastprice(b)) / 1000.00,
		Open:     float64(getOpen(b)) / 1000.00,
		High:     float64(getHigh(b)) / 1000.00,
		Low:      float64(getLow(b)) / 1000.00,
		Close:    float64(getLow(b)) / 1000.00,
		Match:    float64(getTradePrice(b)) / 1000.00,
		AskPrice: _askPrice,
		AskVol:   _askVol,
		BidPrice: _bidPrice,
		BidVol:   _bidVol,
		Volume:   getVol(b),
		IOPV:     float64(iopv) / 1000.0,
		Turnover: getAmount(b) / 100}
}
