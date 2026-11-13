package hackernewsapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupServer(t *testing.T, handler http.HandlerFunc) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	orig := apiEndpoint
	apiEndpoint = srv.URL + "/"
	t.Cleanup(func() { apiEndpoint = orig })
}

func TestGetStoriesReturnsIDs(t *testing.T) {
	setupServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/topstories.json" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte(`[3,2,1]`))
	})

	ids, err := GetStories("Top")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ids) != 3 || ids[0] != 3 {
		t.Fatalf("got %v", ids)
	}
}

func TestGetStoriesRejectsUnknownTarget(t *testing.T) {
	if _, err := GetStories("unknown"); err == nil {
		t.Fatal("expected error for unknown target")
	}
}

func TestGetStoriesReturnsErrorOnHTTPError(t *testing.T) {
	setupServer(t, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	})

	if _, err := GetStories("top"); err == nil {
		t.Fatal("expected error on 500 response")
	}
}

func TestGetItemReturnsErrorWhenServerUnreachable(t *testing.T) {
	orig := apiEndpoint
	apiEndpoint = "http://127.0.0.1:1/"
	t.Cleanup(func() { apiEndpoint = orig })

	if _, err := GetItem(1); err == nil {
		t.Fatal("expected error when server is unreachable")
	}
}

func TestGetItemReturnsErrorOnBrokenJSON(t *testing.T) {
	setupServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{broken`))
	})

	if _, err := GetItem(1); err == nil {
		t.Fatal("expected error on broken JSON")
	}
}

func TestGetItemDecodesStory(t *testing.T) {
	setupServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/item/8863.json" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte(`{"by":"dhouston","id":8863,"time":1175714200,"title":"My YC app","type":"story"}`))
	})

	s, err := GetItem(8863)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.By != "dhouston" || s.ID != 8863 || s.Time != 1175714200 {
		t.Fatalf("got %+v", s)
	}
}

func TestGetUserDecodesNameAndKarma(t *testing.T) {
	setupServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"id":"jl","created":1173923446,"karma":2937,"submitted":[1,2]}`))
	})

	u, err := GetUser("jl")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.ID != "jl" || u.Karma != 2937 {
		t.Fatalf("got %+v", u)
	}
}

func TestGetMaxItemID(t *testing.T) {
	setupServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`12345`))
	})

	n, err := GetMaxItemID()
	if err != nil || n != 12345 {
		t.Fatalf("got %d, %v", n, err)
	}
}

func TestGetUpdates(t *testing.T) {
	setupServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"items":[1,2],"profiles":["a"]}`))
	})

	u, err := GetUpdates()
	if err != nil || len(u.Items) != 2 || u.Profiles[0] != "a" {
		t.Fatalf("got %+v, %v", u, err)
	}
}
