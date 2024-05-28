package builder

// House 定義一個房子的結構
type House struct {
	Foundation string
	Walls      string
	Roof       string
}

// SetFoundation 設定房子的地基.
func (h *House) SetFoundation(foundation string) {
	h.Foundation = foundation
}

// SetWalls 設定房子的牆壁.
func (h *House) SetWalls(walls string) {
	h.Walls = walls
}

// SetRoof 設定房子的屋頂.
func (h *House) SetRoof(roof string) {
	h.Roof = roof
}

// HouseBuilder 定義建造房子的步驟
type HouseBuilder interface {
	BuildFoundation()
	BuildWalls()
	BuildRoof()
	GetHouse() *House
}

// ConcreteHouseBuilder 實現 HouseBuilder 接口
type ConcreteHouseBuilder struct {
	house *House
}

// NewConcreteHouseBuilder 創建 ConcreteHouseBuilder 實例.
func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
	return &ConcreteHouseBuilder{house: &House{}}
}

func (b *ConcreteHouseBuilder) BuildFoundation() {
	b.house.SetFoundation("Concrete Foundation")
}

func (b *ConcreteHouseBuilder) BuildWalls() {
	b.house.SetWalls("Brick Walls")
}

func (b *ConcreteHouseBuilder) BuildRoof() {
	b.house.SetRoof("Slate Roof")
}

func (b *ConcreteHouseBuilder) GetHouse() *House {
	return b.house
}

// Director 負責按順序構建房子
type Director struct {
	builder HouseBuilder
}

// NewDirector 創建 Director 實例.
func NewDirector(builder HouseBuilder) *Director {
	return &Director{builder: builder}
}

// Construct 構建房子.
func (d *Director) Construct() {
	d.builder.BuildFoundation()
	d.builder.BuildWalls()
	d.builder.BuildRoof()
}
