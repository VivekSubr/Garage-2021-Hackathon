package yangserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	demo "hack.com/structs"
)

var yangRoot = demo.SchemaTree["Device"]
var yangData interface{}

func validateAnimal(token string, json string) error {
	fmt.Printf("validating %s\n", token)

	switch token {
	case "cat":
		nCat := &demo.Cat{}
		if err := demo.Unmarshal([]byte(json), nCat); err != nil {
			return err
		}

		return nCat.Validate()
	case "dog":
		nDog := &demo.Dog{}
		if err := demo.Unmarshal([]byte(json), nDog); err != nil {
			return err
		}

		return nDog.Validate()

	case "tiger":
		nTiger := &demo.Tiger{}
		if err := demo.Unmarshal([]byte(json), nTiger); err != nil {
			return err
		}

		return nTiger.Validate()
	}

	return nil
}

func updateAnimal(token string, jMap map[string]string) error {
	jstr, err := json.Marshal(jMap)
	if err != nil {
		return err
	}

	err = validateAnimal(token, string(jstr))
	if err != nil {
		return err
	}

	yangData.(map[string]interface{})["animal"].(map[string]interface{})[token] = string(jstr)
	return nil
}

func serveYangFunc(w http.ResponseWriter, req *http.Request) {
	s := strings.Split(req.URL.Path, "/")
	dataEle := yangData
	for _, tok := range s[1:] {
		dataEle = dataEle.(map[string]interface{})[tok]
	}

	switch req.Method {
	case "GET":
		w.WriteHeader(200)
		w.Write([]byte(dataEle.(string)))
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var t map[string]string
		err := decoder.Decode(&t)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Printf("POST %v\n", t)
		err = updateAnimal(s[2], t)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(200)
	}
}

func ServeYang() {
	yangData = make(map[string]interface{})
	yangData.(map[string]interface{})["animal"] = make(map[string]interface{})
	yangData.(map[string]interface{})["animal"].(map[string]interface{})["cat"] = "{\"does\" : \"meow\" }"
	yangData.(map[string]interface{})["animal"].(map[string]interface{})["dog"] = "{\"does\" : \"bark\" }"
	yangData.(map[string]interface{})["animal"].(map[string]interface{})["tiger"] = "{\"does\" : \"roar\" }"

	http.HandleFunc("/", serveYangFunc)

	fmt.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
