import http from 'k6/http'
import {check, sleep} from 'k6'

export default function(){
    let data = {
        email: 'beatuzcrespo18@gmail.com', password: "test_1234"
    }

    const url = "http://localhost:8000/auth/login/";

    let res = http.post(url, data);

    console.log(res.json())
}