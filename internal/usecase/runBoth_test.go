package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/pastequo/experiments.unit-test/internal/domain/repo/mocks"
	"github.com/pastequo/experiments.unit-test/internal/usecase"
	"github.com/pastequo/experiments.unit-test/internal/utils/observability"
	observabilitymocks "github.com/pastequo/experiments.unit-test/internal/utils/observability/mocks"
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
	// We don't want to test the metrics implementation, so we use a dummy implementation
	metrics := observability.DummyMetrics{}
	us := usecase.NewActAndDo(actor, doer, metrics)

	actor.EXPECT().Act().Return(errAct)

	// Doesn't break if we change call order between Act and Do.
	// Relax mock expectations on collaborators that aren't the focus of this assertion.
	doer.EXPECT().Do().Return(nil).AnyTimes()

	err := us.Run()
	assert.ErrorIs(t, err, errAct)
}

func TestRunBoth_Bad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actor := mocks.NewMockActor(ctrl)
	doer := mocks.NewMockDoer(ctrl)
	// We don't want to test this implementation in this test case
	// Mock is wrong here because it will force us to fill EXPECT for it
	metrics := observabilitymocks.NewMockMetrics(ctrl)

	us := usecase.NewActAndDo(actor, doer, metrics)

	actor.EXPECT().Act().Return(errAct)

	// This is implementation dependent, and we don't want to test it
	// It's an implementation detail
	// It would break if we change the implementation order
	doer.EXPECT().Do().Return(nil).Times(0) // Or commented

	// This is not something that we want to test here
	metrics.EXPECT().Increment("usecase.Run").Times(1)

	err := us.Run()
	assert.ErrorIs(t, err, errAct)
}
