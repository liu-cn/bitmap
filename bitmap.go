package bitmap

type BitMap struct {
	bits        []bit
	maxIndex    int64
	minIndex    int64
	maxCount    int64
	setCount    int64 //置为1的数量
	statusCount int64
}

func (m *BitMap) SetState(index int64, status bool) {
	var bit uint8
	if index < 8 {
		index = 0
		bit = uint8(index)
	} else {
		index = index / 8
		bit = uint8(index % 8)
	}
	if status {
		m.statusCount++
		m.bits[index].setBit1(bit)
	} else {
		m.bits[index].setBit0(bit)
	}
}

func New(maxCount int64) *BitMap {
	return &BitMap{
		maxCount: maxCount,
		bits:     make([]bit, (maxCount/8)+1),
	}
}

func (m *BitMap) GetState(index int64) bool {
	var bit uint8
	if index < 8 {
		index = 0
		bit = uint8(index)
	} else {
		index = index / 8
		bit = uint8(index % 8)
	}
	return m.bits[index].getBit(bit)
}

type bit uint8

// SetBit1 将某一位置为1
func (b *bit) setBit1(p uint8) {
	*b = *b | (1 << p)
	//fmt.Printf("%v: %b\n", *b, *b)
}

// SetBit0 将某一位置为0
func (b *bit) setBit0(p uint8) {
	*b = *b &^ (1 << p)
	//fmt.Printf("%v: %b\n", *b, *b)
}
func (b *bit) getBit(bit uint8) bool {
	eb := *b
	eb.setBit1(bit) //把该位置为1，看该值是否有变化，如果有变化，说明该位置为0，如果没变化说明该位为1
	return eb == *b
}
