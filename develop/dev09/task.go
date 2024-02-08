package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func wget() {
	urlFlag := flag.String("url", "", "URL сайта для скачивания")
	outputFlag := flag.String("output", ".", "Каталог для сохранения скачанных файлов")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Необходимо указать URL сайта для скачивания.")
		return
	}

	baseURL, err := url.Parse(*urlFlag)
	if err != nil {
		fmt.Println("Ошибка при парсинге URL:", err)
		return
	}

	resp, err := http.Get(*urlFlag)
	if err != nil {
		fmt.Println("Ошибка при выполнении HTTP-запроса:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при парсинге HTML:", err)
		return
	}

	downloadResources(doc, baseURL, *outputFlag)
}

func downloadResources(n *html.Node, baseURL *url.URL, outputDir string) {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "link" || n.Data == "script") {
		for _, attr := range n.Attr {
			if attr.Key == "href" || attr.Key == "src" {
				resourceURL, err := baseURL.Parse(attr.Val)
				if err != nil {
					fmt.Println("Ошибка при формировании абсолютного URL:", err)
					continue
				}

				downloadFile(resourceURL.String(), outputDir)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		downloadResources(c, baseURL, outputDir)
	}
}

func downloadFile(url string, outputDir string) {
	fmt.Println("Скачивание файла:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении HTTP-запроса для ресурса", url, ":", err)
		return
	}
	defer resp.Body.Close()

	filePath := filepath.Join(outputDir, getFileName(url))

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла", filePath, ":", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Ошибка при копировании данных в файл", filePath, ":", err)
		return
	}
}

func getFileName(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
