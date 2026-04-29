package main

import (
	"context"
	"fmt"
	"math/rand/v2"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RandomIndices()( int, int, int, int, int ){
	randNum1 := rand.IntN(24)
	randNum2 := rand.IntN(24)
	randNum3 := rand.IntN(24)
	randNum4 := rand.IntN(24)
	randNum5 := rand.IntN(24)
	fmt.Println(randNum1)
	return randNum1, randNum2, randNum3, randNum4, randNum5
}