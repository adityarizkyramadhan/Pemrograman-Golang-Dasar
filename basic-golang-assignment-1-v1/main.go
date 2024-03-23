package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"

type StudentObject struct {
	ID    string
	Name  string
	Major string
}

var StudentList []StudentObject

var firstRunStudent bool = true

func parseStudentList() {
	if !firstRunStudent {
		return
	}
	studentDatas := strings.Split(Students, ", ")
	for _, student := range studentDatas {
		// Split student data into ID, Name, and Major
		studentInfo := strings.Split(student, "_")

		// Create a new StudentObject
		newStudent := StudentObject{
			ID:    studentInfo[0],
			Name:  studentInfo[1],
			Major: studentInfo[2],
		}

		// Append the new StudentObject to StudentList
		StudentList = append(StudentList, newStudent)
	}
	firstRunStudent = false
}

func Login(id string, name string) string {
	parseStudentList()
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}
	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}
	var student *StudentObject
	for _, s := range StudentList {
		if s.ID == id && s.Name == name {
			student = &s
			break
		}
	}
	if student == nil {
		return "Login gagal: data mahasiswa tidak ditemukan"
	}
	return fmt.Sprintf("Login berhasil: %s (%s)", student.Name, student.Major)
}

func Register(id string, name string, major string) string {
	parseStudentList()
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}
	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}
	// Cari id di StudentList apa sudah ada
	for _, s := range StudentList {
		if s.ID == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}
	student := &StudentObject{
		ID:    id,
		Name:  name,
		Major: major,
	}
	StudentList = append(StudentList, *student)
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", student.Name, student.Major)
}

var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

var firstRunStudyProgram bool = true

type StudyProgramObject struct {
	Code string
	Name string
}

var StudyProgramList []StudyProgramObject

func parseStudyProgramList() {
	if !firstRunStudyProgram {
		return
	}
	programDatas := strings.Split(StudentStudyPrograms, ", ")
	for _, program := range programDatas {
		programInfo := strings.Split(program, "_")
		newProgram := StudyProgramObject{
			Code: programInfo[0],
			Name: programInfo[1],
		}
		StudyProgramList = append(StudyProgramList, newProgram)
	}
	firstRunStudyProgram = false
}

func GetStudyProgram(code string) string {
	parseStudyProgramList()
	if code == "" {
		return "Code is undefined!"
	}
	var program *StudyProgramObject
	for _, p := range StudyProgramList {
		if p.Code == code {
			program = &p
			break
		}
	}
	if program == nil {
		return ""
	}
	return strings.Trim(program.Name, " ")
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
