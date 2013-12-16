package statisticalaccumulator



import "math/big"



type StatisticalAccumulator struct {
	n              *big.Rat
	sigmaXI        *big.Rat
	sigmaXISquared *big.Rat

	one *big.Rat
}



func New() (*StatisticalAccumulator) {

	n := new(big.Rat)
	n.SetInt64(0)

	sigmaXI := new(big.Rat)
	sigmaXI.SetInt64(0)

	sigmaXISquared := new(big.Rat)
	sigmaXISquared.SetInt64(0)

	one := new(big.Rat)
	one.SetInt64(1)



	me := StatisticalAccumulator{
		n:n,
		sigmaXI:sigmaXI,
		sigmaXISquared:sigmaXISquared,
		one:one,
	}



	return &me
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



func (me *StatisticalAccumulator) N() (*big.Rat) {
	n := new(big.Rat)
	n.Set(me.n)



	return n
}

func (me *StatisticalAccumulator) Mean() (*big.Rat) {
	mean := new(big.Rat)

	mean.Inv(me.n)

	mean.Mul(mean, me.sigmaXI)



	return mean
}


func (me *StatisticalAccumulator) Variance() (*big.Rat) {
	variance := new(big.Rat)

	variance.Inv(me.n)

	variance.Mul(variance, me.sigmaXISquared)



	temp := new(big.Rat)

	temp.Mul(me.n, me.n)
	temp.Inv(temp)

	temp.Mul(temp, me.sigmaXI)
	temp.Mul(temp, me.sigmaXI)


	variance.Sub(variance, temp)



	return variance
}
