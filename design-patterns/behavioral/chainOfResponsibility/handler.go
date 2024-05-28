package chain

// Handler 定義了處理請求的接口
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(level int) string
}
