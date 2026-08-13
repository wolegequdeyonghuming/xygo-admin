import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import uni from '@dcloudio/vite-plugin-uni'

export default defineConfig({
  plugins: [uni()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  build: {
    target: 'es2015',
    outDir: '../../deploy/uni'
  },
  server: {
    port: 5173,
    proxy: {
      '/staff': {
        target: 'http://127.0.0.1:4096',
        changeOrigin: true
      },
      '/admin': {
        target: 'http://127.0.0.1:4096',
        changeOrigin: true
      },
      '/attachment': {
        target: 'http://127.0.0.1:4096',
        changeOrigin: true
      },
      '/site': {
        target: 'http://127.0.0.1:4096',
        changeOrigin: true
      }
    }
  }
})
