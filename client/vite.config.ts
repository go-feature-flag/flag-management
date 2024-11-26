/// <reference types="vitest" />
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import { chunkSplitPlugin } from "vite-plugin-chunk-split";

export default defineConfig({
  plugins: [react(), chunkSplitPlugin()],
  test: {
    include: ["src/**/*.{test,spec}.{js,mjs,cjs,ts,mts,cts,jsx,tsx}"],
    globals: true,
    environment: "jsdom",
    setupFiles: "./tests/setup.ts",
    disableConsoleIntercept: true, // This line is to be able to console log in tests
  },
});
