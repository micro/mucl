package mucl

import (
	"testing"

	require "github.com/alecthomas/assert/v2"
)

func TestGood(t *testing.T) {
	tree, err := Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	require.Equal(t, 4, len(tree.Entries))
	require.Equal(t, "helloworld", tree.Service.Name)

	require.Equal(t, "http", tree.Service.Entries[0].Broker.Name)
	require.Equal(t, "SearchRequest", tree.Entries[0].Message.Name)
	require.Equal(t, "SearchResponse", tree.Entries[1].Message.Name)
	require.Equal(t, "SearchType", tree.Entries[2].Enum.Name)
	require.Equal(t, "SearchService", tree.Entries[3].Endpoint.Name)
}

func TestEmbeddedGood(t *testing.T) {
	tree, err := Parser.ParseString("", embeddedEnumMucl)
	require.NoError(t, err)
	require.Equal(t, 3, len(tree.Entries))
	require.Equal(t, "helloworld", tree.Service.Name)

	require.Equal(t, "http", tree.Service.Entries[0].Broker.Name)
	require.Equal(t, "SearchRequest", tree.Entries[0].Message.Name)
	require.Equal(t, "SearchResponse", tree.Entries[1].Message.Name)
	require.Equal(t, "SearchService", tree.Entries[2].Endpoint.Name)
}

func TestGoodTwoEndpoints(t *testing.T) {
	tree, err := Parser.ParseString("", twoEndpoints)
	require.NoError(t, err)
	require.Equal(t, 5, len(tree.Entries))
	require.Equal(t, "helloworld", tree.Service.Name)

	require.Equal(t, "http", tree.Service.Entries[0].Broker.Name)
	require.Equal(t, "SearchRequest", tree.Entries[0].Message.Name)
	require.Equal(t, "SearchResponse", tree.Entries[1].Message.Name)
	require.Equal(t, "SearchType", tree.Entries[2].Enum.Name)

	require.Equal(t, "SearchService", tree.Entries[3].Endpoint.Name)
	require.Equal(t, "InternalSearchService", tree.Entries[4].Endpoint.Name)
}

func TestWithTwoServices(t *testing.T) {
	_, err := Parser.ParseString("", twoServices)
	require.Error(t, err)
}

var goodMucl = `
service helloworld {
  broker http
  transport grpc
  registry etcd
  server mucp
}

type SearchRequest {
  query string
  type SearchType
  page_number int32
  result_per_page int32
}

type SearchResponse {
  results string
}

enum SearchType {
  SHALLOW = 0
  DEEP = 1
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`

var twoServices = `
service helloworld {
  broker http
  transport grpc
  registry etcd
  server mucp
}
service helloworld2 {}

type SearchRequest {
  query string
  type SearchType
  page_number int32
  result_per_page int32
}

type SearchResponse {
  results string
}

enum SearchType {
  SHALLOW = 0
  DEEP = 1
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`

var twoEndpoints = `
service helloworld {
  broker http
  transport grpc
  registry etcd
  server mucp
}

type SearchRequest {
  query string
  type SearchType
  page_number int32
  result_per_page int32
}

type SearchResponse {
  results string
}

enum SearchType {
  SHALLOW = 0
  DEEP = 1
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}

endpoint InternalSearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`

var embeddedEnumMucl = `
service helloworld {
  broker http
  transport grpc
  registry etcd
  server mucp
}

type SearchRequest {
  query string
  enum SearchType {
    SHALLOW = 0
    DEEP = 1
  }
  page_number int32
  result_per_page int32
}

type SearchResponse {
  results string
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`
