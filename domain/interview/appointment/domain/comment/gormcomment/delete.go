package gormcomment

func (r repository) DeleteById(id uint) error {
	return r.db.Delete(&Comment{}, id).Error
}
