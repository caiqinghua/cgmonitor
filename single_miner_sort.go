package main

import (
	"fmt"
	"sort"
	"math/rand"
)
// define a function type which accepts two miner statistics and compare by different criterion
type DevsBy func(p1, p2 *DevObject)	bool

type DevSorter struct{
	devs []DevObject
	by	DevsBy
}

func (by DevsBy) Sort(devs []DevObject) {
	ds := &DevSorter{
		devs: devs,
		by: by,
	}
	sort.Sort(ds)
}

func (s *DevSorter) Len() int{
	return len(s.devs)
}

func (s *DevSorter) Swap(i, j int){
	s.devs[i], s.devs[j] = s.devs[j], s.devs[i]
}

func (s *DevSorter) Less(i, j int) bool{
	return s.by(&s.devs[i], &s.devs[j])
}

func DevName(r1, r2 *DevObject) bool{
	return r1.GPU < r2.GPU
}
func ID(r1, r2 *DevObject) bool{
	return r1.ID < r2.ID
}
func Enabled(r1, r2 *DevObject) bool{
	return r1.Enabled < r2.Enabled
}
func Status(r1, r2 *DevObject) bool{
	return r1.Status < r2.Status
}
func DevAccepted(r1, r2 *DevObject) bool{
	return r1.Accepted < r2.Accepted
}
func DevRejected(r1, r2 *DevObject) bool{
	return r1.Rejected < r2.Rejected
}
func DevMHSAv(r1, r2 *DevObject) bool{
	return r1.MHSAv < r2.MHSAv
}
func DevMHS5s(r1, r2 *DevObject) bool{
	return r1.MHS5s < r2.MHS5s
}
func HardwareErrors(r1, r2 *DevObject) bool{
	return r1.HardwareErrors < r2.HardwareErrors
}
func Temperature(r1, r2 *DevObject) bool{
	return r1.Temperature < r2.Temperature
}
func FanSpeed(r1, r2 *DevObject) bool{
	return r1.FanSpeedIn < r2.FanSpeedIn
}

// The caller shall guarantee that the modification is safe for concurrency
func sortDevs(devs []DevObject, criterion string) []DevObject{
	switch criterion{
	case "Name": DevsBy(DevName).Sort(devs)
	case "ID": DevsBy(ID).Sort(devs)
	case "Enabled": DevsBy(Enabled).Sort(devs)
	case "Status": DevsBy(Status).Sort(devs)
	case "Accepted": DevsBy(DevAccepted).Sort(devs)
	case "Rejected": DevsBy(DevRejected).Sort(devs)
	case "MHSAv": DevsBy(DevMHSAv).Sort(devs)
	case "MHS5s": DevsBy(DevMHS5s).Sort(devs)
	case "HardwareErrors": DevsBy(HardwareErrors).Sort(devs)
	case "Temperature": DevsBy(Temperature).Sort(devs)
	case "FanSpeed": DevsBy(FanSpeed).Sort(devs)
	default : DevsBy(DevName).Sort(devs)
	}
	return devs
}

func initDevsRand(devs []DevObject) []DevObject{
	temp := make([]DevObject, 10)
	devs = append(devs, temp...)
	devs_len := len(devs)
	for i:=0; i<devs_len; i++{
		devs[i].GPU = fmt.Sprintf("GPU%03d", i)
		devs[i].ID = rand.Intn(1000);
		if (0==rand.Intn(1000)&0x1){
			devs[i].Enabled = fmt.Sprint("Enabled")
		}else{
			devs[i].Enabled = fmt.Sprint("Disabled")
		}
		if (0==rand.Intn(1000)&0x1){
			devs[i].Status = fmt.Sprintf("Live")
		}	else{
			devs[i].Status = fmt.Sprintf("Dead")
		}
			
		devs[i].Accepted = rand.Intn(5000);
		devs[i].Rejected = rand.Intn(200);
		devs[i].MHSAv = 1000*rand.Float64();
		devs[i].MHS5s = 1000*rand.Float64();
		devs[i].HardwareErrors = rand.Intn(20);
		devs[i].Temperature = float64(rand.Intn(100));
		devs[i].FanSpeedIn = rand.Intn(3000);
	}
	return devs
}
