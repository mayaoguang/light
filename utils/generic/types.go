package generic

// 使用泛型封装的常用方法

type IntEx interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

type UIntEx interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type FloatEx interface {
	~float32 | ~float64
}

type Float interface {
	float32 | float64
}

type NumberEx interface {
	IntEx | UIntEx | FloatEx
}

type Number interface {
	Int | UInt | Float
}

type StringEx interface {
	~string
}

type DataType interface {
	NumberEx | StringEx
}
