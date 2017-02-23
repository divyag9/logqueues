package logger

import "testing"

var logTests = []struct {
	data []byte
	file string
	err  error
}{
	{
		[]byte(`[
    			{"ID":"5893ecc490df111978809d15","Name":"spi","Type":"spi","Depth":1000,"Rate":10,"LastProcessed":"2008-09-16T19:00:00-05:00","LastReported":"2017-02-02T20:52:42.471-06:00"},
				{"ID":"5893f0dbd8ef364087be110b","Name":"spi1","Type":"spi1","Depth":10000,"Rate":100,"LastProcessed":"2009-09-16T19:00:00-05:00","LastReported":"2017-02-02T20:55:12.27-06:00"}
    		]`),
		"logqueue",
		nil,
	},
}

func TestLogQueueDataToFile(t *testing.T) {
	for _, logTest := range logTests {
		err := LogQueueDataToFile(logTest.data, logTest.file)
		if err != logTest.err && err.Error() != logTest.err.Error() {
			t.Errorf("Returned: %v. Expected: %v", err, logTest.err)
		}
	}
}
