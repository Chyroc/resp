package resp

func NewWithErr(err error) *Reply {
	if err != nil {
		return &Reply{errNotFromReply: err}
	}
	return nil
}

func NewWithBytes(bs []byte) *Reply {
	return &Reply{str: string(bs)}
}

func NewWithInt64(i int64) *Reply {
	return &Reply{integer: i}
}

func NewWithNull() *Reply {
	return &Reply{null: true}
}

func NewWithReplies(replies []*Reply) *Reply {
	return &Reply{replys: replies}
}
