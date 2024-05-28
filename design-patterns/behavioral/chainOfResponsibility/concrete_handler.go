package chain

// BaseHandler 提供了處理者接口的基本實現
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) HandleRequest(level int) string {
	if b.next != nil {
		return b.next.HandleRequest(level)
	}
	return "No handler available"
}

// ConcreteHandler1 處理普通請求
type ConcreteHandler1 struct {
	BaseHandler
}

func (h *ConcreteHandler1) HandleRequest(level int) string {
	if level == 1 {
		return "ConcreteHandler1 handled the request"
	}
	return h.BaseHandler.HandleRequest(level)
}

// ConcreteHandler2 處理高級請求
type ConcreteHandler2 struct {
	BaseHandler
}

func (h *ConcreteHandler2) HandleRequest(level int) string {
	if level == 2 {
		return "ConcreteHandler2 handled the request"
	}
	return h.BaseHandler.HandleRequest(level)
}

// ConcreteHandler3 處理最高級請求
type ConcreteHandler3 struct {
	BaseHandler
}

func (h *ConcreteHandler3) HandleRequest(level int) string {
	if level == 3 {
		return "ConcreteHandler3 handled the request"
	}
	return h.BaseHandler.HandleRequest(level)
}
