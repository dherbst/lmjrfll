package lmjrfll

import (
	"google.golang.org/appengine/aetest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApi1UserProfileHandler(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Errorf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	request, err := inst.NewRequest("GET", "/api/1/user/profile", nil)
	if err != nil {
		t.Errorf("Error creating new instance %v", err)
	}

	response := httptest.NewRecorder()
	Api1UserProfileHandler(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Did not get StatusOK got %v", response.Code)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading body %v", err)
	}
	prefix := `{"Name":"","Email":"","LogoutUrl":"","LoginUrl":"/_ah/login?continue=`
	if !strings.HasPrefix(string(data), prefix) {
		t.Errorf("Expected  %s got %v", prefix, string(data))
	}

}

func TestApi1ExpoCurrentHandler(t *testing.T) {

}
