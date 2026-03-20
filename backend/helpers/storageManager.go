package helpers

import (
	"fmt"

	"github.com/theerudito/istla/model/dto"
)

func StorageManager(obj dto.StorageItemDTO) (int, error) {
	var storageId int
	var err error

	switch obj.Option {
	case "INSERT":
		qInsert := `INSERT INTO storage (nombre, url, extencion) VALUES ($1, $2, $3) RETURNING id_storage`
		err = obj.TX.QueryRow(qInsert, obj.FileName, obj.Url, obj.Extension).Scan(&storageId)
	case "UPDATE":
		qUpdate := `UPDATE storage SET nombre = $1, url = $2, extencion = $3 WHERE id_storage = $4 RETURNING id_storage`
		err = obj.TX.QueryRow(qUpdate, obj.FileName, obj.Url, obj.Extension, obj.StorageId).Scan(&storageId)
	case "DELETE":
		_, err = obj.TX.Exec(`DELETE FROM storage WHERE id_storage = $1`, obj.StorageId)
		storageId = obj.StorageId
	}

	if err != nil {
		_ = InsertLogsError(obj.TX, "storage", fmt.Sprintf("error en StorageManager opción %s: %v", obj.Option, err))
		return 0, err
	}

	return storageId, nil
}
