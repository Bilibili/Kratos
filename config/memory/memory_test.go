package memory

import (
	"errors"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/config"
)

const (
	_testJSON = `
{
  "test": {
    "settings" : {
      "int_key": 1000,
      "float_key": 1000.1,
      "duration_key": 10000,
      "string_key": "string_value"
    },
    "server": {
      "addr": "127.0.0.1",
      "port": 8000
    }
  }
}`
)

func TestMemory(t *testing.T) {
	var (
		data = []byte(_testJSON)
	)
	testSource(t, data)
}

func testSource(t *testing.T, data []byte) {
	s := NewSource(WithJSON(data))
	kvs, err := s.Load()
	if err != nil {
		t.Error(err)
	}
	if string(kvs[0].Value) != string(data) {
		t.Errorf("no expected: %s, but got: %s", kvs[0].Value, data)
	}
}

func TestConfig(t *testing.T) {
	var (
		data = []byte(_testJSON)
	)

	c := config.New(
		config.WithSource(
			NewSource(
				WithJSON(data),
			),
		),
	)

	testConfig(t, c)
}

func testConfig(t *testing.T, c config.Config) {
	var expected = map[string]interface{}{
		"test.settings.int_key":      int64(1000),
		"test.settings.float_key":    float64(1000.1),
		"test.settings.string_key":   "string_value",
		"test.settings.duration_key": time.Duration(10000),
		"test.server.addr":           "127.0.0.1",
		"test.server.port":           int64(8000),
	}
	if err := c.Load(); err != nil {
		t.Error(err)
	}
	for key, value := range expected {
		switch value.(type) {
		case int64:
			if v, err := c.Value(key).Int(); err != nil {
				t.Error(key, value, err)
			} else if v != value {
				t.Errorf("no expect key: %s value: %v, but got: %v", key, value, v)
			}
		case float64:
			if v, err := c.Value(key).Float(); err != nil {
				t.Error(key, value, err)
			} else if v != value {
				t.Errorf("no expect key: %s value: %v, but got: %v", key, value, v)
			}
		case string:
			if v, err := c.Value(key).String(); err != nil {
				t.Error(key, value, err)
			} else if v != value {
				t.Errorf("no expect key: %s value: %v, but got: %v", key, value, v)
			}
		case time.Duration:
			if v, err := c.Value(key).Duration(); err != nil {
				t.Error(key, value, err)
			} else if v != value {
				t.Errorf("no expect key: %s value: %v, but got: %v", key, value, v)
			}
		}
	}
	// scan
	var settings struct {
		IntKey      int64         `json:"int_key"`
		FloatKey    float64       `json:"float_key"`
		StringKey   string        `json:"string_key"`
		DurationKey time.Duration `json:"duration_key"`
	}
	if err := c.Value("test.settings").Scan(&settings); err != nil {
		t.Error(err)
	}
	if v := expected["test.settings.int_key"]; settings.IntKey != v {
		t.Errorf("no expect int_key value: %v, but got: %v", settings.IntKey, v)
	}
	if v := expected["test.settings.float_key"]; settings.FloatKey != v {
		t.Errorf("no expect float_key value: %v, but got: %v", settings.FloatKey, v)
	}
	if v := expected["test.settings.string_key"]; settings.StringKey != v {
		t.Errorf("no expect string_key value: %v, but got: %v", settings.StringKey, v)
	}
	if v := expected["test.settings.duration_key"]; settings.DurationKey != v {
		t.Errorf("no expect duration_key value: %v, but got: %v", settings.DurationKey, v)
	}

	// not found
	if _, err := c.Value("not_found_key").Bool(); errors.Is(err, config.ErrNotFound) {
		t.Logf("not_found_key not match: %v", err)
	}

}
