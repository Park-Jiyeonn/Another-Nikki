import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from "path"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],

  server: {
		open: true,
		port: 9999,
		host: '0.0.0.0',
	},

  // 设置路径别名
  resolve: {
    alias:{
      '@': resolve(__dirname, 'src'),
    }
}
})
