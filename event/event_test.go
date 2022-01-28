package event

import "testing"

func TestViewEvent(t *testing.T) {
	wanted := "[Kylling Rev Korn HS ---V \\____________/ _____________________Ø---]"
	event := ViewEvent()
	if event != wanted {
		t.Errorf("Feil, vi fikk %q, ønsket %q", event, wanted)
	}
}

func TestSettInnKylling(t *testing.T) {
	wanted := event
	if event != wanted {
		t.Errorf("Feil, vi fikk %q, ønsket %q", event, wanted)
	}
}
