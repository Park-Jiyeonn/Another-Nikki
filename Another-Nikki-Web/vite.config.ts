import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from "path"
import Components from "unplugin-vue-components/vite";
import { VantResolver } from "unplugin-vue-components/resolvers";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
import { fileURLToPath, URL } from 'node:url'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
		vue(),
		Components({
			resolvers: [VantResolver()],
		}),
		createSvgIconsPlugin({
			// 配置svg文件路径
			iconDirs: [resolve(process.cwd(), "src/assets/svg")],
			// 指定symbolId格式，规则
			symbolId: "icon-[name]",
		}),
	],
	resolve: {
		alias: {
			'@': fileURLToPath(new URL('./src', import.meta.url)),
		},
	},
	css: {
		preprocessorOptions: {
			scss: {
				additionalData: '@use "@/assets/css/base/index.scss" as *;',
			},
		},
	},
	base: '/',
})
