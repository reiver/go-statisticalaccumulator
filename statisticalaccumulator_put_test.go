package statisticalaccumulator



import (
    "testing"
	"math/big"
)


func TestPutFrac(t *testing.T) {

	testsOfTests := [][]struct {
			a                string
			b                string
			expectedN        string
			expectedMean     string
			expectedVariance string
		}{
			{
				{"5" , "3" , "1/1"  , "5/3"  , "0/1"},
				{"5" , "3" , "2/1"  , "5/3"  , "0/1"},
				{"5" , "3" , "3/1"  , "5/3"  , "0/1"},
				{"5" , "3" , "4/1"  , "5/3"  , "0/1"},
			},

			{
				{"4" , "7" , "1/1"  , "4/7"  , "0/1"},
			},

			{
				{"1" , "1" , "1/1" , "1/1" , "0/1"},
				{"2" , "1" , "2/1" , "3/2" , "1/4"},
				{"3" , "1" , "3/1" , "2/1" , "2/3"},
				{"4" , "1" , "4/1" , "5/2" , "5/4"},
				{"5" , "1" , "5/1" , "3/1" , "2/1"},
			},

		}


	for testSetNumber,tests := range testsOfTests {

		statisticalAccumulator := New()

		for datumNumber,datum := range tests {

			a := new(big.Int)
			a.SetString(datum.a, 10)

			b := new(big.Int)
			b.SetString(datum.b, 10)

			statisticalAccumulator.PutFrac(a,b)

			if actualN := statisticalAccumulator.N() ; actualN.String() != datum.expectedN {
				t.Errorf("With test set number [%v], when putting [%v/%v] (i.e., datum number [%v]) into the statistical accumulator, expected N to be [%v], but instead got [%v].", testSetNumber, datum.a, datum.b, datumNumber, datum.expectedN, actualN)
				continue
			}

			if actualMean := statisticalAccumulator.Mean() ; actualMean.String() != datum.expectedMean {
				t.Errorf("With test set number [%v], when putting [%v/%v] (i.e., datum number [%v]) into the statistical accumulator, expected Mean to be [%v], but instead got [%v].", testSetNumber, datum.a, datum.b, datumNumber, datum.expectedMean, actualMean)
				continue
			}

			if actualVariance := statisticalAccumulator.Variance() ; actualVariance.String() != datum.expectedVariance {
				t.Errorf("With test set number [%v], when putting [%v/%v] (i.e., datum number [%v]) into the statistical accumulator, expected Variance to be [%v], but instead got [%v].", testSetNumber, datum.a, datum.b, datumNumber, datum.expectedVariance, actualVariance)
				continue
			}
		}
	}

}



func TestPutFrac64(t *testing.T) {

	testsOfTests := [][]struct {
			a                int64
			b                int64
			expectedN        string
			expectedMean     string
			expectedVariance string
		}{
			{
				{int64(5) , int64(3) , "1/1"  , "5/3"  , "0/1"},
				{int64(5) , int64(3) , "2/1"  , "5/3"  , "0/1"},
				{int64(5) , int64(3) , "3/1"  , "5/3"  , "0/1"},
				{int64(5) , int64(3) , "4/1"  , "5/3"  , "0/1"},
			},

			{
				{int64(4) , int64(7) , "1/1"  , "4/7"  , "0/1"},
			},

			{
				{int64(1) , int64(1) , "1/1" , "1/1" , "0/1"},
				{int64(2) , int64(1) , "2/1" , "3/2" , "1/4"},
				{int64(3) , int64(1) , "3/1" , "2/1" , "2/3"},
				{int64(4) , int64(1) , "4/1" , "5/2" , "5/4"},
				{int64(5) , int64(1) , "5/1" , "3/1" , "2/1"},
			},

		}


	for testSetNumber,tests := range testsOfTests {

		statisticalAccumulator := New()

		for datumNumber,datum := range tests {

			a := datum.a

			b := datum.b

			statisticalAccumulator.PutFrac64(a,b)

			if actualN := statisticalAccumulator.N() ; actualN.String() != datum.expectedN {
				t.Errorf("With test set number [%v], when putting [%v/%v] (i.e., datum number [%v]) into the statistical accumulator, expected N to be [%v], but instead got [%v].", testSetNumber, datum.a, datum.b, datumNumber, datum.expectedN, actualN)
				continue
			}

			if actualMean := statisticalAccumulator.Mean() ; actualMean.String() != datum.expectedMean {
				t.Errorf("With test set number [%v], when putting [%v/%v] (i.e., datum number [%v]) into the statistical accumulator, expected Mean to be [%v], but instead got [%v].", testSetNumber, datum.a, datum.b, datumNumber, datum.expectedMean, actualMean)
				continue
			}

			if actualVariance := statisticalAccumulator.Variance() ; actualVariance.String() != datum.expectedVariance {
				t.Errorf("With test set number [%v], when putting [%v/%v] (i.e., datum number [%v]) into the statistical accumulator, expected Variance to be [%v], but instead got [%v].", testSetNumber, datum.a, datum.b, datumNumber, datum.expectedVariance, actualVariance)
				continue
			}
		}
	}

}



