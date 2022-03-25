package usecase

type Judge interface {
	Increase()
	Decrease()
}

type judgement struct {
}

// TODO ロジックを実装する
func (j judgement) Increase() {

}

func (j judgement) Decrease() {

}
