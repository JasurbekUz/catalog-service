#!/bin/bash

CURRENT_DIR=$(pwd)

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/book-shop_protos/catalog_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/catalog_service/ \
        $CURRENT_DIR/book-shop_protos/catalog_service/*.proto;

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/book-shop_protos/order_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/order_service/ \
        $CURRENT_DIR/book-shop_protos/order_service/*.proto;        

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i "" -e "s/,omitempty//g" $CURRENT_DIR/genproto/catalog_service/*.go
  else
    sed -i -e "s/,omitempty//g" $CURRENT_DIR/genproto/catalog_service/*.go
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i "" -e "s/,omitempty//g" $CURRENT_DIR/genproto/order_service/*.go
  else
    sed -i -e "s/,omitempty//g" $CURRENT_DIR/genproto/order_service/*.go
fi