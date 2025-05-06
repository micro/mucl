package mucl

import "github.com/iancoleman/strcase"

func (m *Message) FileName() string {
	if m == nil {
		return ""
	}
	return strcase.ToSnake(m.Name) + ".go"
}

func (m *Message) Fields() []*Field {
	if m == nil {
		return nil
	}
	var fields []*Field
	for _, entry := range m.Entries {
		if entry.Field != nil {
			fields = append(fields, entry.Field)
		}
	}
	return fields
}

func (m *Message) Messages() []*Message {
	if m == nil {
		return nil
	}
	var msgs []*Message
	for _, entry := range m.Entries {
		if entry.Message != nil {
			msgs = append(msgs, entry.Message)
		}
	}
	return msgs
}

func (m *Message) Enums() []*Enum {
	if m == nil {
		return nil
	}
	var enums []*Enum
	for _, entry := range m.Entries {
		if entry.Enum != nil {
			enums = append(enums, entry.Enum)
		}
	}
	return enums
}

func (m *Message) Oneofs() []*Oneof {
	if m == nil {
		return nil
	}
	var oneofs []*Oneof
	for _, entry := range m.Entries {
		if entry.Oneof != nil {
			oneofs = append(oneofs, entry.Oneof)
		}
	}
	return oneofs
}

func (m *Message) Options() []*Option {
	if m == nil {
		return nil
	}
	var options []*Option
	for _, entry := range m.Entries {
		if entry.Option != nil {
			options = append(options, entry.Option)
		}
	}
	return options
}

func (m *Message) Reserved() []*Reserved {
	if m == nil {
		return nil
	}
	var reserved []*Reserved
	for _, entry := range m.Entries {
		if entry.Reserved != nil {
			reserved = append(reserved, entry.Reserved)
		}
	}
	return reserved
}

func (m *Message) Extensions() []*Extensions {
	if m == nil {
		return nil
	}
	var extensions []*Extensions
	for _, entry := range m.Entries {
		if entry.Extensions != nil {
			extensions = append(extensions, entry.Extensions)
		}
	}
	return extensions
}

func (m *Message) Extend() []*Extend {
	if m == nil {
		return nil
	}
	var extends []*Extend
	for _, entry := range m.Entries {
		if entry.Extend != nil {
			extends = append(extends, entry.Extend)
		}
	}
	return extends
}
