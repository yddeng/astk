import request from '@/utils/request'

/**
 * login func
 * parameter: {
 *     username: '',
 *     password: '',
 *     remember_me: true,
 *     captcha: '12345'
 * }
 * @param parameter
 * @returns {*}
 */
export function cmdList (parameter) {
  return request({
    url: 'cmd/list',
    method: 'post',
    data: parameter
  })
}

export function cmdCreate (parameter) {
  return request({
    url: 'cmd/create',
    method: 'post',
    data: parameter
  })
}

export function cmdDelete (parameter) {
  return request({
    url: 'cmd/delete',
    method: 'post',
    data: parameter
  })
}

export function cmdUpdate (parameter) {
  return request({
    url: 'cmd/update',
    method: 'post',
    data: parameter
  })
}

export function cmdExec (timeout, parameter) {
  return request({
    url: 'cmd/exec',
    method: 'post',
    timeout: timeout,
    data: parameter
  })
}

export function cmdLog (parameter) {
  return request({
    url: 'cmd/log',
    method: 'post',
    data: parameter
  })
}
