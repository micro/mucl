project "mymod"
service users {}

type CreateRequest {
  input string
}

type CreateResponse {
  output string
}

server User {
  rpc Create(CreateRequest) returns (CreateResponse)
}
