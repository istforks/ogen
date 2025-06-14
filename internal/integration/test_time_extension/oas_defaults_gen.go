// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/jx"

	"github.com/istforks/ogen/json"
)

// setDefaults set default value of fields.
func (s *DefaultOK) setDefaults() {
	{
		val, _ := json.DecodeTimeFormat(jx.DecodeStr("\"04/03/2001\""), "02/01/2006")
		s.Date.SetTo(val)
	}
	{
		val, _ := json.DecodeTimeFormat(jx.DecodeStr("\"1:23AM\""), "3:04PM")
		s.Time.SetTo(val)
	}
	{
		val, _ := json.DecodeTimeFormat(jx.DecodeStr("\"2001-03-04T01:23:45.123456789-07:00\""), "2006-01-02T15:04:05.999999999Z07:00")
		s.DateTime.SetTo(val)
	}
}
