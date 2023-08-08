package torrent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/anacrolix/torrent/metainfo"
)

func Check(ctx context.Context, filePath string) (string, error) {
	f, _ := os.Open(filePath)
	mInfo, err := metainfo.Load(f)
	if err != nil {
		fmt.Println(ctx, err)
		return "", errors.New("服务错误")
	}
	info, err := mInfo.UnmarshalInfo()
	if err != nil {
		fmt.Println(ctx, err)
		return "", errors.New("服务错误")
	}
	// 分片数量
	pieceCount := int(math.Ceil(float64(info.Length) / float64(info.PieceLength)))
	bf := new(strings.Builder)
	for i := 0; i < pieceCount; i++ {
		h := info.Piece(i).Hash().String()
		fmt.Println(h)
		bf.WriteString(h)
	}
	return bf.String(), nil
}
