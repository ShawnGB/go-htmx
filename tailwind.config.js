const config = {
  content: ["./ui/**/*.tmpl", "./ui/**/*.html", "./ui/**/*.gohtml"],
  theme: {
    colors: {
      offWhite: "#f8f8f8", // Adjust the color code as needed
      primary: {
        50: "#fce4e4",
        100: "#facccc",
        200: "#f99a9a",
        300: "#f76767",
        400: "#f53535",
        500: "#ef3238", // Primary brand color
        600: "#e52b2f",
        700: "#d02125",
        800: "#ae1a1d",
        900: "#8c1214",
      },
      secondary: {
        50: "#f9f9f9",
        100: "#f0f0f0",
        200: "#d9d9d9",
        300: "#bfbfbf",
        400: "#a6a6a6",
        500: "#8c8c8c",
        600: "#737373",
        700: "#595959",
        800: "#404040",
        900: "#262626",
      },
    },
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
    },
  },
  plugins: [],
};
export default config;
