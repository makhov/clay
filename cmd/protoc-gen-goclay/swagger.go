package main

import (
	"github.com/golang/glog"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"github.com/utrack/grpc-gateway/protoc-gen-swagger/genswagger"
)

func genSwaggerDef(req *plugin.CodeGeneratorRequest, pkgMap map[string]string) ([]byte, error) {
	reg := descriptor.NewRegistry()
	reg.SetPrefix(*importPrefix)
	reg.SetAllowDeleteBody(*allowDeleteBody)

	for k, v := range pkgMap {
		reg.AddPkgMap(k, v)
	}

	gsw := genswagger.New(reg)

	if err := reg.Load(req); err != nil {
		return nil, err
	}

	var targets []*descriptor.File
	for _, target := range req.FileToGenerate {
		f, err := reg.LookupFile(target)
		if err != nil {
			glog.Fatal(err)
		}
		targets = append(targets, f)
	}

	outSwag, err := gsw.Generate(targets)
	if err != nil {
		return nil, err
	}
	return []byte(outSwag[0].GetContent()), nil
}