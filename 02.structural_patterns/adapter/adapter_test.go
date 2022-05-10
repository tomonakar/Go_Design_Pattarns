package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello"

	adapter := PrinterAdapter {OldPrinter: &MyLegacyPrinter{}, Msg: msg}

	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Hello\n" {
		t.Errorf("Messag didn't match: %s", returnedMsg)
	}

}