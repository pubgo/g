package main

import (
	"flag"
	"fmt"
	// "strings"

	"log"
	"sync"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
	"io/ioutil"
)

const parallelism = 3

type input struct {
	Mat      gocv.Mat
	FileName string
}

// func GetFileContentType(out *os.File) (string, error) {

// 	// Only the first 512 bytes are used to sniff the content type.
// 	buffer := make([]byte, 512)

// 	_, err := out.Read(buffer)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Use the net/http package's handy DectectContentType function. Always returns a valid
// 	// content-type by returning "application/octet-stream" if no others seemed to match.
// 	contentType := http.DetectContentType(buffer)

// 	return contentType, nil
// }

var (
	useAll            = flag.Bool("all", false, "Compute all hashes")
	usePHash          = flag.Bool("phash", false, "Compute PHash")
	useAverage        = flag.Bool("average", false, "Compute AverageHash")
	useBlockMean0     = flag.Bool("blockmean0", false, "Compute BlockMeanHash mode 0")
	useBlockMean1     = flag.Bool("blockmean1", false, "Compute BlockMeanHash mode 1")
	useColorMoment    = flag.Bool("colormoment", false, "Compute ColorMomentHash")
	useMarrHildreth   = flag.Bool("marrhildreth", false, "Compute MarrHildrethHash")
	useRadialVariance = flag.Bool("radialvariance", false, "Compute RadialVarianceHash")
)

func setupHashes() []contrib.ImgHashBase {
	var hashes []contrib.ImgHashBase

	if *usePHash || *useAll {
		hashes = append(hashes, contrib.PHash{})
	}
	if *useAverage || *useAll {
		hashes = append(hashes, contrib.AverageHash{})
	}
	if *useBlockMean0 || *useAll {
		hashes = append(hashes, contrib.BlockMeanHash{})
	}
	if *useBlockMean1 || *useAll {
		hashes = append(hashes, contrib.BlockMeanHash{Mode: contrib.BlockMeanHashMode1})
	}
	if *useColorMoment || *useAll {
		hashes = append(hashes, contrib.ColorMomentHash{})
	}
	if *useMarrHildreth || *useAll {
		// MarrHildreth has default parameters for alpha/scale
		hashes = append(hashes, contrib.NewMarrHildrethHash())
	}
	if *useRadialVariance || *useAll {
		// RadialVariance has default parameters too
		hashes = append(hashes, contrib.NewRadialVarianceHash())
	}

	// If no hashes were selected, behave as if all hashes were selected
	if len(hashes) == 0 {
		*useAll = true
		return setupHashes()
	}

	return hashes
}

// func hasher(fileName string)

func main() {
	flag.Usage = func() {
		fmt.Println("How to run:\n\timg-similarity [-flags] [image1.jpg] [image2.jpg]")
		flag.PrintDefaults()
	}

	var procDir string
	flag.StringVar(&procDir, "d", ".", "Directory path to process")

	flag.Parse()

	theMap := &map[gocv.Mat]string{}
	hash := contrib.AverageHash{}

	mapChan, mapWaitGroup := startMapper(theMap)
	hasherChan, hashWaitGroup := startHasher(mapChan, hash)
	readerWaitGroup := startReaders(hasherChan, procDir)

	// Cascade cleanup
	readerWaitGroup.Wait()
	close(hasherChan)
	hashWaitGroup.Wait()
	close(mapChan)
	mapWaitGroup.Wait()

	//// Print map
	// for hash, word := range *theMap {
	// 	fmt.Println(hash, ": ", word)
	// }

	for testHash, testPath := range *theMap {
		fmt.Printf("file://%s :\n", testPath)
		for galHash, galPath := range *theMap {
			if similar := hash.Compare(testHash, galHash); similar < 8 && similar > 0 {
				fmt.Printf("\t%g diff | file://%s\n", similar, galPath)
			}
		}
	}
}

func startReaders(hasherChannel chan *input, directoryName string) *sync.WaitGroup {
	readerWaitGroup := &sync.WaitGroup{}

	// read images
	fmt.Println("Processing from \"" + directoryName + "\"")
	files, err := ioutil.ReadDir(directoryName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		readerWaitGroup.Add(1)
		go reader(hasherChannel, readerWaitGroup, directoryName+"/"+file.Name())
	}

	return readerWaitGroup
}

func reader(hasherChannel chan *input, readerWaitGroup *sync.WaitGroup, pathname string) {

	img := gocv.IMRead(pathname, gocv.IMReadColor)

	if !img.Empty() {
		hasherChannel <- &input{
			Mat:      img,
			FileName: pathname,
		}
	}

	readerWaitGroup.Done()
}

func startHasher(mapperChannel chan *input, hash contrib.ImgHashBase) (chan *input, *sync.WaitGroup) {
	hasherWaitGroup := &sync.WaitGroup{}

	hasherChannel := make(chan *input)

	hasherWaitGroup.Add(1)
	go hasher(mapperChannel, hasherChannel, hasherWaitGroup, hash)

	return hasherChannel, hasherWaitGroup
}

func hasher(mapperChannel chan *input, hasherChannel chan *input, hasherWaitGroup *sync.WaitGroup, hash contrib.ImgHashBase) {
	for p := range hasherChannel {
		result := gocv.NewMat()
		defer p.Mat.Close()
		hash.Compute(p.Mat, &result)
		if result.Empty() {
			fmt.Printf("error computing hash for %s\n", p.FileName)
			continue
		}

		mapperChannel <- &input{
			Mat:      result,
			FileName: p.FileName,
		}
	}

	hasherWaitGroup.Done()
}

func startMapper(theMap *map[gocv.Mat]string) (chan *input, *sync.WaitGroup) {
	mapperWaitGroup := &sync.WaitGroup{}

	mapperChannel := make(chan *input)

	mapperWaitGroup.Add(1)

	go mapper(theMap, mapperChannel, mapperWaitGroup)

	return mapperChannel, mapperWaitGroup
}

func mapper(theMap *map[gocv.Mat]string, mapperChannel chan *input, mapperWaitGroup *sync.WaitGroup) {
	for input := range mapperChannel {
		(*theMap)[input.Mat] = input.FileName
	}

	mapperWaitGroup.Done()
}
