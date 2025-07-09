package main

import (
	"errors"
	"fmt"
	"sync"
	"log"
	"context"
)

// Custom errors
var (
	ErrNotImplemented = errors.New("Not implemented")
	ErrTruckNotFound  = errors.New("Truck not found")
)

// Truck interface
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

// NormalTruck struct
type NormalTruck struct {
	ID    string
	Cargo int
}

// LoadCargo implementation for NormalTruck
func (t *NormalTruck) LoadCargo() error {
	t.Cargo += 1
	return nil
}

// UnloadCargo implementation for NormalTruck
func (t *NormalTruck) UnloadCargo() error {
	t.Cargo = 0
	return nil
}

// ElectricTruck struct
type ElectricTruck struct {
	ID      string
	Cargo   int
	Battery float64
}

// LoadCargo implementation for ElectricTruck
func (e *ElectricTruck) LoadCargo() error {
	e.Cargo += 1
	e.Battery -= 1
	return nil
}

// UnloadCargo implementation for ElectricTruck
func (e *ElectricTruck) UnloadCargo() error {
	e.Cargo = 0
	e.Battery -= 1
	return nil
}

// Process a single truck
func processTruck(ctx context.Context, truck Truck) error {
	// More advanced
	fmt.Printf("Started processing truck %+v \n", truck)

	// access the user id
	userID := ctx.Value("userID")


	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	fmt.Printf("Finished processing truck %+v \n", truck)
	return nil
}

// Process a fleet of trucks
func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1)
		go func(t Truck) {
			if err := processTruck(ctx, t); err != nil{
				log.Println(err)
			}
			wg.Done()
			}(t)
	}
	
	wg.Wait()

	return nil
}

// Main function
func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", 42)

	fleet := []Truck{
		&NormalTruck{ID: "NT1", Cargo: 0},
		&ElectricTruck{ID: "ET1", Cargo: 0, Battery: 100},
		&NormalTruck{ID: "NT2", Cargo: 0},
		&ElectricTruck{ID: "ET2", Cargo: 0, Battery: 100},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Println("Fleet processed successfully.")
}
