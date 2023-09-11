//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwith() Bytes {
	var total Bytes
	for _, amount := range b.amount {
		total += amount
	}
	return total / Bytes(len(b.amount))
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) AverageTemp() Celcius {
	var total Celcius
	for _, temp := range c.temp {
		total += temp
	}
	return total / Celcius(len(c.temp))
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) AverageMemory() Bytes {
	var total Bytes
	for _, amount := range m.amount {
		total += amount
	}
	return total / Bytes(len(m.amount))
}

type Dashboard struct {
	BandwidthUsage
	MemoryUsage
	CpuTemp
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{bandwidth, memory, temp}

	fmt.Printf("Bandwidth: %v\n", dash.AverageBandwith())
	fmt.Printf("Memory: %v\n", dash.AverageMemory())
	fmt.Printf("CPU Temp: %v\n", dash.AverageTemp())
}
