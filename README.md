# ☕ gRPC Coffee Shop

> A minimal Go gRPC demo — stream a menu, place an order, check your status.

---

## Overview

**CoffeeShop** is a simple gRPC service written in Go that demonstrates three core gRPC patterns through a familiar café metaphor. It's a clean starting point for anyone learning gRPC, Protocol Buffers, and Go service design.

```
Client ──► GetMenu()        →  Server-side streaming
Client ──► PlaceOrder()     →  Unary RPC
Client ──► GetOrderStatus() →  Unary RPC
```

---

## Project Structure

```
proto_example/
├── coffee_shop.proto        # Protobuf service definition
├── coffeeshop_proto/        # Generated Go code (pb + grpc)
├── server.go                # gRPC server implementation
├── client/
│   └── client.go            # gRPC client
└── Makefile
```

---

## Service Definition

The service is defined in `coffee_shop.proto`:

```protobuf
service CoffeeShop {
    rpc GetMenu(MenuRequest)         returns (stream Menu)  {}
    rpc PlaceOrder(Order)            returns (Receipt)      {}
    rpc GetOrderStatus(Receipt)      returns (OrderStatus)  {}
}
```

| Method | Type | Description |
|---|---|---|
| `GetMenu` | Server-side streaming | Streams menu items one by one |
| `PlaceOrder` | Unary | Submits an order, returns a receipt ID |
| `GetOrderStatus` | Unary | Returns the current status of an order |

---

## Getting Started

### Prerequisites

- Go 1.21+
- `protoc` — [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Generate Protobuf Code

```bash
make build_proto
```

### Run the Server

```bash
go run server.go
# Listening on :9001
```

### Run the Client

```bash
go run client/client.go
```

#### Example Output

```
Resp received: [id:"1" name:"Black Coffee"]
Resp received: [id:"1" name:"Black Coffee" id:"2" name:"Cappuccino"]
Resp received: [id:"1" name:"Black Coffee" id:"2" name:"Cappuccino" id:"3" name:"Latte"]
receipt: id:"ABC123"
status: orderId:"ABC123" status:"IN PROGRESS"
```

---

## How It Works

1. **GetMenu** — The server streams the menu incrementally. Each message adds one more item, simulating a real-time data feed.
2. **PlaceOrder** — The client collects all streamed items, then places an order with the full menu selection.
3. **GetOrderStatus** — The client polls the order status using the receipt ID returned from `PlaceOrder`.

---

## Tech Stack

- **Go** — Server & client implementation
- **gRPC** — Remote procedure call framework
- **Protocol Buffers (proto3)** — Interface definition language
- **`google.golang.org/grpc`** — Official Go gRPC package
