import request from '@/utils/request'

export function nodeList (parameter) {
  return request({
    url: 'node/list',
    method: 'post',
    data: parameter
  })
}

export function nodeStatus (parameter) {
  return request({
    url: 'node/status',
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

export function nodeBell (parameter) {
  return request({
    url: "node/bell",
    method: 'post',
    data: parameter
  })
}

