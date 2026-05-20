package main

import (
	"testing"
)

func TestGetHeadingFrominputBodyBasic(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "get h1",
			inputBody: "<inputBody><body><h1>Test Title</h1></body></inputBody>",
			expected:  "Test Title",
		},
		{
			name:      "get h1 version. 2",
			inputBody: "<inputBody><body><h1>Test h1 Title</h1><h2>Test h2 Title</h2></body></inputBody>",
			expected:  "Test h1 Title",
		},
		{
			name:      "fallback to h2",
			inputBody: "<inputBody><body><h2>Test h2 Title</h2></body></inputBody>",
			expected:  "Test h2 Title",
		},
		{
			name:      "no body",
			inputBody: "<inputBody><body></body></inputBody>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFrominputBody(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %q, actual: %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFrominputBodyMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "get main paragraph",
			inputBody: `<inputBody><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></inputBody>`,
			expected: "Main paragraph.",
		},
		{
			name: "ignore second paragraph",
			inputBody: `<inputBody><body>
		<main>
			<p>Main paragraph.</p>
			<p>ignore this paragraph.</p>
		</main>
	</body></inputBody>`,
			expected: "Main paragraph.",
		},
		{
			name: "no paragraph",
			inputBody: `<inputBody><body>
		<main>
		</main>
	</body></inputBody>`,
			expected: "",
		},
		{
			name: "no main body",
			inputBody: `<inputBody><body>
		<p>Just me and me alone.</p>
	</body></inputBody>`,
			expected: "Just me and me alone.",
		},
		{
			name: "no main body",
			inputBody: `<inputBody><body>
		<main>
			<p>First</p>
			<p>Second</p>
		</main>
	</body></inputBody>`,
			expected: "First",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFrominputBody(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %q, actual: %q", i, tc.name, tc.expected, actual)
			}
		})
	}

}
