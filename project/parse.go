package project

import "github.com/micro/mu/mucl"

func fromMuCL(mucl *mucl.Definition) (*Service, error) {
	serviceName := mucl.Service.Name
	// todo: put validations somewhere
	// with named errors
	service := &Service{Name: serviceName}
	service.EndpointMap = make(map[string]*Endpoint)
	service.MessageMap = make(map[string]*Message)
	service.EnumMap = make(map[string]*Enum)

	// parse enums
	service.parseEnums(mucl)
	// parse types
	err := service.parseTypes(mucl)
	if err != nil {
		return nil, err
	}
	// parse endpoints
	service.parseEndpoints(mucl)
	return service, nil
}

func (s *Service) parseTypes(mucl *mucl.Definition) error {
	// parse types
	types := mucl.Messages()
	for _, typ := range types {
		err := s.parseMessage(typ)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) parseMessage(m *mucl.Message) error {
	// parse message
	t := NewMessage(m.Name)
	// parse embedded messages
	for _, msg := range m.Messages() {
		// parse message
		err := s.parseMessage(msg)
		if err != nil {
			return err
		}
	}
	t.FieldMap = make(map[string]*Field)
	// parse fields
	for _, field := range m.Fields() {
		f := &Field{
			Name:     field.Name,
			TypeName: field.Type.String(),
			Required: field.Required,
			Repeated: field.Repeated,
		}
		t.FieldMap[field.Name] = f
	}
	// parse enums
	for _, enum := range m.Enums() {
		e := &Enum{
			Name: enum.Name,
		}
		// parse enum values
		for _, value := range enum.Values {
			v := &KeyValue{
				Key:   value.Value.Key,
				Value: value.Value.Value,
			}
			e.Values = append(e.Values, v)
		}
		s.EnumMap[enum.Name] = e
	}
	// parse options
	for _, option := range m.Options() {
		o := NewOption(option.Name, option.Attr, option.Value)
		t.Options[option.Name] = o
	}

	// add message to map
	s.MessageMap[m.Name] = t
	return nil
}

func (s *Service) parseEnums(mucl *mucl.Definition) {
	// parse enums
	enums := mucl.Enums()
	for _, typ := range enums {
		t := &Enum{
			Name: typ.Name,
		}
		// parse enum values
		for _, value := range typ.Values {
			v := &KeyValue{
				Key:   value.Value.Key,
				Value: value.Value.Value,
			}
			t.Values = append(t.Values, v)
		}
		s.EnumMap[typ.Name] = t
	}
}

func (s *Service) parseEndpoints(mucl *mucl.Definition) {
	// parse endpoints
	endpoints := mucl.Servers()
	for _, endpoint := range endpoints {
		ep := &Endpoint{
			Name:      endpoint.Name,
			MethodMap: make(map[string]*Method),
		}
		// parse methods
		mm := endpoint.Methods()
		for _, method := range mm {
			m := &Method{
				Name:              method.Name,
				RequestTypeName:   method.Request.String(),
				ResponseTypeName:  method.Response.String(),
				RequestStreaming:  method.StreamingRequest,
				ResponseStreaming: method.StreamingResponse,
			}
			ep.MethodMap[method.Name] = m
		}
		s.EndpointMap[endpoint.Name] = ep
	}
}
