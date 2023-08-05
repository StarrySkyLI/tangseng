package index

import (
	"fmt"
	"testing"

	"github.com/CocaineCong/tangseng/app/search-engine/logic/query"
	"github.com/CocaineCong/tangseng/config"
	log "github.com/CocaineCong/tangseng/pkg/logger"
)

func TestMain(m *testing.M) {
	// 这个文件相对于config.yaml的位置
	re := config.ConfigReader{FileName: "../../../../config/config.yaml"}
	config.InitConfigForTest(&re)
	log.InitLog()
	query.InitSeg()
	fmt.Println("Write tests on values: ", config.Conf)
	m.Run()
}

func TestIndexRunning(t *testing.T) {
	RunningIndex()
}