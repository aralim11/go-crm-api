package upload

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"

	"github.com/aralim11/go-crm-api/internal/utils/validator"
)

type UploadService interface {
	PeopleList(page int, limit int) ([]*People, error)
	ProcessCSV(file string) error
	Count() (int, error)
}

type uploadService struct {
	repo UploadRepository
}

func NewUploadService(repo UploadRepository) UploadService {
	return &uploadService{repo: repo}
}

func (s *uploadService) PeopleList(page int, limit int) ([]*People, error) {
	peoples, err := s.repo.PeopleList(page, limit)
	if err != nil {
		return nil, err
	}

	return peoples, nil
}

func (s *uploadService) Count() (int, error) {
	count, err := s.repo.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *uploadService) ProcessCSV(dstPath string) error {
	// open csv from server using path
	csvFile, err := os.Open(dstPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// read csv from server
	reader := csv.NewReader(csvFile)

	// queue jobs worker
	jobs := make(chan []People, 10)
	var wg sync.WaitGroup
	workerCount := 5
	batchSize := 200

	// start workers
	for i := 1; i < workerCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for batchData := range jobs {
				log.Printf("Worker %d processing batch size: %d\n", id, len(batchData))

				err := s.repo.CsvUpload(batchData)
				if err != nil {
					log.Printf("Worker %d error: %v\n", id, err)
				}

				log.Printf("Worker %d success.\n", id)
			}
		}(i)
	}

	var batch []People
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// push data on model
		people := People{
			UserID:      record[0],
			FirstName:   record[1],
			LastName:    record[2],
			Sex:         record[3],
			Email:       record[4],
			Phone:       record[5],
			DateOfBirth: validator.ParseDate(record[6]),
			JobTitle:    record[7],
		}

		batch = append(batch, people)

		// when batch full → send to worker
		if len(batch) == batchSize {
			jobs <- batch
			batch = nil // reset
		}
	}

	if len(batch) > 0 {
		jobs <- batch
	}

	close(jobs)
	wg.Wait()
	return nil
}
