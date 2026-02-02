// http://localhost:8080/
package main

import (
	"html/template" // 引入處理 HTML 模板的工具
	"net/http"      // 引入網路工具
)

/* ----------------------------------------------------------- */

// 這原本是你寫好的首頁函式
func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	// 告訴 Go：當有人要看「關於我」，請讀取 about.html
	t, _ := template.ParseFiles("templates/about.html")
	t.Execute(w, nil)
}

func projects(w http.ResponseWriter, r *http.Request) {
	// 告訴 Go：當有人要看「作品」，請讀取 projects.html
	t, _ := template.ParseFiles("templates/projects.html")
	t.Execute(w, nil)
}

func awards(w http.ResponseWriter, r *http.Request) {
	// 告訴 Go：當有人要看「獎項」，請讀取 awards.html
	t, _ := template.ParseFiles("templates/awards.html")
	t.Execute(w, nil)
}

/* ----------------------------------------------------------- */

func main() {

	// 讓 Go 知道：如果網址開頭是 /static/，就去 static 資料夾找檔案
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 告示牌：輸入 / 就去執行 home 函式
	http.HandleFunc("/", home)

	// 告示牌：輸入 /about 就去執行 about 函式
	http.HandleFunc("/about", about)

	// 告示牌：輸入 /projects 就去執行 projects 函式
	http.HandleFunc("/projects", projects)

	// 告示牌：輸入 /awards 就去執行 awards 函式
	http.HandleFunc("/awards", awards)

	println("伺服器已啟動：http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
