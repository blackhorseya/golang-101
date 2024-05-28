package proxy

// Image 定義圖像的接口
type Image interface {
	Display() string
}

// RealImage 是真實的圖像對象
type RealImage struct {
	FileName string
}

func (ri *RealImage) Display() string {
	return "Displaying " + ri.FileName
}

// NewRealImage 用於創建 RealImage 對象
func NewRealImage(fileName string) *RealImage {
	// 模擬從磁盤加載圖像的昂貴操作
	return &RealImage{FileName: fileName}
}

// ProxyImage 是圖像的代理，控制對 RealImage 的訪問
type ProxyImage struct {
	FileName  string
	realImage *RealImage
}

func (pi *ProxyImage) Display() string {
	if pi.realImage == nil {
		pi.realImage = NewRealImage(pi.FileName)
	}
	return pi.realImage.Display()
}

func NewProxyImage(fileName string) *ProxyImage {
	return &ProxyImage{FileName: fileName}
}
