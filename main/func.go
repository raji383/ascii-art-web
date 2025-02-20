package fs

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"text/template"
)

func ExportAsciiArt(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("Text")
	file_format := r.FormValue("Banner") + ".txt"

	content, _ := os.ReadFile("main/" + file_format)

	characters := make(map[rune][]string) // this map holds the runes with their draw
	temp_slice := Splitt(string(content)) // this is a temporary slice used to store the ASCII art for each rune before adding it to the map
	char := 32

	// filling the map
	for _, x := range temp_slice {
		characters[rune(char)] = x
		char++
	}
	clean_input := Split_with_newline(input)

	final := ""
	for i := 0; i < len(clean_input); i++ {
		if clean_input[i] != "\n" {
			final1 := Draw(clean_input[i], characters)
			for i := 0; i < len(final1); i++ {
				final += final1[i] + "\n"
			}
		} else {
			final += "\n"
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=\"ascii-art.txt\"")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(final)))

	w.Write([]byte(final))
}

/*
func name: Splitt

Goal: Split the file content into a slice of slices, where each inner slice represents the ASCII art for
a specific rune (from rune 32 to rune 126).

Return: Slice of slices, [[draw of rune 1], [draw of rune 2], ..., [draw of rune 126]]
Args: The file's content as a string
*/
func Splitt(content string) [][]string {
	slice1 := []string{}
	result := [][]string{}
	line := ""
	// Splitting the Content with \n
	for _, x := range content {
		if x == '\n' && x != '\r' {
			slice1 = append(slice1, line)
			line = ""
		} else if x != '\r' {
			line += string(x)
		}
	}
	if line != "" {
		slice1 = append(slice1, line)
		line = ""
	}

	// Making a slice of slices
	for i := 1; i < len(slice1); i += 8 {
		end := i + 8
		if end > len(slice1) {
			break
		}
		temp := slice1[i:end]
		result = append(result, temp)
		i++
	}

	return result
}

/*
This function splits the input of the user based on newlines.
If there is a newline, the sentence before the newline is appended,
and then `\n` is appended as a string.
*/
func Split_with_newline(input string) []string {
	var words []string
	word := ""
	// fmt.Println(input)

	for _, x := range input {
		if x != '\n' && x != '\r' {
			word += string(x)
		} else {
			// fmt.Print(x)
			// fmt.Println(i)
			if word != "" {
				words = append(words, word)
				// fmt.Println(words)
				word = ""
			} else {
				words = append(words, "\n")
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	if input[len(input)-1] == '\n' {
		// fmt.Println("222222")
		words = append(words, "\n")
	}
	// fmt.Println(words)
	return words
}

/*
This function creates the slice that will be printed,
handling any input (with or without spaces).
*/
func Draw(input string, characters map[rune][]string) []string {
	result := make([]string, 8)
	empty := []string{}
	for _, x := range input {
		if temp, exist := characters[x]; exist {
			for j := 0; j < 8; j++ {
				result[j] = result[j] + temp[j]
			}
		} else {
			return empty
		}
	}
	return result
}

func Css(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/css/" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	// Serve the CSS file
	cssFile := "main/css/" + r.URL.Path[len("/css/"):] // this : r.URL.Path[len("/css/"):] gives me (for example) index.css
	http.ServeFile(w, r, cssFile)
}

/*
It calls the home page
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		dataerror := map[string]string{
			"ErrorCode":    "405",
			"ErrorMessage": "Method Not Allowed",
		}
		w.WriteHeader(405)
		errormesage := template.Must(template.ParseFiles("main/templates/error.html"))
		errormesage.Execute(w, dataerror)
		return
	} else {
		if r.URL.Path == "/" {
			tmpl, err := template.ParseFiles("main/templates/index.html")
			if err != nil {
				dataerror := map[string]string{
					"ErrorCode":    "500",
					"ErrorMessage": "Internal Server Error",
				}
				w.WriteHeader(500)
				errormesage := template.Must(template.ParseFiles("main/templates/error.html"))
				errormesage.Execute(w, dataerror)
				return
			}
			tmpl.Execute(w, nil)
		} else {
			dataerror := map[string]string{
				"ErrorCode":    "404",
				"ErrorMessage": "not funde",
			}
			w.WriteHeader(404)
			errormesage := template.Must(template.ParseFiles("main/templates/error.html"))
			errormesage.Execute(w, dataerror)
			return
		}
	}
}

/*
Draw the final result.

w http.ResponseWriter: This is the HTTP response writer used to send the final output to the client.
r *http.Request: The HTTP request that contains the input parameters.
*/
func Finaldrawing(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("text")
	file_format := r.FormValue("banner") + ".txt"

	if r.Method != http.MethodPost {
		dataerror := map[string]string{
			"ErrorCode":    "405",
			"ErrorMessage": "Method Not Allowed",
		}
		w.WriteHeader(405)
		errormesage := template.Must(template.ParseFiles("main/templates/error.html"))
		errormesage.Execute(w, dataerror)
		return
	}
	// check if the input is empty
	if len(input) == 0 || len(input) >= 3000 {
		dataerror := map[string]string{
			"ErrorCode":    "400",
			"ErrorMessage": "Bad request",
		}
		w.WriteHeader(http.StatusBadRequest)

		tmpl, err := template.ParseFiles("main/templates/error.html")
		if err != nil {
			http.Error(w, "Internal Server Error: Failed to load error template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, dataerror)
		return
	}
	for _, x := range input {
		if x != '\r' && x != '\n' && (x < ' ' || x > '~') {

			dataerror := map[string]string{
				"ErrorCode":    "400",
				"ErrorMessage": "Bad request: Input contains invalid characters.",
			}

			w.WriteHeader(http.StatusBadRequest)

			tmpl, err := template.ParseFiles("main/templates/error.html")
			if err != nil {
				http.Error(w, "Internal Server Error: Failed to load error template", http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, dataerror)
			return
		}
	}

	content, err1 := os.ReadFile("main/" + file_format)
	if err1 != nil {
		dataerror := map[string]string{
			"ErrorCode":    "500",
			"ErrorMessage": "Internal Server Error",
		}
		w.WriteHeader(500)
		errormesage := template.Must(template.ParseFiles("main/templates/error.html"))
		errormesage.Execute(w, dataerror)
		return
	}
	//--------------------------------------------------------------------------------

	characters := make(map[rune][]string) // this map holds the runes with their draw
	temp_slice := Splitt(string(content)) // this is a temporary slice used to store the ASCII art for each rune before adding it to the map
	char := 32
	if len(temp_slice) == 0 {
		dataerror := map[string]string{
			"ErrorCode":    "500",
			"ErrorMessage": "Internal Server Error",
		}
		w.WriteHeader(500)
		errormesage := template.Must(template.ParseFiles("main/templates/error.html"))
		errormesage.Execute(w, dataerror)
		return
	}

	// filling the map
	for _, x := range temp_slice {
		characters[rune(char)] = x
		char++
	}
	clean_input := Split_with_newline(input)
	tmpl := template.Must(template.ParseFiles("main/templates/result.html"))

	final := ""
	for i := 0; i < len(clean_input); i++ {
		if clean_input[i] != "\n" {
			final1 := Draw(clean_input[i], characters)
			for i := 0; i < len(final1); i++ {
				final += final1[i] + "\n"
			}
		} else {
			final += "\n"
		}
	}

	// fmt.Println(final)
	data := map[string]string{
		"Text":   html.EscapeString(input),
		"Banner": r.FormValue("banner"),
		"Output": final,
	}

	tmpl.Execute(w, data)
}
