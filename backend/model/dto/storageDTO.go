package dto

type StorageDTO struct {
	StorageId uint   `json:"id_storage"`
	FileName  string `json:"file_name"`
	Url       string `json:"url"`
	Extencion string `json:"extencion"`
}
