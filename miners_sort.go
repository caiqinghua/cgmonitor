package main

import (
	"sort"
	"math/rand"
)
// define a function type which accepts two miner statistics and compare by different criterion
type By func(p1, p2 *MinerRow)	bool

type MinersSorter struct{
	rows []MinerRow
	by	By
}

func (by By) Sort(rows []MinerRow) {
	ms := &MinersSorter{
		rows: rows,
		by: by,
	}
	sort.Sort(ms)
}

func (s *MinersSorter) Len() int{
	return len(s.rows)
}

func (s *MinersSorter) Swap(i, j int){
	s.rows[i], s.rows[j] = s.rows[j], s.rows[i]
}

func (s *MinersSorter) Less(i, j int) bool{
	return s.by(&s.rows[i], &s.rows[j])
}

func Name(r1, r2 *MinerRow) bool{
	return r1.Name < r2.Name
}
func Elapsed(r1, r2 *MinerRow) bool{
	return r1.Elapsed < r2.Elapsed
}
func Accepted(r1, r2 *MinerRow) bool{
	return r1.Accepted < r2.Accepted
}
func Rejected(r1, r2 *MinerRow) bool{
	return r1.Rejected < r2.Rejected
}
func MHSAv(r1, r2 *MinerRow) bool{
	return r1.MHSAv > r2.MHSAv
}
func BestShare(r1, r2 *MinerRow) bool{
	return r1.BestShare < r2.BestShare
}
func IP(r1, r2 *MinerRow) bool{
	return r1.IP < r2.IP
}

// The caller shall guarantee that the modification is safe for concurrency
func sortRows(rows []MinerRow, criterion string)[]MinerRow{
	switch criterion{
	case "Name": By(Name).Sort(rows)
	case "Elapsed": By(Elapsed).Sort(rows)
	case "Accepted": By(Accepted).Sort(rows)
	case "Rejected": By(Rejected).Sort(rows)
	case "MHSAv": By(MHSAv).Sort(rows)
	case "BestShare": By(BestShare).Sort(rows)
	case "IP": By(IP).Sort(rows)
	default : By(Name).Sort(rows)
	}
	return rows
}

func initRowsRand(rows []MinerRow){
	rows_len := len(rows)
	for i:=0; i<rows_len; i++{
		rows[i].Accepted = rand.Intn(5000);
		rows[i].Rejected = rand.Intn(100);
		rows[i].MHSAv = 1000*rand.Float64();
		rows[i].BestShare = rand.Intn(20);
	}
}
