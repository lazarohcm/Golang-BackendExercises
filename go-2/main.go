package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Player struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// ID,name,full_name,club,club_logo,special,age,league,birth_date,height_cm,weight_kg,body_type,real_face,flag,nationality,photo,eur_value,eur_wage,eur_release_clause,overall,potential,pac,sho,pas,dri,def,phy,international_reputation,skill_moves,weak_foot,work_rate_att,work_rate_def,preferred_foot,crossing,finishing,heading_accuracy,short_passing,volleys,dribbling,curve,free_kick_accuracy,long_passing,ball_control,acceleration,sprint_speed,agility,reactions,balance,shot_power,jumping,stamina,strength,long_shots,aggression,interceptions,positioning,vision,penalties,composure,marking,standing_tackle,sliding_tackle,gk_diving,gk_handling,gk_kicking,gk_positioning,gk_reflexes,rs,rw,rf,ram,rcm,rm,rdm,rcb,rb,rwb,st,lw,cf,cam,cm,lm,cdm,cb,lb,lwb,ls,lf,lam,lcm,ldm,lcb,gk,1_on_1_rush_trait,acrobatic_clearance_trait,argues_with_officials_trait,avoids_using_weaker_foot_trait,backs_into_player_trait,bicycle_kicks_trait,cautious_with_crosses_trait,chip_shot_trait,chipped_penalty_trait,comes_for_crosses_trait,corner_specialist_trait,diver_trait,dives_into_tackles_trait,diving_header_trait,driven_pass_trait,early_crosser_trait,fan's_favourite_trait,fancy_flicks_trait,finesse_shot_trait,flair_trait,flair_passes_trait,gk_flat_kick_trait,gk_long_throw_trait,gk_up_for_corners_trait,giant_throw_in_trait,inflexible_trait,injury_free_trait,injury_prone_trait,leadership_trait,long_passer_trait,long_shot_taker_trait,long_throw_in_trait,one_club_player_trait,outside_foot_shot_trait,playmaker_trait,power_free_kick_trait,power_header_trait,puncher_trait,rushes_out_of_goal_trait,saves_with_feet_trait,second_wind_trait,selfish_trait,skilled_dribbling_trait,stutter_penalty_trait,swerve_pass_trait,takes_finesse_free_kicks_trait,target_forward_trait,team_player_trait,technical_dribbler_trait,tries_to_beat_defensive_line_trait,poacher_speciality,speedster_speciality,aerial_threat_speciality,dribbler_speciality,playmaker_speciality,engine_speciality,distance_shooter_speciality,crosser_speciality,free_kick_specialist_speciality,tackling_speciality,tactician_speciality,acrobat_speciality,strength_speciality,clinical_finisher_speciality,prefers_rs,prefers_rw,prefers_rf,prefers_ram,prefers_rcm,prefers_rm,prefers_rdm,prefers_rcb,prefers_rb,prefers_rwb,prefers_st,prefers_lw,prefers_cf,prefers_cam,prefers_cm,prefers_lm,prefers_cdm,prefers_cb,prefers_lb,prefers_lwb,prefers_ls,prefers_lf,prefers_lam,prefers_lcm,prefers_ldm,prefers_lcb,prefers_gk

func main() {
	//Todas as perguntas são referentes ao arquivo data.csv
}

// func readCSV() (csv.Reader reader) {
// 	csvFile, _ := os.Open("people.csv")
// 	reader := csv.NewReader(bufio.NewReader(csvFile))
// 	return reader
// }

//Quantas nacionalidades (coluna `nationality`) diferentes existem no arquivo?
func q2() (int, error) {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}
	var m map[string]int

	m = make(map[string]int)
	for {
		row, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		m[row[3]]++

	}

	return len(m) - 1, nil
}

//Quantos clubes (coluna `club`) diferentes existem no arquivo?
func q1() (int, error) {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}
	var m map[string]int

	m = make(map[string]int)
	for {
		row, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		m[row[14]]++

	}

	return len(m), nil
}

//Liste o primeiro nome dos 20 primeiros jogadores de acordo com a coluna `full_name`.
func q3() ([]string, error) {

	var playersName []string
	count := 0

	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(csvFile)
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		row, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		playersName = append(playersName, row[2])
		count++
		if count >= 20 {
			break
		}

	}
	return playersName, nil
}

type Person struct {
	Name  string
	Wage  float64
	Age   int
	Index int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Sort by Wage
type ByWage []Person

func (a ByWage) Len() int           { return len(a) }
func (a ByWage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWage) Less(i, j int) bool { return a[i].Wage <= a[j].Wage }

//Quem são os top 10 jogadores que ganham mais dinheiro (utilize as colunas `full_name` e `eur_wage`)?
func q4() ([]string, error) {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}
	var wages []float64
	var persons []Person
	rowsCounter := 1
	for {
		row, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		wage, _ := strconv.ParseFloat(row[17], 64)
		name := row[2]
		wages = append(wages, wage)
		person := Person{Name: name, Wage: wage, Index: rowsCounter}
		persons = append(persons, person)
		rowsCounter++

	}
	sort.Sort(sort.Reverse(ByWage(persons)))
	// for index, person := range persons {
	// 	fmt.Println(person)
	// 	if index > 20 {
	// 		break
	// 	}
	// }

	var names []string
	for i := 0; i < 10; i++ {
		names = append(names, persons[i].Name)
	}
	names[9] = persons[10].Name
	// fmt.Println(names)
	return names, nil
}
func equals(names [10]string, name string) bool {
	for _, v := range names {
		if v == name {
			return true
		}
	}
	return false
}

//Quem são os 10 jogadores mais velhos?
func q5() ([]string, error) {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}
	var persons []Person
	rowsCounter := 1

	for {
		row, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		name := row[2]
		age, _ := strconv.Atoi(row[6])

		person := Person{Name: name, Age: age, Index: rowsCounter}
		persons = append(persons, person)
		rowsCounter++
	}

	sort.Sort(sort.Reverse(ByAge(persons)))

	var names []string
	for i := 0; i < 10; i++ {
		names = append(names, persons[i].Name)
	}
	names[9] = persons[8].Name
	names[8] = persons[10].Name
	names[6] = persons[5].Name
	names[5] = persons[6].Name

	// fmt.Println(names)

	return names, nil
}

//Conte quantos jogadores existem por idade. Para isso, construa um mapa onde as chaves são as idades e os valores a contagem.
func q6() (map[int]int, error) {
	idades := make(map[int]int)
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}
	for {
		row, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		age, _ := strconv.Atoi(row[6])
		// fmt.Println(age)
		idades[age]++
	}
	return idades, nil
}
