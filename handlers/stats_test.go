package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const text = `Mention is made in many alchemical writings of a mythical personage named Hermes
Trismegistus, who is said to have lived a little later than the time of Moses.
Representations of Hermes Trismegistus are found on ancient Egyptian monuments. We
are told that Alexander the Great found his tomb near Hebron; and that the tomb
contained a slab of emerald whereon thirteen sentences were written. The eighth
sentence is rendered many alchemical books as follows:
"Ascend with the greatest sagacity from the earth to heaven, and then again
descend to the earth, and unite together the powers of things superior and things
inferior. Thus you will obtain the glory of the whole world, and obscurity will fly
away from you."
This sentence evidently teaches the unity of things in heaven and things on
earth, and asserts the possibility of gaining, not merely a theoretical, but also a
practical, knowledge of the essential characters of all things. Moreover, the
sentence implies that this fruitful knowledge is be obtained by examining

nature, using as guide the fundamental similarity supposed to exist between things
above and things beneath.
The alchemical writers constantly harp on this theme: follow nature; provided
you never lose the clue, which is simplicity and similarity.
The author of The Only Way (1677) beseeches his readers "to enlist under the
standard of that method which proceeds in strict obedience to the teaching of
nature ... in short, the method which nature herself pursues in the bowels of the
earth."`

func TestStatsWithWrongMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Stats)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusNotFound)
	}

	expected := "404 page not found"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestStatsWithWrongBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/stats", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Stats)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
	}

	expected := "500 internal server error"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestStats(t *testing.T) {
	req, err := http.NewRequest("POST", "/stats", strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Stats)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content-type: got %v want %v",
			rr.Code, http.StatusOK)
	}

	expected := `{"stats":{"lines":22,"characters":1273,"words":247,"top_5_words":["the","of","and","things","to"]}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
