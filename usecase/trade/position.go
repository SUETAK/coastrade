package trade

func NewTrade(criteriaOfBuy int32, criteriaOfSell int32) *trade {
	return &trade{criteriaOfBuy: criteriaOfBuy, criteriaOfSell: criteriaOfSell}
}

type trade struct {
	criteriaOfBuy  int32
	criteriaOfSell int32
}

func (t *trade) UpdateCriteriaOfBuy(value int32) bool {
	if t.criteriaOfBuy-value < 0 {
		t.criteriaOfBuy = value
		return true
	}
	return false
}

func (t *trade) UpdateCriteriaOfSell(value int32) bool {
	if t.criteriaOfSell-value < 0 {
		return false
	}
	t.criteriaOfSell = value
	return true
}

type decision interface {
}
