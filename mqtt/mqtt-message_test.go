package mqtt

import (
	"reflect"
	"testing"
)

func TestParseMqttMessageSuccess(t *testing.T) {
	m, err := ParseMqttMessage("{\"header\": {\"message_id\":\"123\"}, \"body\": {\"light\":\"on\"}}")
	if err != nil {
		t.Errorf("Expected to parse json successfully")
	}
	if err != nil {
		t.Errorf("Expected an error for invalid Json")
	}
	if m.Header.MessageId != "123" {
		t.Errorf("Message Id should be ´123´ got %s", m.Header.MessageId)
	}
	body := map[string]string{"light": "on"}
	if !reflect.DeepEqual(m.Body, body) {
		t.Errorf("Expected body ´light´:`on` got %s", m.Body)
	}
}
