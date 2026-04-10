package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/pastequo/experiments.unit-test/internal/domain/repo/mocks"
	"github.com/pastequo/experiments.unit-test/internal/usecase"
)

var (
	errAct = errors.New("error Act")
	errDo  = errors.New("error Do")
)

func TestRunBoth_Good(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actor := mocks.NewMockActor(ctrl)
	doer := mocks.NewMockDoer(ctrl)
	us := usecase.NewActAndDo(actor, doer)

	actor.EXPECT().Act().Return(errAct)

	// Doesn't break if we change the implementation order
	// Theoretically it should be done for all methods of this interface and all other methods of the other interface
	doer.EXPECT().Do().Return(nil).AnyTimes()

	err := us.Run()
	assert.ErrorIs(t, err, errAct)
}

func TestRunBoth_Bad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actor := mocks.NewMockActor(ctrl)
	doer := mocks.NewMockDoer(ctrl)
	us := usecase.NewActAndDo(actor, doer)

	actor.EXPECT().Act().Return(errAct)

	// This is implementation dependent, and we don't want to test it
	// It's an implementation detail
	// It would break if we change the implementation order
	doer.EXPECT().Do().Return(nil).Times(0) // Or commented

	err := us.Run()
	assert.ErrorIs(t, err, errAct)
}
