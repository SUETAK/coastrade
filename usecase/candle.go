package usecase

// Candle についてのロジックを提供する

type CandleUseCase interface {
	Answer() (string, error)
}

type CandleUsecase struct {
}

func (receiver CandleUsecase) Answer() (string, error) {
	return "Answer関数が呼ばれました", nil
}
