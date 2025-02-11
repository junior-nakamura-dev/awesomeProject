package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	//seconds less than 5 expected see the Hotel booking cancelled! message
	//seconds up than 5 expected see the Room booked message
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	book(ctx)

}

func book(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled!")
	case <-time.After(5 * time.Second):
		fmt.Println("Room booked")
	}

}
