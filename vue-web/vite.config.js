import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path' // 主要用于alias文件路径别名
const pathResolve = (dir) => resolve(__dirname, '.', dir)

// https://vitejs.dev/config/

// 开发模式、生产模式
const project = (url) => {
  return defineConfig({
    plugins: [vue()],
    devtools: false,
    resolve: {
      alias: {
        '/@': resolve(__dirname, '.', 'src')
      }
    },
    base: url,
    // 打包配置
    build: {
      sourcemap: true,
      outDir: 'distp', //指定输出路径
      assetsDir: 'static/img/', // 指定生成静态资源的存放路径
      rollupOptions: {
        output: {
          manualChunks(id) {
            if (id.includes('node_modules')) {
              const arr = id.toString().split('node_modules/')[1].split('/')
              switch(arr[0]) {
                // case '@kangc': // md 编辑器
                // case 'axios':
                case '@ctrl':
                case '@floating-ui':
                case 'dayjs':
                case '@vueuse':
                case '@popperjs':
                case '@vue':
                case 'element-plus': // UI 库
                case '@element-plus': // 图标
                  return '_' + arr[0]
                  break
                case '@naturefw': // 自然框架
                  return '_' + arr[0] + '_' + arr[1]
                  break
                default :
                  return '__vendor'
                  break
              }
            }
          },
          chunkFileNames: 'static/js1/[name]-[hash].js',
          entryFileNames: 'static/js2/[name]-[hash].js',
          assetFileNames: 'static/[ext]/[name]-[hash].[ext]'
        },
        brotliSize: false, // 不统计
        target: 'esnext', 
        minify: 'esbuild' // 混淆器，terser构建后文件体积更小
      }
    }
  })
}


export default ({ mode }) => {
  const url = loadEnv(mode, process.cwd()).VITE_BASEURL
  switch (url) {
    case 'lib': // 打包库文件
      return lib
      break;
    
    default: // 开发模式、生产模式
      return project(url)
      break;
  }
}