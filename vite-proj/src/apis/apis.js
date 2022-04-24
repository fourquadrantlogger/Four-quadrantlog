import request from './request'

export const Quadrant= [
	'时空',
	'肉体',
	'信息知识',
	'社会关系',
	'持有物'
]
export function getTagList(params) {
	return request({
		url: '/api/tag',
		method: 'get',
		params
	})
}
export function getLogList(params) {
	return request({
		url: '/api/log',
		method: 'get',
		params
	})
}
export function getlog(id) {
	return request({
		url: '/api/log/'+id,
		method: 'get',
	})
}
export function createlog(d) {
	return request({
		url: '/api/log',
		method: 'POST',
		data:d,
	})
}
export function updatelog(d) {
	return request({
		url: '/api/log',
		method: 'PUT',
		data:d,
	})
}
export function deletelog(d) {
	return request({
		url: '/api/log/'+d,
		method: 'DELETE',
		data:d,
	})
}