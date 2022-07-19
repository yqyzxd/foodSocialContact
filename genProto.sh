function genProto(){
    DOMAIN=$1
    PROTO=$2
    SKIP_GATEWAY=$3
    PROTO_PATH=./$DOMAIN/proto
    GO_OUT_PATH=./$DOMAIN/proto/gen
    mkdir -p "$GO_OUT_PATH"

  if [ $SKIP_GATEWAY ]; then

    protoc  -I=$PROTO_PATH \
            --go_out $GO_OUT_PATH --go_opt paths=source_relative \
            --go-grpc_out $GO_OUT_PATH --go-grpc_opt paths=source_relative \
            "$PROTO".proto
  else 

    protoc  -I=$PROTO_PATH \
            --go_out $GO_OUT_PATH --go_opt paths=source_relative \
            --go-grpc_out $GO_OUT_PATH --go-grpc_opt paths=source_relative \
            --grpc-gateway_out $GO_OUT_PATH --grpc-gateway_opt paths=source_relative,grpc_api_configuration=$PROTO_PATH/${DOMAIN}.yaml \
            "$PROTO".proto
  fi
}
genProto ms-oauth oauth 1
genProto ms-diner diner 1
genProto ms-follow follow 1
