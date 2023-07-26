// Understanding Context

package main

import (
	"context"
	"fmt"
)

func doSome(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSome(ctx)
}
