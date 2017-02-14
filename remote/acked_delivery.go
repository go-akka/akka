package remote

type SeqNo struct {
	rawValue int64
}

func NewSeqNo(rawValue int64) *SeqNo {
	return &SeqNo{
		rawValue: rawValue,
	}
}

func (p *SeqNo) RawValue() int64 {
	return p.rawValue
}

func (p *SeqNo) IsSuccessor(that SeqNo) bool {
	return p.rawValue-that.RawValue() == 1
}

func (p *SeqNo) Inc() *SeqNo {
	nextValue := p.RawValue() + 1
	return NewSeqNo(nextValue)
}

func (p *SeqNo) CompareTo(other SeqNo) int {
	sgn := 0
	if p.RawValue() < other.RawValue() {
		sgn = -1
	} else if ((p.RawValue() - other.RawValue()) * int64(sgn)) < 0 {
		return -sgn
	}
	return sgn
}

type HasSequenceNumber interface {
	Seq() *SeqNo
}
