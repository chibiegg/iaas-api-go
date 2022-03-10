// Copyright 2022 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/iaas-api-go/internal/define"
	"github.com/sacloud/iaas-api-go/internal/dsl"
	"github.com/sacloud/iaas-api-go/internal/tools"
)

const destination = "fake/zz_store.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-fake-store: ")
}

func main() {
	dsl.IsOutOfSacloudPackage = true

	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), destination),
		Template:   tmpl,
		Parameter:  define.APIs,
	})
	log.Printf("generated: %s\n", filepath.Join(destination))
}

const tmpl = `// generated by 'github.com/sacloud/iaas-api-go/internal/tools/gen-fake-store'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-api-go/accessor"
)

{{ range . }} 
func get{{.TypeName}}(zone string) []*iaas.{{.TypeName}} {
	values := ds().List(Resource{{.TypeName}}, zone)
	var ret []*iaas.{{.TypeName}}
	for _ , v := range values {
		if v, ok := v.(*iaas.{{.TypeName}}); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func get{{.TypeName}}ByID(zone string, id types.ID) *iaas.{{.TypeName}} {
	v := ds().Get(Resource{{.TypeName}}, zone, id)
	if v, ok := v.(*iaas.{{.TypeName}}); ok {
		return v
	}
	return nil
}

func put{{.TypeName}}(zone string, value *iaas.{{.TypeName}}) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(Resource{{.TypeName}}, zone, id.GetID(), value)
		return
	}
	ds().Put(Resource{{.TypeName}}, zone, 0, value)
}
{{ end }}
`