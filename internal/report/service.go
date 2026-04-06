package report

type ReportService interface {
	SearchData(barcode string, expiryDate string) (*Product, error)
}

type reportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) SearchData(barcode string, expiryDate string) (*Product, error) {
	result, err := s.repo.SearchData(barcode, expiryDate)
	if err != nil {
		return nil, err
	}

	return result, nil
}
