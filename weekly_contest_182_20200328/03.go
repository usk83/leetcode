// 5370. Design Underground System
// https://leetcode.com/contest/weekly-contest-182/problems/design-underground-system/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

type UndergroundSystem struct {
	inStations  map[int]string
	inTimes     map[int]int
	totalTimes  map[[2]string]float64
	totalCounts map[[2]string]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		inStations:  map[int]string{},
		inTimes:     map[int]int{},
		totalTimes:  map[[2]string]float64{},
		totalCounts: map[[2]string]int{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.inStations[id] = stationName
	this.inTimes[id] = t
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	from := this.inStations[id]
	to := stationName
	if from > to {
		from, to = to, from
	}
	pair := [2]string{from, to}
	this.totalTimes[pair] += float64(t - this.inTimes[id])
	this.totalCounts[pair]++
	delete(this.inStations, id)
	delete(this.inTimes, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	if startStation > endStation {
		startStation, endStation = endStation, startStation
	}
	pair := [2]string{startStation, endStation}
	return this.totalTimes[pair] / float64(this.totalCounts[pair])
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
