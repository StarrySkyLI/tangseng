package storage

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func TestMmap(t *testing.T) {
	filePath := "../../data/db/0.forward"
	fd, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer fd.Close()

	// 获取文件大小
	fi, err := fd.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败:", err)
		return
	}
	fileSize := fi.Size()
	fmt.Println("fileSize", fileSize)
	// 设置要映射的偏移量，假设从中间开始映射
	offset := int64(10)

	// 获取从偏移量开始到结尾的长度
	length := int(fileSize - offset)

	// 映射整个文件到内存
	mmapData, err := Mmap(int(fd.Fd()), offset, length)
	if err != nil {
		fmt.Println("映射文件失败:", err)
		return
	}
	defer syscall.Munmap(mmapData)

	// 使用 mmapData 可以直接读取文件内容
	fmt.Printf("文件内容：%s\n", string(mmapData))
}
