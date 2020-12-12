package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/g1stavo/taco-api/entity"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func prepareTest() (api *apitest.APITest) {
	r := NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()
	api = apitest.New().Handler(r)
	return
}

func TestGetTaco(t *testing.T) {
	const path = "/taco"

	t.Run("Should return status OK", func(t *testing.T) {
		api := prepareTest()
		api.
			Get(path).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("Should return server error when Taco API fails", func(t *testing.T) {
		mock := apitest.NewMock().
			Get("/random").
			RespondWith().
			Status(http.StatusNotFound).
			End()

		api := prepareTest()
		api.
			Mocks(mock).
			Get(path).
			Expect(t).
			Status(http.StatusInternalServerError).
			End()
	})

	t.Run("Should get taco", func(t *testing.T) {
		api := prepareTest()
		api.
			Get("/taco").
			Expect(t).
			Assert(func(res *http.Response, req *http.Request) error {
				defer res.Body.Close()
				data, _ := ioutil.ReadAll(res.Body)

				var taco entity.Taco
				json.Unmarshal(data, &taco)

				assert.NotNil(t, taco.Seasoning.Name)
				assert.NotNil(t, taco.Seasoning.URL)
				assert.NotNil(t, taco.Condiment.Name)
				assert.NotNil(t, taco.Condiment.URL)
				assert.NotNil(t, taco.Mixin.Name)
				assert.NotNil(t, taco.Mixin.URL)
				assert.NotNil(t, taco.BaseLayer.Name)
				assert.NotNil(t, taco.BaseLayer.URL)
				assert.NotNil(t, taco.Shell.Name)
				assert.NotNil(t, taco.Shell.URL)

				return nil
			}).
			End()
	})
}
