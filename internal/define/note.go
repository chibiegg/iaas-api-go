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

package define

import (
	"github.com/sacloud/iaas-api-go/internal/define/names"
	"github.com/sacloud/iaas-api-go/internal/define/ops"
	"github.com/sacloud/iaas-api-go/internal/dsl"
	"github.com/sacloud/iaas-api-go/internal/dsl/meta"
	"github.com/sacloud/iaas-api-go/naked"
)

const (
	noteAPIName     = "Note"
	noteAPIPathName = "note"
)

var noteAPI = &dsl.Resource{
	Name:       noteAPIName,
	PathName:   noteAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.Find(noteAPIName, noteNakedType, findParameter, noteView),

		// create
		ops.Create(noteAPIName, noteNakedType, noteCreateParam, noteView),

		// read
		ops.Read(noteAPIName, noteNakedType, noteView),

		// update
		ops.Update(noteAPIName, noteNakedType, noteUpdateParam, noteView),

		// delete
		ops.Delete(noteAPIName),
	},
}

var (
	noteNakedType = meta.Static(naked.Note{})

	noteView = &dsl.Model{
		Name:      noteAPIName,
		NakedType: noteNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Scope(),
			fields.NoteClass(),
			fields.NoteContent(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	noteCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(noteAPIName),
		NakedType: noteNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}

	noteUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(noteAPIName),
		NakedType: noteNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}
)