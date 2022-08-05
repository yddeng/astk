
import request from '@/utils/request'

export function monitorInfo (parameter) {
  return request({
    url: "monitor/info",
    method: 'post',
    data: parameter
  })
}

export function monitorUpdate (parameter) {
  return request({
    url: "monitor/update",
    method: 'post',
    data: parameter
  })
}

export function monitorOpened (parameter) {
  return request({
    url: "monitor/opened",
    method: 'post',
    data: parameter
  })
}

