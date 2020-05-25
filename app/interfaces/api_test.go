package interfaces

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/itiroinazawa/golangapi/app/mocks"
	"github.com/itiroinazawa/golangapi/app/usecases"
	"github.com/stretchr/testify/assert"
)

func Test_GetAccount_WithValidData(t *testing.T) {

	mockCtrl := gomock.NewController(t)

	iterator := mocks.NewMockInteractor(mockCtrl)

	expected := usecases.Account{ID: 1, DocumentNumber: "12345"}

	iterator.EXPECT().GetAccount(gomock.Any()).DoAndReturn(func(accountID int) (usecases.Account, error) {
		return expected, nil
	})

	req, _ := http.NewRequest("GET", "/accounts/1", nil)
	rr := httptest.NewRecorder()

	api := ApiHandler{Interactor: iterator}

	api.GetAccount(rr, req)

	var acc usecases.Account
	body, _ := ioutil.ReadAll(rr.Body)

	json.Unmarshal(body, &acc)

	assert.Equal(t, expected, acc)
}

func Test_GetAccount_WithInvalidData(t *testing.T) {

	mockCtrl := gomock.NewController(t)

	iterator := mocks.NewMockInteractor(mockCtrl)
	var expected usecases.Account

	iterator.EXPECT().GetAccount(gomock.Any()).DoAndReturn(func(accountID int) (usecases.Account, error) {
		return expected, fmt.Errorf("Account not found")
	})

	req, _ := http.NewRequest("GET", "/accounts/2", nil)
	rr := httptest.NewRecorder()

	api := ApiHandler{Interactor: iterator}

	api.GetAccount(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Result().StatusCode)
}
