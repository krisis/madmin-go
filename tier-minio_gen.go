package madmin

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TierMinIO) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Endpoint":
			z.Endpoint, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Endpoint")
				return
			}
		case "AccessKey":
			z.AccessKey, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "AccessKey")
				return
			}
		case "SecretKey":
			z.SecretKey, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "SecretKey")
				return
			}
		case "Bucket":
			z.Bucket, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "Prefix":
			z.Prefix, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Prefix")
				return
			}
		case "Region":
			z.Region, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Region")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TierMinIO) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "Endpoint"
	err = en.Append(0x86, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Endpoint)
	if err != nil {
		err = msgp.WrapError(err, "Endpoint")
		return
	}
	// write "AccessKey"
	err = en.Append(0xa9, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.AccessKey)
	if err != nil {
		err = msgp.WrapError(err, "AccessKey")
		return
	}
	// write "SecretKey"
	err = en.Append(0xa9, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.SecretKey)
	if err != nil {
		err = msgp.WrapError(err, "SecretKey")
		return
	}
	// write "Bucket"
	err = en.Append(0xa6, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bucket)
	if err != nil {
		err = msgp.WrapError(err, "Bucket")
		return
	}
	// write "Prefix"
	err = en.Append(0xa6, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return
	}
	err = en.WriteString(z.Prefix)
	if err != nil {
		err = msgp.WrapError(err, "Prefix")
		return
	}
	// write "Region"
	err = en.Append(0xa6, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Region)
	if err != nil {
		err = msgp.WrapError(err, "Region")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TierMinIO) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Endpoint"
	o = append(o, 0x86, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Endpoint)
	// string "AccessKey"
	o = append(o, 0xa9, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.AccessKey)
	// string "SecretKey"
	o = append(o, 0xa9, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.SecretKey)
	// string "Bucket"
	o = append(o, 0xa6, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74)
	o = msgp.AppendString(o, z.Bucket)
	// string "Prefix"
	o = append(o, 0xa6, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.Prefix)
	// string "Region"
	o = append(o, 0xa6, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Region)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TierMinIO) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Endpoint":
			z.Endpoint, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Endpoint")
				return
			}
		case "AccessKey":
			z.AccessKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AccessKey")
				return
			}
		case "SecretKey":
			z.SecretKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SecretKey")
				return
			}
		case "Bucket":
			z.Bucket, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "Prefix":
			z.Prefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Prefix")
				return
			}
		case "Region":
			z.Region, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Region")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TierMinIO) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Endpoint) + 10 + msgp.StringPrefixSize + len(z.AccessKey) + 10 + msgp.StringPrefixSize + len(z.SecretKey) + 7 + msgp.StringPrefixSize + len(z.Bucket) + 7 + msgp.StringPrefixSize + len(z.Prefix) + 7 + msgp.StringPrefixSize + len(z.Region)
	return
}