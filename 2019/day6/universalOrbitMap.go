package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var reOrbitDefinition = regexp.MustCompile(`^(\w+)\)(\w+)$`)

type universalOrbitMap struct {
	orbitData      map[string]string
	orbitCountData map[string]int
}

func (u *universalOrbitMap) addOrbit(orbiter string, orbitee string) error {
	if v, ok := u.orbitData[orbiter]; ok && v != orbitee {
		return fmt.Errorf("orbiter %q already has a defined orbitee", orbiter)
	}

	u.orbitData[orbiter] = orbitee
	return nil
}

func (u *universalOrbitMap) checkCOM() error {
	for _, v := range u.orbitData {
		if v == "COM" {
			return nil
		}
	}

	return fmt.Errorf("%q orbitee is not defined", "COM")
}

func (u *universalOrbitMap) generateOrbitCountData() error {
	if err := u.checkCOM(); err != nil {
		return err
	}

	u.orbitCountData = make(map[string]int, len(u.orbitData)+1)
	u.orbitCountData["COM"] = 0

	for orbiter, _ := range u.orbitData {
		if _, err := u.calculateOrbitCount(orbiter); err != nil {
			return err
		}
	}

	return nil
}

func (u *universalOrbitMap) calculateOrbitCount(orbiter string) (int, error) {
	orbitee, exists := u.orbitData[orbiter]
	if !exists {
		return 0, fmt.Errorf("orbiter %q has no defined orbitee", orbiter)
	}

	if v, ok := u.orbitCountData[orbitee]; ok {
		orbitCount := v + 1
		u.orbitCountData[orbiter] = orbitCount
		return orbitCount, nil
	}

	v, err := u.calculateOrbitCount(orbitee)
	if err != nil {
		return 0, err
	}

	orbitCount := v + 1
	u.orbitCountData[orbiter] = orbitCount
	return orbitCount, nil
}

func newUniversalOrbitMap() *universalOrbitMap {
	return &universalOrbitMap{
		orbitData: make(map[string]string),
	}
}

func NewUniversalOrbitMapFromFile(inputFile string) (*universalOrbitMap, error) {
	uom := newUniversalOrbitMap()

	fileHandle, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileLineNr := 1; fileScanner.Scan(); fileLineNr++ {
		fileLine := fileScanner.Text()

		if reOrbitDefinition.MatchString(fileLine) {
			matches := reOrbitDefinition.FindStringSubmatch(fileLine)
			err := uom.addOrbit(matches[2], matches[1])
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("unexpected syntax near line %d of file %q", fileLineNr, inputFile)
		}
	}

	if err := uom.generateOrbitCountData(); err != nil {
		return nil, err
	}

	return uom, nil
}

func (u *universalOrbitMap) CheckSum() int {
	checkSum := 0
	for k, _ := range u.orbitData {
		orbitCount, _ := u.orbitCountData[k]
		checkSum += orbitCount
	}

	return checkSum
}

func (u *universalOrbitMap) GetOrbiteeList(orbiter string) *orbiteeList {
	items := orbiteeList{}

	for item := orbiter; item != "COM"; item, _ = u.orbitData[item] {
		items.Add(item)
	}

	fmt.Printf("[DEBUG]: universalOrbitMap::GetOrbiteeList: items=%+v\n", items)

	return &items
}

type orbiteeList struct {
	data []string
}

func (u *orbiteeList) Add(orbitee string) {
	u.data = append(u.data, orbitee)
}

func (u *orbiteeList) GetIndexByName(orbitee string) int {
	for i, v := range u.data {
		if v == orbitee {
			return i
		}
	}

	return -1
}

func (u *orbiteeList) CalculateOrbitalDistance(v *orbiteeList) int {
	minOrbitalDistance := len(u.data) + len(v.data)

	for i := range u.data {
		if vi := v.GetIndexByName(u.data[i]); vi > 0 {
			orbitalDistance := i + vi - 2
			// fmt.Printf("[DEBUG]: orbiteeList::CalculateDivergencePoint: commonOrbitee=%s i=%d vi=%d orbitalDistance=%d\n", u.data[i], i, vi, orbitalDistance)

			if orbitalDistance < minOrbitalDistance {
				minOrbitalDistance = orbitalDistance
			}
		}
	}

	return minOrbitalDistance
}
