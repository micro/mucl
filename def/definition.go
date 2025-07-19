package mucl

func (d *Definition) ServiceName() string {
	// if d == nil {
	// 	return ""
	// }
	// for _, entry := range d.Entries {
	// 	if entry.Service != nil {
	// 		return strings.ReplaceAll(entry.Service.Name, "\"", "")
	// 	}
	// }
	return ""
}

func (d *Definition) Import() string {
	if d == nil {
		return ""
	}
	for _, entry := range d.Entries {
		if entry.Import != "" {
			return entry.Import
		}
	}
	return ""
}

func (d *Definition) Messages() []*Message {
	if d == nil {
		return nil
	}
	var messages []*Message
	for _, entry := range d.Entries {
		if entry.Message != nil {
			messages = append(messages, entry.Message)
		}
	}
	return messages
}

func (d *Definition) Servers() []*Endpoint {
	if d == nil {
		return nil
	}
	var servers []*Endpoint
	for _, entry := range d.Entries {
		if entry.Endpoint != nil {
			servers = append(servers, entry.Endpoint)
		}
	}
	return servers
}

func (d *Definition) Enums() []*Enum {
	if d == nil {
		return nil
	}
	var enums []*Enum
	for _, entry := range d.Entries {
		if entry.Enum != nil {
			enums = append(enums, entry.Enum)
		}
	}
	return enums
}
