package dumper

import (
	"encoding/json"
	"fmt"
	"io"

	"lsif-jsonnet/protocol"
	"lsif-jsonnet/refs"
	"lsif-jsonnet/types"
)

// Dumps LSIF: https://github.com/Microsoft/language-server-protocol/blob/master/indexFormat/specification.md
type Dumper struct {
	projectRoot string
	w           io.Writer
	id          int
	encoder     *json.Encoder
	docIDByFile map[string]int
}

func NewDumper(projectRoot string, w io.Writer) *Dumper {
	encoder := json.NewEncoder(w)

	return &Dumper{
		projectRoot: projectRoot,
		w:           w,
		encoder:     encoder,
		docIDByFile: map[string]int{},
	}
}

func (dmpr *Dumper) DumpProject(info *protocol.ToolInfo, ll *refs.Listener) error {
	_, err := dmpr.emitMetaData("file://"+dmpr.projectRoot, info)
	if err != nil {
		return fmt.Errorf(`emit "metadata": %w`, err)
	}
	proID, err := dmpr.emitProject()
	if err != nil {
		return fmt.Errorf(`emit "project": %w`, err)
	}

	_, err = dmpr.emitBeginEvent("project", proID)
	if err != nil {
		return fmt.Errorf(`emit "begin": %w`, err)
	}

	err = dmpr.dumpListener(ll, proID)
	if err != nil {
		return err
	}

	_, err = dmpr.emitEndEvent("project", proID)
	return err
}

func (dmpr *Dumper) dumpListener(ll *refs.Listener, proID int) error {
	docID, err := dmpr.emitDocument(ll.File())
	if err != nil {
		return err
	}
	_, err = dmpr.emitBeginEvent("document", docID)
	if err != nil {
		return err
	}

	_, err = dmpr.emitContains(proID, []int{docID})
	if err != nil {
		return err
	}

	for _, dcl := range ll.Declarations() {
		err = dmpr.emitDeclaration(dcl, docID)
		if err != nil {
			return err
		}
	}

	for _, ill := range ll.Imports() {
		err = dmpr.dumpListener(ill, proID)
		if err != nil {
			return err
		}
	}

	_, err = dmpr.emitEndEvent("document", docID)
	return err
}

func (dmpr *Dumper) emitDeclaration(dcl *types.Declaration, docID int) error {
	startPos := protocol.Pos{
		Line:      dcl.StartPos.Line - 1,
		Character: dcl.StartPos.Column,
	}
	endPos := protocol.Pos{
		Line:      dcl.StartPos.Line - 1,
		Character: dcl.StartPos.Column + len(dcl.Name),
	}

	rangeID, err := dmpr.emitRange(startPos, endPos)
	if err != nil {
		return err
	}

	_, err = dmpr.emitContains(docID, []int{rangeID})
	if err != nil {
		return err
	}

	resultSetID, err := dmpr.emitResultSet()
	if err != nil {
		return err
	}

	_, err = dmpr.emitNext(rangeID, resultSetID)
	if err != nil {
		return err
	}

	defResultID, err := dmpr.emitDefinitionResult()
	if err != nil {
		return err
	}

	_, err = dmpr.emitTextDocumentDefinition(resultSetID, defResultID)
	if err != nil {
		return err
	}

	_, err = dmpr.emitItem(defResultID, []int{rangeID}, docID)
	if err != nil {
		return err
	}

	for _, use := range dcl.Uses {
		err = dmpr.emitUse(use, dcl, resultSetID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dmpr *Dumper) emitUse(use types.Use, dcl *types.Declaration, resultSetID int) error {
	startPos := protocol.Pos{
		Line:      use.StartPos.Line - 1,
		Character: use.StartPos.Column,
	}
	endPos := protocol.Pos{
		Line:      use.StartPos.Line - 1,
		Character: use.StartPos.Column + len(dcl.Name),
	}

	rangeID, err := dmpr.emitRange(startPos, endPos)
	if err != nil {
		return err
	}
	_, err = dmpr.emitNext(rangeID, resultSetID)
	if err != nil {
		return err
	}

	docID, ok := dmpr.docIDByFile[use.File]
	if !ok {
		return fmt.Errorf("document %s not declared", use.File)
	}

	_, err = dmpr.emitContains(docID, []int{rangeID})
	return err
}

var newLineBytes = []byte("\n")

func (dmpr *Dumper) writeNewLine() error {
	_, err := dmpr.w.Write(newLineBytes)
	return err
}

func (dmpr *Dumper) nextID() int {
	dmpr.id++
	return dmpr.id
}

func (dmpr *Dumper) emit(v interface{}) error {
	return dmpr.encoder.Encode(v)
}

func (dmpr *Dumper) emitMetaData(root string, info *protocol.ToolInfo) (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewMetaData(id, root, info))
}

func (dmpr *Dumper) emitBeginEvent(scope string, data int) (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewEvent(id, "begin", scope, data))
}

func (dmpr *Dumper) emitEndEvent(scope string, data int) (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewEvent(id, "end", scope, data))
}

func (dmpr *Dumper) emitProject() (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewProject(id))
}

func (dmpr *Dumper) emitDocument(path string) (int, error) {
	id, ok := dmpr.docIDByFile[path]
	if !ok {
		id = dmpr.nextID()
		dmpr.docIDByFile[path] = id
	}
	return id, dmpr.emit(protocol.NewDocument(id, "file://"+path, nil))
}

func (dmpr *Dumper) emitContains(outV int, inVs []int) (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewContains(id, outV, inVs))
}

func (dmpr *Dumper) emitResultSet() (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewResultSet(id))
}

func (dmpr *Dumper) emitRange(start, end protocol.Pos) (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewRange(id, start, end))
}

func (dmpr *Dumper) emitNext(outV, inV int) (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewNext(id, outV, inV))
}

func (dmpr *Dumper) emitDefinitionResult() (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewDefinitionResult(id))
}

func (dmpr *Dumper) emitTextDocumentDefinition(outV, inV int) (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewTextDocumentDefinition(id, outV, inV))
}

func (dmpr *Dumper) emitItem(outV int, inVs []int, docID int) (int, error) {
	id := dmpr.nextID()
	return id, dmpr.emit(protocol.NewItem(id, outV, inVs, docID))
}
