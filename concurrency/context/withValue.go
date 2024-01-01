package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	cCtx := context.WithValue(ctx, "shanliao", "333")
	ccCtx := context.WithValue(cCtx, "king", "222")
	fmt.Println(ctx.Value("shanliao"))
	fmt.Println(ctx.Value("king"))
	fmt.Println(cCtx.Value("king"))
	fmt.Println(ccCtx.Value("shanliao"))
}
