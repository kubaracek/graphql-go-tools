package logger

type Field struct {
	kind           FieldKind
	Key            string
	StringValue    string
	StringsValue   []string
	IntValue       int64
	ByteValue      []byte
	InterfaceValue interface{}
	ErrorValue     error
}

type FieldKind int

const (
	StringField FieldKind = iota + 1
	StringsField
	IntField
	BoolField
	ByteStringField
	InterfaceField
	ErrorField
	NamedErrorField
)

func Any(key string, value interface{}) Field {
	return Field{
		kind:           InterfaceField,
		Key:            key,
		InterfaceValue: value,
	}
}

func Error(err error) Field {
	return Field{
		kind:       ErrorField,
		Key:        "error",
		ErrorValue: err,
	}
}

func NamedError(key string, err error) Field {
	return Field{
		kind:       NamedErrorField,
		Key:        key,
		ErrorValue: err,
	}
}

func String(key, value string) Field {
	return Field{
		kind:        StringField,
		Key:         key,
		StringValue: value,
	}
}

func Strings(key string, value []string) Field {
	return Field{
		Key:          key,
		kind:         StringsField,
		StringsValue: value,
	}
}

func Int(key string, value int) Field {
	return Field{
		kind:     IntField,
		Key:      key,
		IntValue: int64(value),
	}
}

func Bool(key string, value bool) Field {
	var integer int64
	if value {
		integer = 1
	}
	return Field{
		kind:     BoolField,
		Key:      key,
		IntValue: integer,
	}
}

func ByteString(key string, value []byte) Field {
	return Field{
		kind:      ByteStringField,
		Key:       key,
		ByteValue: value,
	}
}
