package structs

import (
	"fmt"
	"strings"
)

// 1) Student and a function to print details
// ---------------------------------------------------------------------
type Student struct {
	Name  string
	Age   int
	Grade string
}

// PrintStudentDetails prints a student's details in a readable format.
func PrintStudentDetails() {
	s := Student{
		Name:  "Alice",
		Age:   20,
		Grade: "A",
	}
	fmt.Printf("Student: Name=%s, Age=%d, Grade=%s\n", s.Name, s.Age, s.Grade)
}

// 2) Book with a Display() method
// ---------------------------------------------------------------------
type Book struct {
	Title  string
	Author string
	Price  float64
}

func NewBook(title, author string, price float64) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Price:  price,
	}
}

// Display prints the book details.
func (b *Book) Display() {
	fmt.Printf("Book: \"%s\" by %s (%.2f)\n", b.Title, b.Author, b.Price)
}

// 3) Rectangle with Area() float64
// ---------------------------------------------------------------------
type Rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

// Area returns the area of the rectangle.
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// ExampleRectangleArea creates a rectangle and returns its computed area.
// Useful for demonstrations/tests without printing to stdout.
func ExampleRectangleArea(width, height float64) {
	r := NewRectangle(width, height)
	area := r.Area()
	fmt.Printf("Rectangle Area (%.2fx%.2f) = %.2f\n", width, height, area)
}

// 4) Movie slice utilities and printing titles after a given year
// ---------------------------------------------------------------------
type Movie struct {
	Title string
	Year  int
}

func NewMovie(title string, year int) *Movie {
	return &Movie{
		Title: title,
		Year:  year,
	}
}

// PrintMovieTitlesAfter prints titles of movies released after the given year.
func PrintMovieTitlesAfter(movies []Movie, year int) {
	after := MoviesAfter(movies, year)
	if len(after) == 0 {
		fmt.Printf("No movies found after %d\n", year)
		return
	}
	titles := make([]string, 0, len(after))
	for _, m := range after {
		titles = append(titles, m.Title)
	}
	fmt.Printf("Movies after %d: %s\n", year, strings.Join(titles, ", "))
}

// MoviesAfter filters movies released after the given year (exclusive).
func MoviesAfter(movies []Movie, year int) []Movie {
	res := make([]Movie, 0, len(movies))
	for _, m := range movies {
		if m.Year > year {
			res = append(res, m)
		}
	}
	return res
}

func ExamplePrintMovieTitlesAfter() {
	movie1 := NewMovie("Inception", 2010)
	movie2 := NewMovie("Interstellar", 2014)
	movie3 := NewMovie("The Dark Knight", 2008)
	movie4 := NewMovie("Avengers: Endgame", 2019)

	movies := []Movie{*movie1, *movie2, *movie3, *movie4}
	PrintMovieTitlesAfter(movies, 2010)
}

// 5) Address and Person, printing nested details
// ---------------------------------------------------------------------
type Address struct {
	City string
	Zip  int
}

type Person struct {
	Name    string
	Address Address
}

func NewPerson(name, city string, zip int) *Person {
	return &Person{
		Name: name,
		Address: Address{
			City: city,
			Zip:  zip,
		},
	}
}

// PrintPerson prints the person's name and address details.
func (p *Person) PrintPerson() {
	fmt.Printf("Person: %s\n", p.Name)
	fmt.Printf("Address: %s, %d\n", p.Address.City, p.Address.Zip)
}

func DisplayPerson() {
	p := NewPerson("John Doe", "New York", 10001)
	p.PrintPerson()
}

func StructExamples() {
	PrintStudentDetails()
	b := NewBook("1984", "George Orwell", 9.99)
	b.Display()
	ExampleRectangleArea(5.0, 3.0)
	ExamplePrintMovieTitlesAfter()
	DisplayPerson()
}
