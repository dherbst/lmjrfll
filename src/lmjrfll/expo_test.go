package lmjrfll

import (
	"google.golang.org/appengine/aetest"
	"testing"
)

func TestGetCurrentExpoMissing(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	expo, err := GetCurrentExpo(c)
	if err != nil && expo == nil {
		t.Logf("Expo =%v, nil expected", expo)
		return
	}
	t.Errorf("Expected to get an error that the expo did not exist")
}