func TestPutInt(t *testing.T) {

	testsOfTests := [][]struct {
			x                string
			expectedN        string
			expectedMean     string
			expectedVariance string
		}{
			{
				{"5" , "1/1"  , "5/1"  , "0/1"},
				{"5" , "2/1"  , "5/1"  , "0/1"},
				{"5" , "3/1"  , "5/1"  , "0/1"},
				{"5" , "4/1"  , "5/1"  , "0/1"},
			},

			{
				{"4" , "1/1"  , "4/1"  , "0/1"},
			},

			{
				{"1" , "1/1" , "1/1" , "0/1"},
				{"2" , "2/1" , "3/2" , "1/4"},
				{"3" , "3/1" , "2/1" , "2/3"},
				{"4" , "4/1" , "5/2" , "5/4"},
				{"5" , "5/1" , "3/1" , "2/1"},
			},

		}


	for testSetNumber,tests := range testsOfTests {

		statisticalAccumulator := New()

		for datumNumber,datum := range tests {

			x := new(big.Int)
			x.SetString(datum.x, 10)

			statisticalAccumulator.PutInt(x)

			if actualN := statisticalAccumulator.N() ; actualN.String() != datum.expectedN {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected N to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedN, actualN)
				continue
			}

			if actualMean := statisticalAccumulator.Mean() ; actualMean.String() != datum.expectedMean {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected Mean to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedMean, actualMean)
				continue
			}

			if actualVariance := statisticalAccumulator.Variance() ; actualVariance.String() != datum.expectedVariance {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected Variance to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedVariance, actualVariance)
				continue
			}
		}
	}

}



func TestPutInt64(t *testing.T) {

	testsOfTests := [][]struct {
			x                int64
			expectedN        string
			expectedMean     string
			expectedVariance string
		}{
			{
				{int64(5) , "1/1"  , "5/1"  , "0/1"},
				{int64(5) , "2/1"  , "5/1"  , "0/1"},
				{int64(5) , "3/1"  , "5/1"  , "0/1"},
				{int64(5) , "4/1"  , "5/1"  , "0/1"},
			},

			{
				{int64(4) , "1/1"  , "4/1"  , "0/1"},
			},

			{
				{int64(1) , "1/1" , "1/1" , "0/1"},
				{int64(2) , "2/1" , "3/2" , "1/4"},
				{int64(3) , "3/1" , "2/1" , "2/3"},
				{int64(4) , "4/1" , "5/2" , "5/4"},
				{int64(5) , "5/1" , "3/1" , "2/1"},
			},

		}


	for testSetNumber,tests := range testsOfTests {

		statisticalAccumulator := New()

		for datumNumber,datum := range tests {

			x := datum.x

			statisticalAccumulator.PutInt64(x)

			if actualN := statisticalAccumulator.N() ; actualN.String() != datum.expectedN {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected N to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedN, actualN)
				continue
			}

			if actualMean := statisticalAccumulator.Mean() ; actualMean.String() != datum.expectedMean {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected Mean to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedMean, actualMean)
				continue
			}

			if actualVariance := statisticalAccumulator.Variance() ; actualVariance.String() != datum.expectedVariance {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected Variance to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedVariance, actualVariance)
				continue
			}
		}
	}

}



