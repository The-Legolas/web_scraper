package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      PageData
		errorContains string
	}{
		{
			name:     "basic test",
			inputURL: "https://crawler-test.com",
			inputBody: `
<inputBody>
	<body>
	  <h1>Test Title</h1>
	  <p>This is the first paragraph.</p>
	  <a href="/link1">Link 1</a>
	  <img src="/image1.jpg" alt="Image 1">
	</body>
</inputBody>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://crawler-test.com/link1"},
				ImageURLs:      []string{"https://crawler-test.com/image1.jpg"},
			},
		},
		{
			name:     "fallback paragraph when no <main>",
			inputURL: "https://crawler-test.com",
			inputBody: `
<inputBody>
  <body>
    <h1>Title</h1>
    <p>Outside paragraph wins.</p>
    <a href="/x">x</a>
    <img src="/img.png">
  </body>
</inputBody>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Title",
				FirstParagraph: "Outside paragraph wins.",
				OutgoingLinks:  []string{"https://crawler-test.com/x"},
				ImageURLs:      []string{"https://crawler-test.com/img.png"},
			},
		},
		{
			name:     "malformed inputBody still parsed; absolute link and image",
			inputURL: "https://crawler-test.com",
			inputBody: `
<inputBody body>
  <h1>Messy</h1>
  <a href="https://other.com/path">Other</a>
  <img src="https://cdn.boot.dev/banner.jpg">
</inputBody body>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Messy",
				FirstParagraph: "",
				OutgoingLinks:  []string{"https://other.com/path"},
				ImageURLs:      []string{"https://cdn.boot.dev/banner.jpg"},
			},
		},
		{
			name:     "no h1 and no paragraph",
			inputURL: "https://crawler-test.com",
			inputBody: `
<inputBody>
  <body>
    <a href="/only-link">Only link</a>
    <img src="/only.png">
  </body>
</inputBody>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "",
				FirstParagraph: "",
				OutgoingLinks:  []string{"https://crawler-test.com/only-link"},
				ImageURLs:      []string{"https://crawler-test.com/only.png"},
			},
		},
		{
			name:     "multiple links and images preserve order",
			inputURL: "https://crawler-test.com",
			inputBody: `
<inputBody><body>
  <h1>t</h1>
  <main><p>p</p></main>
  <a href="/a1">a1</a>
  <a href="https://x.dev/a2">a2</a>
  <img src="/i1.png">
  <img src="https://x.dev/i2.png">
</body></inputBody>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "t",
				FirstParagraph: "p",
				OutgoingLinks: []string{
					"https://crawler-test.com/a1",
					"https://x.dev/a2",
				},
				ImageURLs: []string{
					"https://crawler-test.com/i1.png",
					"https://x.dev/i2.png",
				},
			},
		},
		{
			name:     "invalid base URL → empty link/image slices",
			inputURL: `:\\invalidBaseURL`,
			inputBody: `
<inputBody>
  <body>
    <h1>Title</h1>
    <p>Paragraph</p>
    <a href="/path">path</a>
    <img src="/logo.png">
  </body>
</inputBody>`,
			expected: PageData{
				URL:            `:\\invalidBaseURL`,
				Heading:        "Title",
				FirstParagraph: "Paragraph",
				OutgoingLinks:  nil,
				ImageURLs:      nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			actual := extractPageData(tc.inputBody, tc.inputURL)

			if actual.URL != tc.expected.URL {
				t.Errorf("URL: expected %q, actual %q", tc.expected.URL, actual.URL)
			}
			if actual.Heading != tc.expected.Heading {
				t.Errorf("Heading: expected %q, actual %q", tc.expected.Heading, actual.Heading)
			}
			if actual.FirstParagraph != tc.expected.FirstParagraph {
				t.Errorf("FirstParagraph: expected %q, actual %q", tc.expected.FirstParagraph, actual.FirstParagraph)
			}
			if !reflect.DeepEqual(actual.OutgoingLinks, tc.expected.OutgoingLinks) {
				t.Errorf("OutgoingLinks: expected %v, actual %v", tc.expected.OutgoingLinks, actual.OutgoingLinks)
			}
			if !reflect.DeepEqual(actual.ImageURLs, tc.expected.ImageURLs) {
				t.Errorf("ImageURLs: expected %v, actual %v", tc.expected.ImageURLs, actual.ImageURLs)
			}
		})
	}
}
