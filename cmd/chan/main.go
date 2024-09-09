package main

import (
	"cases/internal/controller"
	"cases/internal/infrastructure"
	"cases/internal/presenter"
	"cases/internal/usecase"
	"context"
)

func main() {
	subscriber := &controller.ChanSubscriber{
		UseCase: &usecase.ChanUseCase{
			Uploader: &infrastructure.Uploader{},
		},
		Presenter: &presenter.Presenter{},
	}

	subscriber.UploadFiles(context.Background(), &controller.Event{})
}
