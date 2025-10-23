package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	DEFAULT_TIMEOUT = 30 * time.Second
	SIZE_KB         = 1 << (10 * 1)
)

func main() {
	testDiskSymbol()
}

func getPcieTopologyFromSysFile() (map[string][]string, error) {
	/*
	   /sys/devices/pci0000:42/0000:42:02.0/0000:43:00.0/0000:44:08.0/0000:5d:00.0/0000:5e:00.0/0000:5f:00.0
	*/

	buf := "/sys/devices/pci0000:42/0000:42:02.0/0000:43:00.0/0000:44:08.0/0000:5d:00.0/0000:5e:00.0/0000:5f:00.0"
	ret := make(map[string][]string)
	lines := strings.Split(strings.TrimSuffix(string(buf), "\n"), "\n")
	for _, line := range lines {
		pciline := strings.TrimSpace(line)
		addrs := strings.Split(pciline, "pci")
		if len(addrs) != 2 {
			continue
		}
		pcieTree := addrs[1]
		pcies := strings.Split(pcieTree, "/")
		if len(pcies) < 3 {
			// rc rootPort endpoint
			continue
		}
		// pcieRC := pcies[0]
		pcieRootPort := pcies[1]
		//endpoint := pcies[len(pcies)-1]
		pos := strings.Index(pcieTree, "/")
		//lastPos := strings.LastIndex(pcieTree, "/")
		if _, ok := ret[pcieRootPort]; !ok {
			ret[pcieRootPort] = make([]string, 0, 10)
		}
		ret[pcieRootPort] = append(ret[pcieRootPort], pcieTree[pos+1:])
	}

	return ret, nil
}

func testReadConsoleFile() {
	file, modifyTime, err := readConsoleLogFile("/Users/lijinlong100/workspace/caiflower-go/cmd/test.log", 25, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
	fmt.Println(modifyTime.Format("2006-01-02 15:04:05"))
}

func readConsoleLogFile(path string, reqLines, maxKB int) (string, time.Time, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", time.Time{}, err
	}
	defer f.Close()

	fileLen, err := f.Seek(0, 2)
	if err != nil {
		return "", time.Time{}, err
	}

	if maxKB == 0 {
		maxKB = 100
	}
	stat, err := f.Stat()
	if err != nil {
		return "", time.Time{}, err
	}

	maxSize := int64(maxKB * SIZE_KB)
	if maxSize > fileLen {
		maxSize = fileLen
	}
	if _, err := f.Seek(-maxSize, 2); err != nil {
		return "", time.Time{}, err
	}

	chunks := make([]byte, maxSize)
	if _, err := f.Read(chunks); err != nil {
		return "", time.Time{}, err
	}

	logContent := string(chunks)
	logContent = strings.TrimSuffix(logContent, "\n")
	logLines := strings.Split(logContent, "\n")

	// 大于1行的时候截断对齐
	if len(logLines) > 1 {
		logLines = logLines[1:]
	}

	realLines := len(logLines)
	if reqLines == 0 || realLines <= reqLines {
		logContent = strings.Join(logLines, "\n")
	} else if realLines > reqLines {
		logContent = strings.Join(logLines[realLines-reqLines:], "\n")
	}

	return logContent, stat.ModTime(), nil
}

func imageTest() {
	base64Image, err := os.ReadFile("/Users/lijinlong100/workspace/caiflower-go/cmd/image.ppn")

	decodeBytes, _ := base64.StdEncoding.DecodeString(string(base64Image))

	img, err := decodePPM(decodeBytes)
	if err != nil {
		panic(err)
	}

	// 2. 编码为PNG
	var pngBuf bytes.Buffer
	//err = png.Encode(&pngBuf, img)
	err = jpeg.Encode(&pngBuf, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		panic(err)
	}

	toString := base64.StdEncoding.EncodeToString(pngBuf.Bytes())
	fmt.Println(toString)
}

func decodePPM(data []byte) (image.Image, error) {
	r := bytes.NewReader(data)
	// 读取 PPM 头部（魔数、宽、高、最大色值），跳过注释
	var header [4]string
	n := 0
	buf := make([]byte, 1)
	token := ""
	for n < 4 {
		if _, err := r.Read(buf); err != nil {
			return nil, err
		}
		b := buf[0]
		if b == '#' {
			// 跳过注释行
			for {
				if _, err := r.Read(buf); err != nil {
					return nil, err
				}
				if buf[0] == '\n' {
					break
				}
			}
			continue
		}
		if b == ' ' || b == '\n' || b == '\t' || b == '\r' {
			if len(token) > 0 {
				header[n] = token
				n++
				token = ""
			}
		} else {
			token += string(b)
		}
	}
	if header[0] != "P6" {
		return nil, fmt.Errorf("only P6 PPM supported")
	}
	width, _ := strconv.Atoi(header[1])
	height, _ := strconv.Atoi(header[2])
	maxval, _ := strconv.Atoi(header[3])
	if maxval != 255 {
		return nil, fmt.Errorf("only maxval=255 supported")
	}

	// 读取像素数据
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	pixel := make([]byte, width*height*3)
	if _, err := io.ReadFull(r, pixel); err != nil {
		return nil, err
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := (y*width + x) * 3
			r, g, b := pixel[idx], pixel[idx+1], pixel[idx+2]
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	return img, nil
}

func testDiskSymbol() {
	str := "vda,vdb,vdc,vdd,vde,vdf,vdg,vdh,vdi,vdj,vdk,vdl,vdm,vdn,vdo,vdp,vdq,vdr,vds,vdt,vdu,vdv,vdw,vdx,vdy,vdz,vdaa,vdab,vdac,vdad,vdae,vdaf,vdag,vdah,vdai,vdaj,vdak,vdal,vdam,vdan,vdao,vdap,vdaq,vdar,vdas,vdat,vdau,vdav,vdaw,vdax,vday,vdaz,vdba,vdbb,vdbc,vdbd,vdbe,vdbf,vdbg,vdbh,vdbi,vdbj,vdbk,vdbl,vdbm"
	slice := strings.Split(str, ",")
	fmt.Println(len(slice))
}
