package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const GeoipUrl = "https://raw.githubusercontent.com/runetfreedom/russia-v2ray-rules-dat/release/geoip.dat"
const GeositeUrl = "https://raw.githubusercontent.com/runetfreedom/russia-v2ray-rules-dat/release/geosite.dat"

func main() {
	tmpDirSrc := "/tmp/geo-xray"
	tmpGeoipSrc := filepath.Join(tmpDirSrc, "geoip.dat")
	tmpGeositeSrc := filepath.Join(tmpDirSrc, "geosite.dat")
	
	destDirSrc := "/usr/share/xray"
	geoipDest := filepath.Join(destDirSrc, "geoip.dat")
	geositeDest := filepath.Join(destDirSrc, "geosite.dat")

	// Создаем временную директорию
	if err := os.MkdirAll(tmpDirSrc, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		os.Exit(1)
	}

	// Скачиваем geoip.dat
	if err := downloadFile(GeoipUrl, tmpGeoipSrc); err != nil {
		fmt.Printf("Failed to download geoip: %v\n", err)
		os.Exit(1)
	}

	// Скачиваем geosite.dat
	if err := downloadFile(GeositeUrl, tmpGeositeSrc); err != nil {
		fmt.Printf("Failed to download geosite: %v\n", err)
		os.Exit(1)
	}

	// Копируем файлы в целевую директорию
	if err := copyFile(tmpGeoipSrc, geoipDest); err != nil {
		fmt.Printf("Failed to copy geoip: %v\n", err)
		os.Exit(1)
	}

	if err := copyFile(tmpGeositeSrc, geositeDest); err != nil {
		fmt.Printf("Failed to copy geosite: %v\n", err)
		os.Exit(1)
	}

	// Удаляем временную директорию
	if err := os.RemoveAll(tmpDirSrc); err != nil {
		fmt.Printf("Failed to remove directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Files downloaded and installed successfully")
}

// downloadFile скачивает файл по URL и сохраняет его по указанному пути
func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// copyFile копирует файл из источника в назначение
func copyFile(src, dst string) error {
	// Создаем целевую директорию, если она не существует
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
