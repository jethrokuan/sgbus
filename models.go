package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type busService struct {
	BusNumber int
	Status    int
	NextBus   []time.Time
}

type busStop struct {
	BusStopID int
	Services  []busService
}

func (s *busStop) UnmarshalJSON(data []byte) error {
	type Alias busStop
	aux := &struct {
		BusStopID string `json:"BusStopID"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.BusStopID, _ = strconv.Atoi(aux.BusStopID)
	return nil
}

func (s *busService) UnmarshalJSON(data []byte) error {
	type Alias busService
	type ba struct {
		EstimatedArrival string `json:"EstimatedArrival"`
	}
	aux := &struct {
		ServiceNumber  string `json:"ServiceNo"`
		Status         string `json:"Status"`
		NextBus        ba     `json:"NextBus"`
		SubsequentBus  ba     `json:"SubsequentBus"`
		SubsequentBus3 ba     `json:"SubsequentBus3"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.BusNumber, _ = strconv.Atoi(aux.ServiceNumber)
	if aux.Status == "In Operation" {
		s.Status = 1
	} else {
		s.Status = 0
	}

	a, err := time.Parse(time.RFC3339, aux.NextBus.EstimatedArrival)
	if err != nil {
		log.Printf("Error! Failed to parse NextBus time: %s\n", err)
		return err
	}
	b, err := time.Parse(time.RFC3339, aux.SubsequentBus.EstimatedArrival)
	if err != nil {
		log.Printf("Error! Failed to parse SubsequentBus time: %s\n", err)
		return err
	}
	c, err := time.Parse(time.RFC3339, aux.SubsequentBus3.EstimatedArrival)
	if err != nil {
		log.Printf("Error! Failed to parse SubsequentBus3 time: %s\n", err)
		return err
	}

	s.NextBus = []time.Time{a, b, c}
	return nil
}
