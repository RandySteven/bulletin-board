import http from 'k6/http'
import {check, sleep} from 'k6'

export const Post = (
    expectedResponseCode,
    endpoint,
    jwtToken,
    payload
) => {
    const defaultHost = 'http://localhost:8080'
    let url = defaultHost + endpoint
    const reqBody = JSON.stringify(payload)

    let param = {
        headers: {
            'Content-Type': 'application/json'
        }
    }

    if (jwtToken !== null) {
        param.headers.Authorization = `Bearer ${jwtToken}`
    }
    let response = http.post(url, reqBody, param)

    check(response, {
        'The response code is expected': (r) => response.status === expectedResponseCode
    })

    return response
}

