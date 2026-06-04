package upload

import (
	"encoding/csv"
	"io"
	"os"

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

		// fmt.Println(people)

		err = s.repo.CsvUpload(&people)
		if err != nil {
			return err
		}
	}

	return nil
}
