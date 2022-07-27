
import request from '@/utils/request'

export function monitorInfo (parameter) {
  return request({
    url: "monitor/info",
    method: 'post',
    data: parameter
  })
}

export function monitorRule (parameter) {
  return request({
    url: "monitor/rule",
    method: 'post',
    data: parameter
  })
}

export function monitorNotify (parameter) {
  return request({
    url: "monitor/notify",
    method: 'post',
    data: parameter
  })
}

