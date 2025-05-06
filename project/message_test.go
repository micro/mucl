package project

import (
	"testing"

	require "github.com/alecthomas/assert/v2"
	"github.com/micro/mu/mucl"
)

func TestMessages(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 2, len(svc.MessageMap))
}

func TestEmbeddedEnumMessages(t *testing.T) {
	def, err := mucl.Parser.ParseString("", embeddedEnumMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 2, len(svc.MessageMap))
	searchType, ok := svc.GetMessage("SearchRequest")
	require.True(t, ok)
	require.Equal(t, "SearchRequest", searchType.Name)
	require.Equal(t, 4, len(searchType.GetAllFields()))
}

func TestEmbeddedMsgMessages(t *testing.T) {
	def, err := mucl.Parser.ParseString("", embeddedTypeMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 3, len(svc.MessageMap))
	responseType, ok := svc.GetMessage("SearchResponse")
	require.True(t, ok)
	require.Equal(t, "SearchResponse", responseType.Name)
	require.Equal(t, 2, len(responseType.GetAllFields()))
}

func TestOptionsMsg(t *testing.T) {
	def, err := mucl.Parser.ParseString("", optionMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 2, len(svc.MessageMap))
	searchType, ok := svc.GetMessage("SearchRequest")
	require.True(t, ok)
	require.Equal(t, "SearchRequest", searchType.Name)
	require.Equal(t, 4, len(searchType.GetAllFields()))

	require.Equal(t, 1, len(searchType.Options))
	// check if the options are correct
	options := searchType.Options
	o, ok := options.Get("struct_tags")
	require.True(t, ok)
	require.Equal(t, "struct_tags", o.Name)
	sval, ok := o.Value.GetValue().(int64)
	require.True(t, ok)
	require.Equal(t, 32, sval)
}
