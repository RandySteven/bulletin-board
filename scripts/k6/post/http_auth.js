import http from 'k6/http'
import {check, sleep} from 'k6'

export default function(){
    const url = "http://localhost:8080/auth/login"
    const payload = JSON.stringify({
        email: "randysteven12@gmail.com",
        password: "test_1234"
    })

    const param = {
        headers: {
            'Content-Type': 'application/json'
        }
    }

    const res = http.post(url, payload, param)

    check(res, {
            'Post status is 200 ': (r) => res.status === 200
        }
    )
    console.log(res)
}