func TestPutString(t *testing.T) {

	testsOfTests := [][]struct {
			x                string
			expectedN        string
			expectedMean     string
			expectedVariance string
		}{
			{
				{"1/2" , "1/1"  , "1/2"  , "0/1"},
				{"0.5" , "2/1"  , "1/2"  , "0/1"},
				{"0.50" , "3/1"  , "1/2"  , "0/1"},
				{"2/4" , "4/1"  , "1/2"  , "0/1"},
			},

			{
				{"4/7" , "1/1"  , "4/7"  , "0/1"},
			},

			{
				{"1", "1/1" , "1/1" , "0/1"},
				{"2", "2/1" , "3/2" , "1/4"},
				{"3", "3/1" , "2/1" , "2/3"},
				{"4", "4/1" , "5/2" , "5/4"},
				{"5", "5/1" , "3/1" , "2/1"},
			},

		}


	for testSetNumber,tests := range testsOfTests {

		statisticalAccumulator := New()

		for datumNumber,datum := range tests {

			x := datum.x

			statisticalAccumulator.PutString(x)

			if actualN := statisticalAccumulator.N() ; actualN.String() != datum.expectedN {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected N to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedN, actualN)
				continue
			}

			if actualMean := statisticalAccumulator.Mean() ; actualMean.String() != datum.expectedMean {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected Mean to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedMean, actualMean)
				continue
			}

			if actualVariance := statisticalAccumulator.Variance() ; actualVariance.String() != datum.expectedVariance {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected Variance to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedVariance, actualVariance)
				continue
			}
		}
	}

}



func TestPutRat(t *testing.T) {

	testsOfTests := [][]struct {
			x                string
			expectedN        string
			expectedMean     string
			expectedVariance string
		}{
			{
				{"5/3" , "1/1"  , "5/3"  , "0/1"},
				{"5/3" , "2/1"  , "5/3"  , "0/1"},
				{"5/3" , "3/1"  , "5/3"  , "0/1"},
				{"5/3" , "4/1"  , "5/3"  , "0/1"},
			},

			{
				{"4/7" , "1/1"  , "4/7"  , "0/1"},
			},

			{
				{"1"   , "1/1" , "1/1" , "0/1"},
				{"2/1" , "2/1" , "3/2" , "1/4"},
				{"3.0" , "3/1" , "2/1" , "2/3"},
				{"8/2" , "4/1" , "5/2" , "5/4"},
				{"5"   , "5/1" , "3/1" , "2/1"},
			},

		}


	for testSetNumber,tests := range testsOfTests {

		statisticalAccumulator := New()

		for datumNumber,datum := range tests {

			x := new(big.Rat)
			x.SetString(datum.x)

			statisticalAccumulator.PutRat(x)

			if actualN := statisticalAccumulator.N() ; actualN.String() != datum.expectedN {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected N to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedN, actualN)
				continue
			}

			if actualMean := statisticalAccumulator.Mean() ; actualMean.String() != datum.expectedMean {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected Mean to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedMean, actualMean)
				continue
			}

			if actualVariance := statisticalAccumulator.Variance() ; actualVariance.String() != datum.expectedVariance {
				t.Errorf("With test set number [%v], when putting [%v] (i.e., datum number [%v]) into the statistical accumulator, expected Variance to be [%v], but instead got [%v].", testSetNumber, datum.x, datumNumber, datum.expectedVariance, actualVariance)
				continue
			}
		}
	}

}
