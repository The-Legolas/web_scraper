package main

import (
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "get h1",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "get h1 version. 2",
			inputBody: "<html><body><h1>Test h1 Title</h1><h2>Test h2 Title</h2></body></html>",
			expected:  "Test h1 Title",
		},
		{
			name:      "fallback to h2",
			inputBody: "<html><body><h2>Test h2 Title</h2></body></html>",
			expected:  "Test h2 Title",
		},
		{
			name:      "no body",
			inputBody: "<html><body></body></html>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %q, actual: %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "get main paragraph",
			inputBody: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "ignore second paragraph",
			inputBody: `<html><body>
		<main>
			<p>Main paragraph.</p>
			<p>ignore this paragraph.</p>
		</main>
	</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "no paragraph",
			inputBody: `<html><body>
		<main>
		</main>
	</body></html>`,
			expected: "",
		},
		{
			name: "no main body",
			inputBody: `<html><body>
		<p>Just me and me alone.</p>
	</body></html>`,
			expected: "Just me and me alone.",
		},
		{
			name: "no main body",
			inputBody: `<html><body>
		<main>
			<p>First</p>
			<p>Second</p>
		</main>
	</body></html>`,
			expected: "First",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %q, actual: %q", i, tc.name, tc.expected, actual)
			}
		})
	}

}
