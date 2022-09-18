package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
)

func ChangeContact(file *os.File) (*os.File, error) {
	fmt.Println("CHAnge")
	correctFile, err := os.Create("newContacts.vcf")
	if err != nil {
		return nil, err
	}
	defer func(correctFile *os.File) {
		err := correctFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(correctFile)

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "TEL;") {
			if line[9:15] == "+38071" {
				_, err := correctFile.WriteString("TEL;CELL:+7949")
				if err != nil {
					return nil, err
				}
				_, err = correctFile.WriteString(line[15:])
				if err != nil {
					return nil, err
				}
			} else if line[9:12] == "071" {
				_, err := correctFile.WriteString("TEL;CELL:+7949")
				if err != nil {
					return nil, err
				}
				correctLine := strings.ReplaceAll(line[12:], "-", "")
				_, err = correctFile.WriteString(correctLine)
				if err != nil {
					return nil, err
				}
			} else {
				_, err := correctFile.WriteString(line)
				if err != nil {
					return nil, err
				}
			}
			continue
		}
		_, err = correctFile.WriteString(line)
		if err != nil {
			return nil, err
		}
	}
	return correctFile, err
}

func ChangeContactsMultiPart(file multipart.File) (*bytes.Buffer, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "./contacts.vcf")
	if err != nil {
		return nil, err
	}
	bodyBuf.Truncate(0)
	fileScanner := bufio.NewScanner(file)
	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(fileScanner.Text(), "TEL;") {
			line = strings.ReplaceAll(line, ";PREF", ":")
			line = strings.ReplaceAll(line, "-", "")
			if strings.HasPrefix(line[9:], "+38071") || strings.HasPrefix(line[9:], "380-71") {
				line = "TEL;CELL;PREF:+7949" + line[15:]
			} else if strings.HasPrefix(line[9:], "071") {
				line = "TEL;CELL;PREF:+7949" + line[12:]
			} else {
				lines = append(lines, fileScanner.Text())
				continue
			}
			lines = append(lines, line)
			lines = append(lines, fileScanner.Text())
			continue
		}
		lines = append(lines, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		return nil, err
	}

	_, err = fileWriter.Write([]byte(strings.Join(lines, "\n")))
	if err != nil {
		return nil, err
	}
	return bodyBuf, nil
}
