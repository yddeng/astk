import request from '@/utils/request'

export function tags (parameter) {
  return request({
    url: "process/tags",
    method: 'post',
    data: parameter
  })
}

export function processList (parameter) {
  return request({
    url: "process/list",
    method: 'post',
    data: parameter
  })
}

export function processCreate (parameter) {
  return request({
    url: "process/create",
    method: 'post',
    data: parameter
  })
}

export function processUpdate (parameter) {
  return request({
    url: "process/update",
    method: 'post',
    data: parameter
  })
}

export function processDelete (parameter) {
  return request({
    url: "process/delete",
    method: 'post',
    data: parameter
  })
}

export function processStart (parameter) {
  return request({
    url: "process/start",
    method: 'post',
    data: parameter
  })
}

export function processStop (parameter) {
  return request({
    url: "process/stop",
    method: 'post',
    data: parameter
  })
}

export function processBatchStart (parameter) {
  return request({
    url: "process/batch/start",
    method: 'post',
    data: parameter
  })
}

export function processBatchStop (parameter) {
  return request({
    url: "process/batch/stop",
    method: 'post',
    data: parameter
  })
}

export function processTail (parameter) {
  return request({
    url: "process/tail",
    method: 'post',
    data: parameter
  })
}

export function processBell (parameter) {
  return request({
    url: "process/bell",
    method: 'post',
    data: parameter
  })
}
