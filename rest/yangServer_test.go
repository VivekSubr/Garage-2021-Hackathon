package yangserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServe(t *testing.T) {
	go ServeYang()
	time.Sleep(2 * time.Second)

	url := "http://localhost:3000/animal/cat"
	rsp, err := http.Get(url)
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode)

	body, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(t, err)
	fmt.Printf("Rsp1 : %v\n", string(body))

	nCat := map[string]interface{}{"does": "meow1"}
	bCat, err := json.Marshal(nCat)
	assert.Nil(t, err)

	rsp, err = http.Post(url, "application/json", bytes.NewBuffer(bCat))
	assert.Nil(t, err)
	assert.Equal(t, 500, rsp.StatusCode)

	nCat = map[string]interface{}{"does": "meowwwwww"}
	bCat, err = json.Marshal(nCat)
	assert.Nil(t, err)

	rsp, err = http.Post(url, "application/json", bytes.NewBuffer(bCat))
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode)

	rsp, err = http.Get(url)
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode)
}
