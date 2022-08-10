import request from '@/utils/request'

export function gitCreate (parameter) {
  return request({
    url: "githook/create",
    method: 'post',
    data: parameter
  })
}

export function gitDelete (parameter) {
  return request({
    url: "githook/delete",
    method: 'post',
    data: parameter
  })
}
export function gitList (parameter) {
  return request({
    url: "githook/list",
    method: 'post',
    data: parameter
  })
}
