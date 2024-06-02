import http from 'k6/http'
import {check, sleep} from 'k6'
import {Post} from "../api_script.js";

export default function Login(){
    const endpoint = "/auth/login"
    const payload ={
        email: "randysteven12@gmail.com",
        password: "test_1234"
    }

    const res = Post(200, endpoint, null, payload)

    let response = res.json()
    console.log(response.message)
    return response.data.user.token
}