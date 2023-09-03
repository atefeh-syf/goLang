package examples

import "context"

type KeyValue string

func SendValueExample() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, KeyValue("api-key"), "1234567890")
	ctx = context.WithValue(ctx, KeyValue("validation-code"), "adsfc444fsd")
	Print(ctx)
}

func Print(ctx context.Context) {
	var apiKey = ctx.Value(KeyValue("api-key"))
	var validationCode = ctx.Value(KeyValue("validation-code"))

	println(apiKey.(string), validationCode.(string))
}
