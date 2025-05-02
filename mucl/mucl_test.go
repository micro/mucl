package mucl

import (
	"testing"

	require "github.com/alecthomas/assert/v2"
)

func TestExe(t *testing.T) {
	tree, err := Parser.ParseString("", `
project "github.com/test/test"
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

server SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`)
	require.NoError(t, err)
	require.Equal(t, 6, len(tree.Entries))
	require.Equal(t, "\"github.com/test/test\"", tree.Entries[0].Project)
	require.Equal(t, "helloworld", tree.Entries[1].Service.Name)

	require.Equal(t, "http", tree.Entries[1].Service.Entries[0].Broker.Name)
	require.Equal(t, "SearchRequest", tree.Entries[2].Message.Name)
	require.Equal(t, "SearchResponse", tree.Entries[3].Message.Name)
	require.Equal(t, "SearchType", tree.Entries[4].Enum.Name)
}
