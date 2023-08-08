package torrent

import (
	"errors"
	"fmt"
	"io"

	"github.com/anacrolix/torrent/metainfo"
)

// The info dictionary.
type Info struct {
	PieceLength int64  `bencode:"piece length"` // BEP3
	Pieces      []byte `bencode:"pieces"`       // BEP3
	Name        string `bencode:"name"`         // BEP3
	NameUtf8    string `bencode:"name.utf-8,omitempty"`
	Length      int64  `bencode:"length,omitempty"`  // BEP3, mutually exclusive with Files
	Private     *bool  `bencode:"private,omitempty"` // BEP27
	// TODO: Document this field.
	Source string              `bencode:"source,omitempty"`
	Files  []metainfo.FileInfo `bencode:"files,omitempty"` // BEP3, mutually exclusive with Length
}

func (info *Info) writeFiles(w io.Writer, open func(fi metainfo.FileInfo) (io.ReadCloser, error)) error {
	for _, fi := range info.UpvertedFiles() {
		r, err := open(fi)
		if err != nil {
			return fmt.Errorf("error opening %v: %s", fi, err)
		}
		wn, err := io.CopyN(w, r, fi.Length)
		r.Close()
		if wn != fi.Length {
			return fmt.Errorf("error copying %v: %s", fi, err)
		}
	}
	return nil
}

func (info *Info) GeneratePieces(open func(fi metainfo.FileInfo) (io.ReadCloser, error)) (err error) {
	if info.PieceLength == 0 {
		return errors.New("piece length must be non-zero")
	}
	pr, pw := io.Pipe()
	go func() {
		err := info.writeFiles(pw, open)
		pw.CloseWithError(err)
	}()
	defer pr.Close()
	info.Pieces, err = metainfo.GeneratePieces(pr, info.PieceLength, nil)
	return
}

func (info *Info) TotalLength() (ret int64) {
	return info.Length
}

func (info *Info) UpvertedFiles() []metainfo.FileInfo {
	if len(info.Files) == 0 {
		return []metainfo.FileInfo{{
			Length: info.Length,
			// Callers should determine that Info.Name is the basename, and
			// thus a regular file.
			Path: nil,
		}}
	}
	return info.Files
}

// This is a helper that sets Files and Pieces from a root path and its children.
func (info *Info) BuildFromReader(r io.ReadCloser) (err error) {
	info.Files = nil

	if info.PieceLength == 0 {
		info.PieceLength = metainfo.ChoosePieceLength(info.TotalLength())
	}
	err = info.GeneratePieces(func(fi metainfo.FileInfo) (io.ReadCloser, error) { return r, nil })
	if err != nil {
		err = fmt.Errorf("error generating pieces: %s", err)
	}
	return
}
