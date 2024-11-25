/// <reference types="vitest" />
import react from "@vitejs/plugin-react";
import {defineConfig} from "vite";
import {chunkSplitPlugin} from "vite-plugin-chunk-split";

export default defineConfig({
    plugins: [react(), chunkSplitPlugin()],
    test: {},
});
