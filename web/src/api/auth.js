import request from '@/utils/request'

export function login (parameter) {
  return request({
    url: "auth/login",
    method: 'post',
    data: parameter
  })
}

export function logout () {
  return request({
    url: "auth/logout",
    method: 'post',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}
