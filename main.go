package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

var (
	d = flag.String("d", ".", "Directory to process")
	a = flag.Bool("a", false, "Print all info")
	h = flag.Bool("h", false, "Normal file size (ex. 1024 -> 2 KB)")
	s = flag.String("s", "no", "Files sort option")
)

func hrSize(fsize int64) string {
	newSize := fsize

	start := 1024
	prefix := "КБ"

	if newSize >= 1048576 {
		start = 1048576
		prefix = "МБ"
	}

	for i := 1; ; i++ {
		if fsize < int64(start * i) {
			newSize = int64(i)
			break
		}
	}

	return strconv.Itoa(int(newSize)) + prefix
}

func printAll(file os.FileInfo, h bool) {
	time := file.ModTime().Format("Jan 06 15:4")

	var fSize string

	if h {
		fSize = hrSize(file.Size())
	} else {
		fSize = strconv.Itoa(int(file.Size()))
	}

	fmt.Printf("%s %s %s \n", fSize, time, file.Name())
}

type DateSort []os.FileInfo

func (ds DateSort) Len() int {
	return len(ds)
}

func (ds DateSort) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func (ds DateSort) Less(i, j int) bool {
	return ds[i].ModTime().Second() < ds[j].ModTime().Second()
}

type SizeSort []os.FileInfo

func (ss SizeSort) Len() int {
	return len(ss)
}

func (ss SizeSort) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss SizeSort) Less(i, j int) bool {
	return ss[i].Size() < ss[j].Size()
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)

	if *s == "date" {
		sort.Sort(DateSort(files))
	} else if *s == "size" {
		sort.Sort(SizeSort(files))
	}

	for _, file := range files {
		if *a {
			printAll(file, *h)
		} else {
			fmt.Println(file.Name())
		}
	}
}
