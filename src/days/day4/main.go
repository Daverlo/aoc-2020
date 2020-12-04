package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	Passport       string
	Country        string
}

func (p *Passport) AddFields(fields map[string]string) {
	for k, v := range fields {
		switch k {
		case "byr":
			p.BirthYear = v
		case "iyr":
			p.IssueYear = v
		case "eyr":
			p.ExpirationYear = v
		case "hgt":
			p.Height = v
		case "hcl":
			p.HairColor = v
		case "ecl":
			p.EyeColor = v
		case "pid":
			p.Passport = v
		case "cid":
			p.Country = v
		}
	}
}

func (p *Passport) IsValid() bool {
	if p.BirthYear == "" {
		return false
	}

	if p.IssueYear == "" {
		return false
	}

	if p.ExpirationYear == "" {
		return false
	}

	if p.Height == "" {
		return false
	}

	if p.HairColor == "" {
		return false
	}

	if p.EyeColor == "" {
		return false
	}

	if p.Passport == "" {
		return false
	}

	// Ignore this ¯\_(ツ)_/¯
	// if p.Country == "" {
	// 	return false
	// }

	return true
}

func (p *Passport) IsValid2() bool {
	if p.BirthYear == "" {
		return false
	} else {
		byr, err := strconv.Atoi(p.BirthYear)
		if err != nil {
			return false
		}
		if byr < 1920 || byr > 2002 {
			return false
		}
	}

	if p.IssueYear == "" {
		return false
	} else {
		iyr, err := strconv.Atoi(p.IssueYear)
		if err != nil {
			return false
		}

		if iyr < 2010 || iyr > 2020 {
			return false
		}
	}

	if p.ExpirationYear == "" {
		return false
	} else {
		eyr, err := strconv.Atoi(p.ExpirationYear)
		if err != nil {
			return false
		}

		if eyr < 2020 || eyr > 2030 {
			return false
		}
	}

	if p.Height == "" {
		return false
	} else {
		height := p.Height[:len(p.Height)-2]
		hgt, err := strconv.Atoi(height)
		if err != nil {
			return false
		}

		unit := p.Height[len(p.Height)-2:]
		if unit == "cm" {
			if hgt < 150 || hgt > 193 {
				return false
			}
		} else if unit == "in" {
			if hgt < 59 || hgt > 76 {
				return false
			}
		} else {
			return false
		}
	}

	if p.HairColor == "" {
		return false
	} else {
		match, err := regexp.MatchString("^#[a-f0-9]{6}$", p.HairColor)
		if err != nil {
			return false
		}
		if !match {
			return false
		}
	}

	if p.EyeColor == "" {
		return false
	} else {
		validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		if !contains(validColors, p.EyeColor) {
			return false
		}
	}

	if p.Passport == "" {
		return false
	} else {
		match, err := regexp.MatchString("^[0-9]{9}$", p.Passport)
		if err != nil {
			return false
		}
		if !match {
			return false
		}
	}

	// Ignore this ¯\_(ツ)_/¯
	// if p.Country == "" {
	// 	return false
	// }

	return true
}

func contains(a []string, k string) bool {
	for _, v := range a {
		if v == k {
			return true
		}
	}

	return false
}

func parseInput(path string) ([]Passport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(f)

	var passports []Passport
	for {
		var passport Passport
		var err error

		for {
			var line string
			line, err = rd.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}
			line = strings.TrimSuffix(line, "\n")

			// Passport delimiter
			if len(line) == 0 {
				break
			}

			fields := parseLine(line)
			passport.AddFields(fields)
		}

		passports = append(passports, passport)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}

	return passports, nil
}

func parseLine(line string) map[string]string {
	fields := make(map[string]string)

	tokens := strings.Split(line, " ")
	for _, token := range tokens {
		field := strings.Split(token, ":")
		fields[field[0]] = field[1]
	}

	return fields
}

func part1(passports []Passport) int {
	validPassports := 0

	for _, p := range passports {
		if p.IsValid() {
			validPassports++
		}
	}

	return validPassports
}

func part2(passports []Passport) int {
	validPassports := 0

	for _, p := range passports {
		if p.IsValid2() {
			validPassports++
		}
	}

	return validPassports
}

func main() {
	args := os.Args[1:]
	passports, err := parseInput(args[0])

	if err != nil {
		panic(err)
	}

	output := part1(passports)
	fmt.Println(output)

	output = part2(passports)
	fmt.Println(output)
}
