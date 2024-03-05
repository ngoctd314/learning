package main

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_notifier_onlyGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockstorage(ctrl)
	mock.EXPECT().Get().Return(gomock.Any(), nil)
	// must not run Create()

	notifier := newNotifier(mock)
	_, err := notifier.getOrCreate()

	if err != nil {
		t.Errorf("want err == nil, got err = %v", err)
	}
}

func Test_notifier_getAndCreateWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockstorage(ctrl)
	mock.EXPECT().Get().Return(gomock.Any(), errors.New("not found")).Times(1)
	createErr := errors.New("create error")
	mock.EXPECT().Create().Return(nil, createErr).Times(1)

	notifier := newNotifier(mock)
	_, err := notifier.getOrCreate()

	if err != createErr {
		t.Errorf("want err == %v, got err = %v", createErr, err)
	}
}

func Test_notifier_getAndCreateWithoutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockstorage(ctrl)
	mock.EXPECT().Get().Return(nil, errors.New("not found")).Times(1)
	mock.EXPECT().Create().Return(gomock.Any(), nil).Times(1)

	notifier := newNotifier(mock)
	_, err := notifier.getOrCreate()

	if err != nil {
		t.Errorf("want err == nil, got err = %v", err)
	}
}
