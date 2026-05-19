package main

import "fmt"

func main() {
	p := `
<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>
`
	_, _ = normalizeURL(p)
	val := getHeadingFromHTML(p)
	fmt.Println(val)
}
