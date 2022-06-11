package trade

func NewTrade(criteriaOfBuy float64, criteriaOfSell float64) *trade {
	return &trade{criteriaOfBuy: criteriaOfBuy, criteriaOfSell: criteriaOfSell}
}

type trade struct {
	criteriaOfBuy         float64
	updateResultListOfCOB []bool
	criteriaOfSell        float64
	updateResultListOfCOS []bool
}

func (t *trade) UpdateCriteriaOfBuy(value float64) bool {
	if value < t.criteriaOfBuy {
		t.criteriaOfBuy = value
		return true
	}
	return false
}

func (t *trade) UpdateCriteriaOfSell(value float64) bool {
	if t.criteriaOfSell < value {
		t.criteriaOfSell = value
		return true
	}
	return false
}

func (t *trade) saveUpdateResultOfCOB(updateResult bool) {
	t.updateResultListOfCOB = append(t.updateResultListOfCOB, updateResult)
}

func (t *trade) saveUpdateResultOfCOS(updateResult bool) {
	t.updateResultListOfCOS = append(t.updateResultListOfCOS, updateResult)
}

type Decide interface {
	DecidePosition() error
}

type position struct{}

func (p position) DecidePosition(trade *trade, value float64) (string, error) {

	// 現在のETHの値を取得する(Ticker)
	// 基準値を更新するかどうか決める関数を呼ぶ(trade)
	cobResult := trade.UpdateCriteriaOfBuy(value)
	trade.saveUpdateResultOfCOB(cobResult)
	cosResult := trade.UpdateCriteriaOfSell(value)
	trade.saveUpdateResultOfCOS(cosResult)
	resultLastIndex := trade.updateResultListOfCOB[len(trade.updateResultListOfCOB)-2]
	if cobResult != resultLastIndex {
		if value >= trade.criteriaOfBuy*1.01 {
			return "buy", nil
		}
	}
	if cosResult != trade.updateResultListOfCOS[len(trade.updateResultListOfCOS)-2] {
		if value <= trade.criteriaOfSell*0.99 {
			return "sell", nil
		}
	}
	return "", nil
}
