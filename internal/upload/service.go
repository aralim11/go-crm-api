package upload

type UploadService interface {
	ImageUpload()
}

type uploadService struct {
	repo UploadRepository
}

func NewUploadService(repo UploadRepository) UploadService {
	return &uploadService{repo: repo}
}

func (s *uploadService) ImageUpload() {

}
