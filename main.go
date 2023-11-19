package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("カレントディレクトリの取得に失敗しました: %v\n", err)
		return
	}
	dir := filepath.Dir(exePath)
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("ディレクトリの読み込みに失敗しました: %v", err)
	}
	files := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			log.Fatalf("ファイル情報の取得に失敗しました: %v", err)
		}
		files = append(files, info)
	}
	for _, file := range files {
		for i := 1; i < len(os.Args); i++ {
			argsBase := filepath.Base(os.Args[i])
			fileBase := filepath.Base(file.Name())
			if strings.TrimSuffix(argsBase, filepath.Ext(argsBase)) == strings.TrimSuffix(fileBase, filepath.Ext(fileBase)) {
				rmfile := filepath.Join(dir, file.Name())
				log.Printf("ファイル: %vを削除します。\n", rmfile)
				if err := os.Remove(rmfile); err != nil {
					log.Fatalf("ファイルの削除に失敗しました: %v", err)
				}
				break
			}
		}
	}
}
