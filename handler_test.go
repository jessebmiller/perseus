package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestRequestResponse tests requests and responses
func TestRequestResponse(t *testing.T) {
	DEPS = make(map[string]interface{})
	tsets := []struct {
		req      *http.Request // a request
		inStore  Store         // given some state
		outStore Store         // should leave the store in this state
		resp     string        // should write this to the response writer
	}{
		{
			httptest.NewRequest("GET", "/abc", nil),
			MapStore{
				map[string][]string{
					"/abc": {"a", "b", "c"},
				},
			},
			MapStore{
				map[string][]string{
					"/abc": {"a", "b", "c"},
				},
			},
			"a\nb\nc",
		},
		{
			httptest.NewRequest("POST", "/abc?message=d", nil),
			MapStore{
				map[string][]string{
					"/abc": {"a", "b", "c"},
				},
			},
			MapStore{
				map[string][]string{
					"/abc": {"a", "b", "c", "d"},
				},
			},
			"1",
		},
	}

	for _, tset := range tsets {
		DEPS["store"] = tset.inStore
		w := httptest.NewRecorder()
		rootHandler(w, tset.req)

		if w.Body.String() != tset.resp {
			t.Errorf(
				"rootHandler responded %v, want %v",
				w.Body.String(),
				tset.resp,
			)
		}

		if !reflect.DeepEqual(DEPS["store"], tset.outStore) {
			t.Errorf(
				"store was left as %v, wanted %v",
				DEPS["store"],
				tset.outStore,
			)
		}
	}
}
