const flowbite = require("flowbite-react/tailwind");

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
    flowbite.content(),
    "./node_modules/react-tailwindcss-select/dist/index.esm.js",
    "./node_modules/react-tailwindcss-datetimepicker/dist/react-tailwindcss-datetimepicker.js",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: "#eff6ff",
          100: "#dbeafe",
          200: "#bfdbfe",
          300: "#93c5fd",
          400: "#60a5fa",
          500: "#3b82f6",
          600: "#2563eb",
          700: "#1d4ed8",
          800: "#1e40af",
          900: "#1e3a8a",
        },
        goff: {
          50: "#edfcf7",
          100: "#cdf7e7",
          200: "#abefd9",
          300: "#74e1c4",
          400: "#3ccbaa",
          500: "#18b192",
          600: "#0c8f77",
          700: "#0a7263",
          800: "#0a5b4f",
          900: "#0a4a41",
          950: "#042a26",
        },
        variationColor: {
          50: "#1f77b4",
          100: "#ff7f0e",
          150: "#2ca02c",
          200: "#d62728",
          250: "#9467bd",
          300: "#8c564b",
          400: "#e377c2",
          450: "#7f7f7f",
          500: "#bcbd22",
          550: "#17becf",
          600: "#aec7e8",
          650: "#ffbb78",
          700: "#98df8a",
          750: "#ff9896",
          800: "#c5b0d5",
          850: "#c49c94",
          900: "#f7b6d2",
          950: "#c7c7c7",
        },
      },
      fontFamily: {
        sans: [
          "Inter",
          "ui-sans-serif",
          "system-ui",
          "-apple-system",
          "system-ui",
          "Segoe UI",
          "Roboto",
          "Helvetica Neue",
          "Arial",
          "Noto Sans",
          "sans-serif",
          "Apple Color Emoji",
          "Segoe UI Emoji",
          "Segoe UI Symbol",
          "Noto Color Emoji",
        ],
        body: [
          "Inter",
          "ui-sans-serif",
          "system-ui",
          "-apple-system",
          "system-ui",
          "Segoe UI",
          "Roboto",
          "Helvetica Neue",
          "Arial",
          "Noto Sans",
          "sans-serif",
          "Apple Color Emoji",
          "Segoe UI Emoji",
          "Segoe UI Symbol",
          "Noto Color Emoji",
        ],
        mono: [
          "ui-monospace",
          "SFMono-Regular",
          "Menlo",
          "Monaco",
          "Consolas",
          "Liberation Mono",
          "Courier New",
          "monospace",
        ],
      },
    },
  },
  plugins: [flowbite.plugin()],
};