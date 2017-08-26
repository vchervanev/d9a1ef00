package zip

import (
	"archive/zip"
	"bufio"
	"log"
)

type RecordHandler func(content []byte, filter byte)

// LoadObjects unzips `filename`, skips bytes until `[` is found
// and then invokes `handler` for every record separated by `}`
// After each `}` bytes unless a comma found is skipped
//
// `sequence` defines the order of files to be handled (by the 1st char of their filename)
func LoadObjects(filename string, sequence []byte, handler RecordHandler) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, filter := range sequence {
		for _, file := range r.File {
			if file.FileHeader.Name[0] != filter {
				continue
			}
			log.Printf("Processing %v... ", file.FileHeader.Name)
			rc, _ := file.Open()
			buf := bufio.NewReader(rc)
			buf.ReadBytes('[')
			for {
				tmp, _ := buf.ReadBytes('}')

				if len(tmp) == 0 {
					break
				}

				handler(tmp, filter)
				buf.ReadBytes(',')
			}

			rc.Close()
			log.Println("Done")
		}
	}
}
