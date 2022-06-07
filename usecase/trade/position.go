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
	if t.criteriaOfBuy-value < 0 {
		t.criteriaOfBuy = value
		return true
	}
	return false
}

func (t *trade) UpdateCriteriaOfSell(value float64) bool {
	if t.criteriaOfSell-value < 0 {
		return false
	}
	t.criteriaOfSell = value
	return true
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

func (p position) DecidePosition(trade *trade) error {

	// 現在のETHの値を取得する(Ticker)
	value := 123.123
	// 基準値を更新するかどうか決める関数を呼ぶ(trade)
	cobResult := trade.UpdateCriteriaOfBuy(value)
	trade.saveUpdateResultOfCOB(cobResult)
	cosResult := trade.UpdateCriteriaOfSell(value)
	trade.saveUpdateResultOfCOS(cosResult)
	// 1. COBの最新とその前の値を見て、true->false になったか確認する
	if cobResult != trade.updateResultListOfCOB[len(trade.updateResultListOfCOB)-1] {
		// {1}. がtrue, {2}. がfalse の時に差分が1% 以上かどうかをみて、1%以上なら購入
		if value*0.01 > trade.criteriaOfBuy {
			// 購入
		}
	}
	// 2. COSの最新とその前の値を見て、true->false になったか確認する
	if cosResult != trade.updateResultListOfCOS[len(trade.updateResultListOfCOS)-1] {
		// {2}. がtrue, {1}. がfalse の時に差分が1% 以上かどうかをみて、1%以下なら売却
		if value*0.01 < trade.criteriaOfSell {
			// 売却
		}
	}
	return nil
}
