import axios from "axios"

// 创建axios实例，设置配置得默认值
const instance = axios.create({
	baseUrl: '127.0.0.1:10008',   // 这里写接口的http地址，
	timeout: 1000000,  // 设置请求超时的默认值 
})
// 设置请求拦截器
instance.interceptors.request.use(
	config => {
		// 判断当前是否有token，有则在请求头上加上token
		if (localStorage.getItem('token')) {
			config.headers.Authorization = localStorage.getItem('token')
		}
		return config
	},
	error => {
		// 请求错误进行拦截并返回错误信息
		console.log(error)
		return Promise.reject(error)
	}
)
// 设置响应拦截
instance.interceptors.response.use(
	reponse => {
		const res = reponse.data
		return res
	},
	error => {
		return Promise.reject(error)
	}	
)
// 向外暴露axios实例
export default instance
