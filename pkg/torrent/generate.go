package torrent

import (
	"bytes"
	"fmt"
	"io"

	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"
)

var builtinAnnounceList = [][]string{
	{"http://xx.xx.com:1337/announce"},
	{"udp://xx.xx.com:2800/announce"},
	{"udp://xx.xx.com:2800/announce"},
	{"udp://xx.xx.com:2800/announce"},
	{"udp://xx.xx.com:2800/announce"},
}

func GenerateTorrentBase64(reader io.ReadCloser, length int64, name string) (r string, err error) {
	mi := metainfo.MetaInfo{
		AnnounceList: builtinAnnounceList,
	}
	mi.SetDefaults()

	info := Info{
		PieceLength: 10 * 1024 * 1024,
	}
	info.Length = length
	if err = info.BuildFromReader(reader); err != nil {
		fmt.Println("错误", err)
		return "", err
	}
	info.Name = name

	mi.InfoBytes, err = bencode.Marshal(info)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = mi.Write(&buf); err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}
