package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var req = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func main() {
	b, err := ioutil.ReadFile("4.in")
	if err != nil {
		log.Fatal(err)
	}

	ps := strings.Split(string(b), "\n\n")

	count := 0
	for _, p := range ps {
		fields := map[string]bool{}
		for _, s := range strings.Split(strings.Replace(p, "\n", " ", -1), " ") {
			if s == "" {
				continue
			}
			ss := strings.Split(s, ":")
			key, value := ss[0], ss[1]
			if os.Getenv("VALID") == "1" {
				if valid(key)(value) {
					fmt.Println(key, value, "is valid")
					fields[key] = true
				} else {
					fmt.Println(key, value, "is invalid")
				}
			} else {
				fields[key] = true
			}
		}

		missingFields := false
		for _, r := range req {
			if !fields[r] {
				missingFields = true
				break
			}
		}

		if !missingFields {
			count++
		}
	}

	fmt.Println(count)
}

func valid(key string) func(string) bool {
	switch key {
	case "byr":
		return func(s string) bool {
			if strings.HasPrefix(s, "0") {
				log.Fatal(s + "has prefix 0")
			}

			match, err := regexp.MatchString("[0-9]{4}", s)
			if err != nil {
				log.Fatal(err)
			}
			if !match {
				return false
			}

			if i, err := strconv.Atoi(s); err == nil {
				if i >= 1920 && i <= 2002 {
					return true
				}
			}
			return false
		}
	case "iyr":
		return func(s string) bool {
			if strings.HasPrefix(s, "0") {
				log.Fatal(s + "has prefix 0")
			}

			match, err := regexp.MatchString("[0-9]{4}", s)
			if err != nil {
				log.Fatal(err)
			}
			if !match {
				return false
			}

			if i, err := strconv.Atoi(s); err == nil {
				if i >= 2010 && i <= 2020 {
					return true
				}
			}
			return false
		}
	case "eyr":
		return func(s string) bool {
			if strings.HasPrefix(s, "0") {
				log.Fatal(s + "has prefix 0")
			}

			match, err := regexp.MatchString("[0-9]{4}", s)
			if err != nil {
				log.Fatal(err)
			}
			if !match {
				return false
			}

			if i, err := strconv.Atoi(s); err == nil {
				if i >= 2020 && i <= 2030 {
					return true
				}
			}
			return false
		}
	case "hgt":
		return func(s string) bool {
			if strings.HasSuffix(s, "cm") {
				if i, err := strconv.Atoi(strings.TrimSuffix(s, "cm")); err == nil {
					if i >= 150 && i <= 193 {
						return true
					}
				}
			} else if strings.HasSuffix(s, "in") {
				if i, err := strconv.Atoi(strings.TrimSuffix(s, "in")); err == nil {
					if i >= 59 && i <= 76 {
						return true
					}
				}
			}
			return false
		}
	case "hcl":
		return func(s string) bool {
			match, err := regexp.MatchString("^#[0-9a-f]{6}$", s)
			if err != nil {
				log.Fatal(err)
			}
			return match
		}
	case "ecl":
		return func(s string) bool {
			var valid = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			for _, v := range valid {
				if s == v {
					return true
				}
			}
			return false
		}
	case "pid":
		return func(s string) bool {
			match, err := regexp.MatchString("^[0-9]{9}$", s)
			if err != nil {
				log.Fatal(err)
			}
			return match
		}
	}

	return func(string) bool { return false }
}
