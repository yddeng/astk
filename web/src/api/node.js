import request from '@/utils/request'

export function nodeList (parameter) {
  return request({
    url: 'node/list',
    method: 'post',
    data: parameter
  })
}

export function nodeNames (parameter) {
  return request({
    url: 'node/names',
    method: 'post',
    data: parameter
  })
}

export function nodeRemove (parameter) {
  return request({
    url: 'node/remove',
    method: 'post',
    data: parameter
  })
}
