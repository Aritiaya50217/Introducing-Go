package main

import "fmt"

func main() {
	// map[key]value
	x := make(map[string]int)
	x["ten"] = 10
	fmt.Println(x["ten"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	// กรณีไม่มีข้อมูลอยู่จะคืนค่า zero value
	fmt.Printf("%#v\n", elements["Un"])

	// check value
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	// map ซ้อน map
	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements2["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}
