//go:generate goa gen  github/Shitomo/my-chat/design --output ../
package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("chat", func() {
	Title("Chat Service")
	Description("Service for chat manage")
	Server("chat", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("chat", func() {
	Description("The chat service.")

	Method("SendMessage", func() {
		Payload(SendMessageRequestBody)

		Result(SendMessageResponseBody)

		HTTP(func() {
			POST("/message")
			Response(StatusOK)
		})
	})
})

var ChatMessage = Type("Message", func() {
	Attribute("id", String, func() {
		Format(FormatUUID)
		Example("ac65bc46-3854-455f-82b9-54b6ec733b53")
	})
	Attribute("senderId", String, func() {
		Example("Golang coding")
	})
	Attribute("content", String, func() {
		Example("Golang coding")
	})
	Attribute("createdAt", String, func() {
		Format(FormatDateTime)
		Example("2006-01-02T15:04:05Z")
	})
	Attribute("updatedAt", String, func() {
		Format(FormatDateTime)
		Example("2006-01-02T15:04:05Z")
	})
})

var SendMessageRequestBody = Type("SendMessageRequestBody", func() {
	Attribute("senderId", String, func() {
		Example("Golang coding")
	})
	Attribute("content", String, func() {
		Example("Golang coding")
	})

})

var SendMessageResponseBody = Type("SendMessageResponseBody", func() {
	Attribute("message", ChatMessage, func() {})
})
