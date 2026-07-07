import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react';

// Backend Go chưa bật CORS, nên dev server proxy /api sang VITE_API_URL
// để trình duyệt luôn gọi same-origin.
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '');
  return {
    plugins: [react()],
    server: {
      proxy: {
        '/api': {
          target: env.VITE_API_URL || 'http://localhost:8080',
          changeOrigin: true,
        },
      },
    },
  };
});
