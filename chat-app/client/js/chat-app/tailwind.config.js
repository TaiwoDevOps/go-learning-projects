/** @type {import('tailwindcss').Config} */
module.exports = {
   content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        'dark-primary': '#131a1c',
        'dark-secondary': '#1b2224',
        red: '#ff0000',
        green: '#00ff00',
        blue: '#0000ff',
        grey: '#808080',
        white: '#ffffff',
      },
    },
  },
  plugins: [],
}

