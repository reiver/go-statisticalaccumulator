package statisticalaccumulator



import (
    "testing"
)



func TestRunningAccumulation(t *testing.T) {

//	tests := []struct {



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
				{"3" , "1/1"  , "3/1"  , "0/1"},
				{"3" , "2/1"  , "3/1"  , "0/1"},
				{"3" , "3/1"  , "3/1"  , "0/1"},
				{"3" , "4/1"  , "3/1"  , "0/1"},
				{"3" , "5/1"  , "3/1"  , "0/1"},
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

			{
				{"3" , "1/1" , "3/1" , "0/1"},
				{"4" , "2/1" , "7/2" , "1/4"},
				{"5" , "3/1" , "4/1" , "2/3"},
				{"6" , "4/1" , "9/2" , "5/4"},
				{"7" , "5/1" , "5/1" , "2/1"},
			},

			{
				{"8"  , "1/1" , "8/1" , "0/1"},
				{"9"  , "2/1" , "17/2" , "1/4"},
				{"10" , "3/1" , "9/1" , "2/3"},
				{"11" , "4/1" , "19/2" , "5/4"},
				{"12" , "5/1" , "10/1" , "2/1"},
			},

			{
				{"-2" , "1/1" , "-2/1" , "0/1"},
				{"-1" , "2/1" , "-3/2" , "1/4"},
				{"0"  , "3/1" , "-1/1" , "2/3"},
				{"1"  , "4/1" , "-1/2" , "5/4"},
				{"2"  , "5/1" , "0/1"  , "2/1"},
			},

			{
				{"15" , "1/1" , "15/1" , "0/1"},
				{"20" , "2/1" , "35/2" , "25/4"},
				{"25" , "3/1" , "20/1" , "50/3"},
				{"30" , "4/1" , "45/2" , "125/4"},
				{"35" , "5/1" , "25/1" , "50/1"},
			},

			{
				{"17" , "1/1" , "17/1" , "0/1"},
				{"19" , "2/1" , "18/1" , "1/1"},
				{"18" , "3/1" , "18/1" , "2/3"},
				{"17" , "4/1" , "71/4" , "11/16"},
				{"19" , "5/1" , "18/1" , "4/5"},
			},

			{
				{"7"  , "1/1" , "7/1"  , "0/1"},
				{"38" , "2/1" , "45/2" , "961/4"},
				{"4"  , "3/1" , "49/3" , "2126/9"},
				{"23" , "4/1" , "18/1" , "371/2"},
				{"18" , "5/1" , "18/1" , "742/5"},
			},

		}



	for testSetNumber,tests := range testsOfTests {

		statisticalAccumulator := NewStatisticalAccumulator()

		for datumNumber,datum := range tests {

			statisticalAccumulator.PutString(datum.x)

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



