package typings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

// attach a function to a user-defined type, so that it acts like
// a method in OOP

type _WebPage struct {
	url  string
	body []byte
	err  error
}

func (w *_WebPage) get() {
	if resp, err := http.Get(w.url); err == nil {
		defer resp.Body.Close()
		if w.body, err = ioutil.ReadAll(resp.Body); err != nil {
			w.err = err
		}
	} else {
		w.err = err
	}
}

func (w *_WebPage) isOk() bool {
	return w.err == nil
}

func TestWebPageGet(t *testing.T) {
	// using the "tuple form"
	//w := &_WebPage{"http://www.google.com", nil, nil}
	// using the field-accessor
	w := &_WebPage{url: "http://www.google.com"}
	w.get()
	assert.Equal(t, nil, w.err)
	assert.True(t, w.isOk())
	fmt.Printf("%d\n", len(w.body))

	// create a zero-struct (all members init-ed with zero values)
	zero := _WebPage{}
	zero.get()
	assert.False(t, zero.isOk())

	// to construct a new web page struct - commonly used in constructor
	// functions
	// https://stackoverflow.com/questions/18125625/constructors-in-go
	//w2 := new(_WebPage)
}
