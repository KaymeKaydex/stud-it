package service

func (s *Service) GetYoungOrOld(age int) string {
	if age > 20 {
		return "old"
	} else {
		return "young"
	}
}
