package trade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trade_UpdateCriteria(t1 *testing.T) {
	type fields struct {
		criteriaOfBuy  float64
		criteriaOfSell float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		isCOBUpdate      bool
		expectedCOBValue int32
		isCOSUpdate      bool
		expectedCOSValue int32
	}{
		{
			name: "COBを更新し、COSを更新しない",
			fields: fields{
				500,
				500,
			},
			args: args{
				700,
			},
			isCOBUpdate:      true,
			expectedCOBValue: 700,
			isCOSUpdate:      false,
			expectedCOSValue: 500,
		},
		{
			name: "COBを更新しない/COSを更新する",
			fields: fields{
				500,
				500,
			},
			args: args{
				400,
			},
			isCOBUpdate:      false,
			expectedCOBValue: 500,
			isCOSUpdate:      true,
			expectedCOSValue: 400,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := trade{
				criteriaOfBuy:  tt.fields.criteriaOfBuy,
				criteriaOfSell: tt.fields.criteriaOfSell,
			}
			isUpdate := t.UpdateCriteriaOfBuy(tt.args.value)
			isCosUpdate := t.UpdateCriteriaOfSell(tt.args.value)
			assert.Equal(t1, tt.expectedCOBValue, t.criteriaOfBuy)
			assert.Equal(t1, tt.isCOBUpdate, isUpdate)
			assert.Equal(t1, tt.expectedCOSValue, t.criteriaOfSell)
			assert.Equal(t1, tt.isCOSUpdate, isCosUpdate)
		})
	}
}

func Test_trade_saveUpdateResult(t1 *testing.T) {
	type fields struct {
		criteriaOfBuy         float64
		updateResultListOfCOB []bool
		criteriaOfSell        float64
		updateResultListOfCOS []bool
	}
	tests := []struct {
		name   string
		fields fields
		args   bool
		want   []bool
	}{
		{
			"updateResultListOfCOBを更新する",
			fields{
				criteriaOfBuy:         0,
				updateResultListOfCOB: []bool{},
				criteriaOfSell:        0,
				updateResultListOfCOS: []bool{},
			},
			true,
			[]bool{true},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &trade{
				criteriaOfBuy:         tt.fields.criteriaOfBuy,
				updateResultListOfCOB: tt.fields.updateResultListOfCOB,
				criteriaOfSell:        tt.fields.criteriaOfSell,
				updateResultListOfCOS: tt.fields.updateResultListOfCOS,
			}
			t.saveUpdateResultOfCOB(tt.args)
			assert.Equal(t1, tt.want, t.updateResultListOfCOB)
		})
	}
}
