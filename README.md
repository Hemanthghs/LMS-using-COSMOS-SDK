# Leave Management System using Cosmos SDK

This is a module developed based on Cosmos SDK for leave management.

## Requirements
- Golang
- Docker
- Protobuf
- Cobra-cli
- gRPC
- Cosmos SDK

## Methods Available
- Register Admin
- Add Students
- Apply Leave
- Approve Leave
- Get Leaves
- Get Leave Status
- Get Admins
- Get Students
- Get Approved Leaves

 ## Setup
  
**Install protocol compiler**
  
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
   
**Update PATH**
  
    export PATH="$PATH:$(go env GOPATH)/bin"

**Install Cosmos SDK**
    
    go get github.com/cosmos/cosmos-sdk

**Install Docker Engine**
Reference: https://docs.docker.com/engine/install/ubuntu/

**Install cobra-cli**
  
    go install github.com/spf13/cobra-cli@latest
    

## Client

### CLI

A user can query can interact with `lms` module using the CLI.

#### Query

The `query` commands allow users to query `leave` state.

```shell 
lmsd query leave --help
```

##### admins

The `admins` command allows users to query the list of existing admins.

```shell
lmsd query leave admins [flags]
```

Example:

```shell
lmsd query leave admins
```

Example Output:

```yml
admins:
- address: cosmos12fyxskc7cfjju80r8y2gk2d2a7hn7te78w53tz
  name: admin1
 ```

##### students

The `students` command allows users to query the list of existing students.

```shell
lmsd query leave students [flags]
```

Example:

```shell
lmsd query leave students
```

Example Output:

```yml
students:
- address: cosmos1cpr5lf7nx3h7fkz5fm0stehang7zeymv4wxpph
  id: "2"
  name: Sai
- address: cosmos1ez9zcdp5wa3hha8j7ysmehlzfa4hc3fvetycca
  id: "1"
  name: Hemanth
 ```
##### leaves

The `leaves` command allows users to query the list of all leaves applied by students.

```shell
lmsd query leave leaves [flags]
```

Example:

```shell
lmsd query leave leaves
```

Example Output:

```yml
leaverequests:
- address: cosmos1cpr5lf7nx3h7fkz5fm0stehang7zeymv4wxpph
  from: "2023-03-11T00:00:00Z"
  reason: Headache
  signer: cosmos12fyxskc7cfjju80r8y2gk2d2a7hn7te78w53tz
  status: STATUS_UNDEFINED
  to: "2023-03-13T00:00:00Z"
- address: cosmos1ez9zcdp5wa3hha8j7ysmehlzfa4hc3fvetycca
  from: "2023-03-19T00:00:00Z"
  reason: Fever
  signer: cosmos12fyxskc7cfjju80r8y2gk2d2a7hn7te78w53tz
  status: STATUS_UNDEFINED
  to: "2023-03-24T00:00:00Z"
 ```
##### approved

The `approved` command allows users to query the list of approved leaves.

```shell
lmsd query leave approved [flags]
```

Example:

```shell
lmsd query leave approved
```

Example Output:

```yml
leaverequests:
- admin: cosmos12fyxskc7cfjju80r8y2gk2d2a7hn7te78w53tz
  leave_id: cosmos1cpr5lf7nx3h7fkz5fm0stehang7zeymv4wxpph
  status: STATUS_ACCEPTED
 ```

##### status

The `status` command allows users to query the status of the leave applied.

```shell
lmsd query leave status [flags]
```

Example:

```shell
lmsd query leave status
```

Example Output:

```yml
leaverequest:
  address: cosmos1cpr5lf7nx3h7fkz5fm0stehang7zeymv4wxpph
  from: "2023-03-11T00:00:00Z"
  reason: Headache
  signer: cosmos12fyxskc7cfjju80r8y2gk2d2a7hn7te78w53tz
  status: STATUS_ACCEPTED
  to: "2023-03-13T00:00:00Z"
 ```

#### Transactions

The `tx` commands allow users to interact with `leave` module.

```shell 
lmsd tx leave --help
```

##### registeradmin

The `registeradmin` command allows users to register the new admins.

```shell
lmsd tx leave registeradmin [name] [address] [flags]
```

Example:

```shell
lmsd tx leave registeradmin admin1 cosmos12fyxskc7cfjju80r8y2gk2d2a7hn7te78w53tz --from validator-key --chain-id testnet
```

##### addstudents

The `addstudents` command allows admins to add new students.

```shell
lmsd tx leave addstudents [student name] [id] [address] [flags]
```

Example:

```shell
lmsd tx leave addstudents Hemanth 1 cosmos1ez9zcdp5wa3hha8j7ysmehlzfa4hc3fvetycca --from validator-key --chain-id testnet
```

##### applyleave

The `applyleave` command allows students to apply for a leave.

```shell
lmsd tx leave applyleave [Address] [Reason] [from] [to] [flags]
```

Example:

```shell
lmsd tx leave applyleave cosmos1ez9zcdp5wa3hha8j7ysmehlzfa4hc3fvetycca Fever 2023-Mar-01 2023-Mar-03 --from validator-key --chain-id testnet
```

##### acceptleave

The `acceptleave` command allows admins to approve a leave request.

```shell
lmsd tx leave acceptleave [LeaveId] [Status] [flags]
```

Example:

```shell
lmsd tx leave acceptleave cosmos1ez9zcdp5wa3hha8j7ysmehlzfa4hc3fvetycca 1 --from validator-key --chain-id testnet
```








