package main

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
	"net/http"
)

func generateWebContent() string {
	// Initialize StyleManager
	styleMgr := styles.NewStyleManager()

	// Button with hover effect
	buttonClass := styleMgr.AddCompositeStyle(styles.CompositeStyle{
		Default: styles.Props{
			styles.Padding:         "10px 20px",
			styles.BackgroundColor: "blue",
			styles.Color:           "white",
			styles.Border:          "none",
			styles.Cursor:          "pointer",
			styles.Margin:          "10px",
		},
		PseudoClasses: map[string]styles.Props{
			"hover": {
				styles.BackgroundColor: "darkblue",
			},
		},
	})

	// Animated element
	animationName := styleMgr.AddAnimation(styles.Keyframes{
		"0%":   {styles.Transform: "translateY(0px)"},
		"50%":  {styles.Transform: "translateY(-20px)"},
		"100%": {styles.Transform: "translateY(0px)"},
	})
	animatedClass := styleMgr.AddStyle(styles.Props{
		styles.Width:                   "100px",
		styles.Height:                  "100px",
		styles.BackgroundColor:         "green",
		styles.AnimationName:           animationName,
		styles.AnimationDuration:       "2s",
		styles.AnimationIterationCount: "infinite",
	})

	// Responsive design
	responsiveClass := styleMgr.AddCompositeStyle(styles.CompositeStyle{
		Default: styles.Props{
			styles.Padding:         "20px",
			styles.BackgroundColor: "lightgray",
			styles.Margin:          "10px",
		},
		MediaQueries: map[string]styles.Props{
			"@media (max-width: 600px)": {
				styles.Padding:         "10px",
				styles.BackgroundColor: "lightblue",
			},
		},
	})

	pseudoElementClass := styleMgr.AddCompositeStyle(styles.CompositeStyle{
		Default: styles.Props{
			styles.Color:           "black",
			styles.BackgroundColor: "white",
			styles.Padding:         "10px",
			styles.Margin:          "10px",
			styles.Border:          "1px solid gray",
			styles.Position:        "relative",
		},
		PseudoElements: map[string]styles.Props{
			styles.PseudoBefore: {
				styles.Content:    `"Before "`,
				styles.Color:      "red",
				styles.FontWeight: "bold",
				styles.Position:   "absolute",
				styles.Left:       "10px",
				styles.Top:        "-20px",
			},
			styles.PseudoAfter: {
				styles.Content:    `" After"`,
				styles.Color:      "blue",
				styles.FontWeight: "bold",
				styles.Position:   "absolute",
				styles.Right:      "10px",
				styles.Bottom:     "-20px",
			},
		},
	})

	// Composing the page
	pageContent := elem.Div(nil,
		elem.Button(attrs.Props{attrs.Class: buttonClass}, elem.Text("Hover Over Me")),
		elem.Div(attrs.Props{attrs.Class: animatedClass}, elem.Text("I animate!")),
		elem.Div(attrs.Props{attrs.Class: responsiveClass}, elem.Text("Resize the window")),
		elem.Div(attrs.Props{attrs.Class: pseudoElementClass}, elem.Text("I have pseudo-elements")),
	)

	// Render with StyleManager
	return pageContent.RenderWithOptions(elem.RenderOptions{StyleManager: styleMgr})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlContent := generateWebContent() // Assume this returns the HTML string
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(htmlContent))
	})

	http.ListenAndServe(":8080", nil)
}
