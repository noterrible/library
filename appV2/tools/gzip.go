package tools

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// 对象压缩成string，方便直接存入redis
func ToGzip[T any](res T) (string, error) {
	// 创建 gzip 编码器
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)

	// 将切片转换为 JSON 数据，并压缩
	jsonBytes, err := json.Marshal(&res)
	if err != nil {
		fmt.Printf("Error during  json.Marshal :%+v\n", err.Error())
		return "", err
	}
	if _, err := gzipWriter.Write(jsonBytes); err != nil {
		fmt.Printf("Error during  gzipWriter.Write :%+v\n", err.Error())
		return "", err

	}
	if err := gzipWriter.Close(); err != nil {
		fmt.Printf("Error during gzipWriter.Close :%+v\n", err.Error())
		return "", err
	}

	return string(buf.Bytes()), nil
}

// string解压成对象，方便redis数据解压后直接返回数据
func GzipTo[T any](res *T, key string) (err error) {
	// 创建 gzip 解压缩器
	gzipReader, err := gzip.NewReader(bytes.NewReader([]byte(key)))
	defer func() {
		if errC := gzipReader.Close(); errC != nil {
			fmt.Printf("Error during gzipReader.Close: %+v\n", errC)
		}
	}()
	if err != nil {
		return err
	}
	// 读取并解压缩 JSON 数据
	jsonBytes, errR := ioutil.ReadAll(gzipReader)
	if errR != nil {
		return errors.New("读取解压缩JSON错误")
	}
	// 解码 JSON 数据
	if errU := json.Unmarshal(jsonBytes, res); errU != nil {
		return errors.New("解码JSON错误")
	}
	return nil
}
