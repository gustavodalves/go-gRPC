# gRPC Communication Patterns in Go

## Communication Patterns

### Unary Communication

Unary RPC is the simplest form of gRPC communication. In this pattern, the client sends a single request to the server and receives a single response in return.

### Server-side Streaming

Server-side streaming RPC allows the server to send a stream of responses to a single request from the client. The client sends a unary request, and the server responds with a stream of messages.

### Client-side Streaming

Client-side streaming RPC allows the client to send a stream of requests to the server. The server processes these requests and sends back a single response.

### Bidirectional Streaming

Bidirectional streaming RPC enables both the client and server to send a stream of messages to each other. This allows for real-time communication between the two parties.
