package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SubjectScore struct {
	testid  string
	Subject string
	Score   int
	Date    string
}

type Student struct {
	ID            string
	Name          string
	SubjectScores []SubjectScore
}

var students []Student

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("請問你的身分 老師(T) / 學生(S): ")
		identity, _ := reader.ReadString('\n')
		identity = strings.TrimSpace(identity)

		switch identity {
		case "T":
			showTeacherMenu(reader)
		case "S":
			showStudentMenu(reader)
		default:
			fmt.Println("無效的身份。請重新輸入。")
		}
	}
}

func showTeacherMenu(reader *bufio.Reader) {
	for {
		fmt.Println("\n請選擇操作:")
		fmt.Println("(a) 創建學生")
		fmt.Println("(b) 刪除學生")
		fmt.Println("(c) 輸入成績")
		fmt.Println("(d) 印出學生清單")
		fmt.Println("(e) 印出全班成績")
		fmt.Println("(f) 退出")
		fmt.Print("選擇操作: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "a":
			students = append(students, createStudent(reader))
		case "b":
			deleteStudent(reader)
		case "c":
			inputScore(reader)
		case "d":
			printAllStudents()
		case "e":
			printAllScores(reader)
		case "f":
			fmt.Println("返回到輸入身分。")
			return
		default:
			fmt.Println("無效的選擇。請重新輸入。")
		}
	}
}

func showStudentMenu(reader *bufio.Reader) {
	fmt.Print("請輸入你的學生編號: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	var student *Student
	for _, s := range students {
		if s.ID == id {
			student = &s
			break
		}
	}

	if student == nil {
		fmt.Println("沒有找到此學生。")
		return
	}

	for {
		fmt.Println("\n請選擇操作:")
		fmt.Println("(a) 印出所有成績")
		fmt.Println("(b) 返回")
		fmt.Print("選擇操作: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "a":
			printStudentScores(*student)
		case "b":
			fmt.Println("返回到輸入身分。")
			return
		default:
			fmt.Println("無效的選擇。請重新輸入。")
		}
	}
}

func printStudentScores(student Student) {
	fmt.Println("學生名字: ", student.Name)
	for _, subjectScore := range student.SubjectScores {
		fmt.Println("科目: ", subjectScore.Subject, " 分數: ", subjectScore.Score)
	}
	fmt.Println()
}

func checkStudentIDExists(id string) bool {
	for _, s := range students {
		if s.ID == id {
			return true
		}
	}
	return false
}

func createStudent(reader *bufio.Reader) Student {
	var id, name string

	for {
		fmt.Print("請輸入學生的編號: ")
		id, _ = reader.ReadString('\n')
		id = strings.TrimSpace(id)

		if !checkStudentIDExists(id) {
			break
		}

		fmt.Println("該學生編號已經存在，請重新輸入。")
	}

	fmt.Print("請輸入學生的名字: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return Student{ID: id, Name: name}
}

func deleteStudent(reader *bufio.Reader) {
	fmt.Print("請輸入要刪除的學生編號: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	for i, student := range students {
		if student.ID == id {
			// Delete the student from the slice.
			students = append(students[:i], students[i+1:]...)
			fmt.Println("學生已刪除。")
			return
		}
	}

	fmt.Println("找不到該學生。")
}

func createSubjectScore(reader *bufio.Reader, testid, subject string) SubjectScore {
	fmt.Printf("請輸入 %s 的分數: ", subject)
	scoreStr, _ := reader.ReadString('\n')
	scoreStr = strings.TrimSpace(scoreStr)
	score, _ := strconv.Atoi(scoreStr)

	return SubjectScore{testid: testid, Subject: subject, Score: score, Date: testid}
}

func inputScore(reader *bufio.Reader) {
	fmt.Print("請輸入考試編號: ")
	testid, _ := reader.ReadString('\n')
	testid = strings.TrimSpace(testid)

	fmt.Print("請輸入學生編號: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	// find student
	var targetStudent *Student
	for i, s := range students {
		if s.ID == id {
			targetStudent = &students[i]
			break
		}
	}

	if targetStudent == nil {
		fmt.Println("此學生不存在。")
		return
	}

	subjects := []string{"數學", "國文", "英文"}
	for _, subject := range subjects {
		targetStudent.SubjectScores = append(targetStudent.SubjectScores, createSubjectScore(reader, testid, subject))
	}
}

func printAllStudents() {
	fmt.Println("+----------+--------+")
	fmt.Println("| 學生編號 |   姓名 |")
	fmt.Println("+----------+--------+")
	for _, student := range students {
		fmt.Printf("| %8s | %6s |\n", student.ID, student.Name)
	}
	fmt.Println("+----------+--------+")
}

func printAllScores(reader *bufio.Reader) {
	fmt.Print("請輸入考試編號: ")
	testid, _ := reader.ReadString('\n')
	testid = strings.TrimSpace(testid)

	for _, student := range students {
		fmt.Println("學生名字: ", student.Name)
		hasScore := false
		for _, subjectScore := range student.SubjectScores {
			if subjectScore.testid == testid {
				fmt.Println("科目: ", subjectScore.Subject, " 分數: ", subjectScore.Score)
				hasScore = true
			}
		}
		if !hasScore {
			fmt.Println("該學生沒有這個考試的成績。")
		}
		fmt.Println()
	}
}
