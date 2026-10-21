protoc.exe --go_out=./rpc ./rpcproto/*.proto
protoc.exe --go_out=./msg ./msgproto/*.proto
protoc.exe --go_out=../db -I ../db/ ../db/*.proto

protoc.exe --go_out=./notify ./notify_proto/*.proto



PAUSE