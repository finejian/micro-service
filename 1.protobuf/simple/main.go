// Copyright © 2018 FineJian <fanjian.j@qq.com>
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"log"

	"github.com/finejian/micro-service/1.protobuf/simple/pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	hello := &pb.Hello{
		Message: "Micor-Service",
	}
	// 编码
	data, err := proto.Marshal(hello)
	if err != nil {
		log.Fatal("marshal error: ", err)
	}

	// 解码
	newHello := &pb.Hello{}
	err = proto.Unmarshal(data, newHello)
	if err != nil {
		log.Fatal("unmarshal error: ", err)
	}

	// 测试结果
	log.Printf("hello %s", hello.String())
	log.Printf("new hello %s", newHello.String())
	log.Printf("same hello: %v", hello.GetMessage() == newHello.GetMessage())
}
