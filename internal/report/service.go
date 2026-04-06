package report

type ReportService interface {
	Search(barcode string, expiryDate string) error
}

type reportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) Search(barcode string, expiryDate string) error {
	result, _ := s.repo.SearchData()
	if  {
		
	}
}
