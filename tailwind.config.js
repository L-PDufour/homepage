// tailwind.config.js
/** @type {import('tailwindcss').Config} */

module.exports = {
  content: ["./internal/views/*.{templ,go,html}"],
  theme: {
    extend: {
      colors: {
        base: "#303446", //Base00
        mantle: "#292c3c", //Base01
        surface0: "#414559", //Base02
        surface1: "#51576d", //Base03
        surface2: "#626880", //Base04
        text: "#e0e4fc", // Base05
        rosewater: "#f2d5cf", //Base06
        lavender: "#babbf1", //Base07
        red: "#e78284", //Base08
        peach: "#ef9f76", //Base09
        yellow: "#e5c890", //Base0A
        green: "#a6d189", //Base0B
        teal: "#81c8be", //Base0C
        blue: "#8caaee", //Base0D
        mauve: "#ca9ee6", //Base0E
        flamingo: "#eebebe", //Base0F
      },
      typography: (theme) => ({
        DEFAULT: {
          css: {
            color: theme("colors.text"),
            lineHeight: "1.75",
            a: {
              color: theme("colors.blue"),
              "&:hover": {
                color: theme("colors.mauve"),
              },
            },
            "h1, h2, h3, h4, h5, h6": {
              color: theme("colors.text"),
              fontWeight: "700",
            },
            h1: {
              fontSize: "2.5em",
              color: theme("colors.blue"),
            },
            h2: {
              fontSize: "2em",
              color: theme("colors.teal"),
            },
            h3: {
              fontSize: "1.75em",
              color: theme("colors.lavender"),
            },
            h4: {
              fontSize: "1.5em",
              color: theme("colors.peach"),
            },
            h5: {
              fontSize: "1.25em",
              color: theme("colors.yellow"),
            },
            h6: {
              fontSize: "1.1em",
              color: theme("colors.green"),
            },
            strong: {
              color: theme("colors.rosewater"),
              fontWeight: "700",
            },
            em: {
              color: theme("colors.mauve"),
              fontStyle: "italic",
            },
            code: {
              color: theme("colors.green"),
              backgroundColor: theme("colors.mantle"),
              padding: "0.25rem",
              borderRadius: "0.25rem",
              fontWeight: "500",
            },
            "code::before": {
              content: '""',
            },
            "code::after": {
              content: '""',
            },
            pre: {
              backgroundColor: theme("colors.mantle"),
              color: theme("colors.text"),
              fontSize: "0.875rem",
              padding: "1rem",
            },
            blockquote: {
              borderLeftColor: theme("colors.surface0"),
              color: theme("colors.surface1"),
            },
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/typography"), require("@tailwindcss/forms")],
};
