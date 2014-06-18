package Frontend

import (
	"Datastore"
	"appengine"
	"net/http"
	"fmt"
)

type Test struct {
	Test1 string
	Test2 string
}

func init() {
	http.HandleFunc("/", hello)
}

func hello(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("hello")
	
	
	err := Datastore.ClearNamespaceAndMemcache(c, "TEST")
	if err!=nil {
		c.Errorf("Error 1 - "+err.Error())
		return
	}
	
	test1:=Test{"123", "456"}
	
	_, err=Datastore.PutInDatastoreSimpleAndMemcache(c, "TEST", "TEST", "TEST", &test1)
	if err!=nil {
		c.Errorf("Error 2 - "+err.Error())
		return
	}
	
	list := []Test{}
	keys, err := Datastore.QueryGetAllKeys(c, "TEST", &list)
	c.Infof("keys - %v", keys)
	fmt.Fprintf(w, "keys - %v", keys)
}
