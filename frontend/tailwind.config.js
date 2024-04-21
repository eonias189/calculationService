/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      gridTemplateColumns: {
        "custom-layout": "auto min-content",
      },
      gridTemplateRows: {
        "cutom-layout": "min-content auto",
      },
      colors: {
        primary: "#0d6efd",
        secondary: "#6c757d",
      },
    },
  },
  plugins: [],
};
