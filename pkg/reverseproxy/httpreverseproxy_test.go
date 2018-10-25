package reverseproxy

import (
	"errors"
	"github.com/function61/gokit/assert"
	"testing"
)

func TestDestinationPortFromVirtualHost(t *testing.T) {
	testCase := func(input string, destinationPortExpected int, errExpected error) {
		t.Helper()

		destinationPort, err := destinationPortFromVirtualHost(input)

		if errExpected == nil {
			assert.True(t, err == errExpected)
		} else {
			assert.EqualString(t, err.Error(), errExpected.Error())
		}

		assert.True(t, destinationPort == destinationPortExpected)

	}

	testCase("8081.punch.fn61.net", 8081, nil)
	testCase("4431.com", 4431, nil)
	testCase("4431.com:80", 4431, nil)

	testCase("80.punch.fn61.net", 0, errors.New("destination port is disallowed"))
	testCase("4431", 0, errors.New("failed to determine destination port from vhost"))
	testCase("4431:80", 0, errors.New("failed to determine destination port from vhost"))
}
