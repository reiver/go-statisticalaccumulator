package statisticalaccumulator



import "math/big"



//func (me *StatisticalAccumulator) PutFloat64(x float64) {
//
//	xx := new(big.Rat)
//	xx.SetFloat64(x)
//
//	me.PutRat(xx)
//}

func (me *StatisticalAccumulator) PutFrac(a,b *big.Int) {

	xx := new(big.Rat)
	xx.SetFrac(a,b)

	me.PutRat(xx)
}

func (me *StatisticalAccumulator) PutFrac64(a,b int64) {

	xx := new(big.Rat)
	xx.SetFrac64(a,b)

	me.PutRat(xx)
}

func (me *StatisticalAccumulator) PutInt(x *big.Int) {

	xx := new(big.Rat)
	xx.SetInt(x)

	me.PutRat(xx)
}

func (me *StatisticalAccumulator) PutInt64(x int64) {

	xx := new(big.Rat)
	xx.SetInt64(x)

	me.PutRat(xx)
}

func (me *StatisticalAccumulator) PutString(x string) {

	xx := new(big.Rat)
	xx.SetString(x)

	me.PutRat(xx)
}



func (me *StatisticalAccumulator) PutRat(x *big.Rat) {

	me.n.Add(me.n, me.one)



	me.sigmaXI.Add(me.sigmaXI, x)



	xSquared := new(big.Rat)
	xSquared.Mul(x, x)
	
	me.sigmaXISquared.Add(me.sigmaXISquared, xSquared)
}
