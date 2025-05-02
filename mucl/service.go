package mucl

func (s *Service) Broker() string {
	if s == nil {
		return ""
	}
	for _, entry := range s.Entries {
		if entry.Broker != nil {
			return entry.Broker.Name
		}
	}
	return ""
}

func (s *Service) Registry() string {
	if s == nil {
		return ""
	}
	for _, entry := range s.Entries {
		if entry.Registry != nil {
			return entry.Registry.Name
		}
	}
	return ""
}

func (s *Service) Transport() string {
	if s == nil {
		return ""
	}
	for _, entry := range s.Entries {
		if entry.Transport != nil {
			return entry.Transport.Name
		}
	}
	return ""
}

func (s *Service) Protocol() string {
	if s == nil {
		return ""
	}
	for _, entry := range s.Entries {
		if entry.Protocol != nil {
			return entry.Protocol.Name
		}
	}
	return ""
}

func (s *Service) Server() string {
	if s == nil {
		return ""
	}
	for _, entry := range s.Entries {
		if entry.Server != nil {
			return entry.Server.Name
		}
	}
	return ""
}
