package sdk

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getAPIMaxTimeout(t *testing.T) {
	timeout, ok := getAPIMaxTimeout("ecs", "UnassociateEipAddress")
	assert.True(t, ok)
	assert.Equal(t, 16*time.Second, timeout)

	timeout, ok = getAPIMaxTimeout("Ecs", "UnassociateEipAddress")
	assert.True(t, ok)
	assert.Equal(t, 16*time.Second, timeout)

	timeout, ok = getAPIMaxTimeout("Acs", "UnassociateEipAddress")
	assert.False(t, ok)
	assert.Equal(t, 0*time.Second, timeout)

	tmp := apiTimeouts
	apiTimeouts = ""
	timeout, ok = getAPIMaxTimeout("Acs", "UnassociateEipAddress")
	apiTimeouts = tmp
	assert.False(t, ok)
	assert.Equal(t, 0*time.Second, timeout)
}

func TestGetAPIMaxTimeout(t *testing.T) {
	for _, tc := range []struct {
		product, action string
		want            time.Duration
		found           bool
	}{
		{"Ecs", "CreateInstance", 86 * time.Second, true},
		{"ecs", "DescribeInstances", 10 * time.Second, true},
		{"rds", "CreateDBInstance", 0, false},
		{"ecs", "UnknownAPI", 0, false},
	} {
		got, found := GetAPIMaxTimeout(tc.product, tc.action)
		if got != tc.want || found != tc.found {
			t.Errorf("%s/%s: got (%v, %v), want (%v, %v)", tc.product, tc.action, got, found, tc.want, tc.found)
		}
	}
}
