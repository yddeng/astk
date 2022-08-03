import request from '@/utils/request'

export function incCreate (parameter) {
  return request({
    url: "inc/create",
    method: 'post',
    data: parameter
  })
}

export function incDelete (parameter) {
  return request({
    url: "inc/delete",
    method: 'post',
    data: parameter
  })
}
export function incList (parameter) {
  return request({
    url: "inc/list",
    method: 'post',
    data: parameter
  })
}
export function incOpened (parameter) {
  return request({
    url: "inc/opened",
    method: 'post',
    data: parameter
  })
}
