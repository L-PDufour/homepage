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
        text: "#c6d0f5", //Base05
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
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
