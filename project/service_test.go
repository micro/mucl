package project

import (
	"testing"

	require "github.com/alecthomas/assert/v2"
	"github.com/micro/mu/mucl"
)

func TestGoodService(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, "helloworld", svc.Name)
}

func TestStreamService(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodStreamMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, "hellostream", svc.Name)
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

var goodStreamMucl = `
service hellostream {
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
  rpc Search(SearchRequest) returns (stream SearchResponse)
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
  enum SearchType {
    SHALLOW = 0
    DEEP = 1
  }
  type SearchType
  query string
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

var embeddedTypeMucl = `
service helloworld {
  broker http
  transport grpc
  registry etcd
  server mucp
}

type SearchRequest {
  enum SearchType {
    SHALLOW = 0
    DEEP = 1
  }
  type SearchType
  query string
  page_number int32
  result_per_page int32

}

type SearchResponse {
  type PaginationResponse {
    page_number int32
    page_count int32
  }
  repeated results string
  pagination PaginationResponse
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`
var optionMucl = `
service helloworld {
  broker http
  transport grpc
  registry etcd
  server mucp
}

type SearchRequest {
  option struct_tags = 32;
  query string
  type SearchType
  page_number int32
  result_per_page int32
}

type SearchResponse {
  option (micro).tags = "json";
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
