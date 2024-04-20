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
    },
  },
  plugins: [],
};
