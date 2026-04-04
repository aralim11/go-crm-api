package report

type ReportService interface {
	Search()
}

type reportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) Search() {
	s.repo.SearchData()
}
