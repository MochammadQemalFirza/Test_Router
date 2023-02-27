package Router

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/genre/:genrename", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		genre := params.ByName("genrename")
		text := "product " + id + " genre " + genre
		fmt.Fprint(w, text)
	})
	r := httptest.NewRequest("GET", "htpp://localhost:8085/product/1/genre/metal", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "product 1 genre metal", string(body))

}
