# distributed-key-value-store

This is a learning project that consists of an implementation of a distributed key-value store using the Go programming language  and is meant to solidify my knowledge in the following topics :
- Distributed systems patterns 
- Distributed consenus
- Replication based consensus (Raft or Paxos)
- Storage engines data structures, this project is a LSM-tree based implementation based
- Go programming language 
- gRPC protocol and protocol buffers (Protobuf)

## The Goal 

The goal is to implement a simple but fully functional distributed key-value database that follows the ACID properties.

The functionalaties will include :
- Adding key/value
- Updating key/value
- Getting all keys and their values 
- Getting value by key
- Deleting a particular key


## Implementation Plan
1. Writing an in-memory single server key-value store using the gRPC protocol
2. Adding authentication using JWT 
3. Implementing a write-ahead-log for data backup to disk 
4. Implementing the SSTable and the memtable  
5. Implementing a Raft Consenus-Based algorithm for supporting replication of data across multiple storage nodes 
6. Adding a tracing framework