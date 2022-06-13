package trade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trade_UpdateCriteria(t1 *testing.T) {
	type fields struct {
		criteriaOfBuy float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		isCOBUpdate      bool
		expectedCOBValue float64
	}{
		{
			name: "COBを更新する",
			fields: fields{
				500,
			},
			args: args{
				400,
			},
			isCOBUpdate:      true,
			expectedCOBValue: 400,
		},
		{
			name: "COBを更新しない",
			fields: fields{
				500,
			},
			args: args{
				600,
			},
			isCOBUpdate:      false,
			expectedCOBValue: 500,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := criteria{
				criteriaOfBuy: tt.fields.criteriaOfBuy,
			}
			isUpdate := t.UpdateCriteriaOfBuy(tt.args.value)
			assert.Equal(t1, tt.expectedCOBValue, t.criteriaOfBuy)
			assert.Equal(t1, tt.isCOBUpdate, isUpdate)
		})
	}
}

func Test_trade_UpdateCriteriaOfSell(t1 *testing.T) {
	type fields struct {
		criteriaOfSell float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		isCOSUpdate      bool
		expectedCOSValue float64
	}{
		{
			name: "COSを更新する",
			fields: fields{
				500,
			},
			args: args{
				600,
			},
			isCOSUpdate:      true,
			expectedCOSValue: 600,
		},
		{
			name: "COSを更新しない",
			fields: fields{
				500,
			},
			args: args{
				400,
			},
			isCOSUpdate:      false,
			expectedCOSValue: 500,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := criteria{
				criteriaOfSell: tt.fields.criteriaOfSell,
			}
			isUpdate := t.UpdateCriteriaOfSell(tt.args.value)
			assert.Equal(t1, tt.expectedCOSValue, t.criteriaOfSell)
			assert.Equal(t1, tt.isCOSUpdate, isUpdate)
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
			t := &criteria{
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

func Test_position_DecidePosition(t *testing.T) {
	type args struct {
		trade *criteria
		value float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		wantStr string
	}{
		{
			name: "value が値上がりして、COBが更新されず、買いポジションになる",
			args: args{
				&criteria{
					criteriaOfBuy:         500.0,
					updateResultListOfCOB: []bool{true},
					criteriaOfSell:        0,
					updateResultListOfCOS: []bool{false},
				},
				600,
			},
			wantStr: "buy",
			wantErr: nil,
		},
		{
			name: "value が値下がりして、COSが更新されず、売りポジションになる",
			args: args{
				&criteria{
					criteriaOfBuy:         0,
					updateResultListOfCOB: []bool{false},
					criteriaOfSell:        500.0,
					updateResultListOfCOS: []bool{true},
				},
				100,
			},
			wantStr: "sell",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := position{}
			actual, _ := p.DecidePosition(tt.args.trade, tt.args.value)
			assert.Equal(t, tt.wantStr, actual)
		})
	}
}
