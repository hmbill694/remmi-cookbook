/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./cmd/web/**/*.html",
    "./cmd/web/**/*.templ",
    "./internal/views/**/*.html",
    "./internal/views/**/*.templ",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
