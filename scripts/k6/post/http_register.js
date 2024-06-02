import http from 'k6/http'
import {check, sleep} from 'k6'
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

const firstNameChars = 'abcdefghijklmnopqrstuvwxyz';
const lastNameChars = firstNameChars + 'ABCDEFGHIJKLMNOPQRSTUVWXYZ-';
const number = '1234567890'

function randomFirstName() {
    const firstNameLength = Math.floor(Math.random() * 6) + 5; // Random length between 5 and 10 characters
    const firstName = randomString(firstNameLength, firstNameChars);
    return `${firstName}`;
}

function randomLastName() {
    const lastNameLength = Math.floor(Math.random() * 8) + 7; // Random length between 7 and 15 characters
    const lastName = randomString(lastNameLength, lastNameChars);
    return `${lastName}`;
}

function randomUserName() {
    const numberLength = Math.floor(Math.random() * 6) + 10;
    const numberGet = randomString(numberLength, number);
    return `${randomFirstName()}_${randomLastName()}_${numberGet}`;
}

function randomEmail() {
    return `${randomFirstName()}.${randomLastName()}@gmail.com`;
}

export default function(){
    const url = "http://localhost:8080/auth/register"
    const payload = JSON.stringify({
        first_name: randomFirstName(),
        last_name: randomLastName(),
        user_name: randomUserName(),
        email: randomEmail(),
        gender: 'male',
        password: "test_1234",
        date_of_birth: '2001-01-01'
    })

    const param = {
        headers: {
            'Content-Type': 'application/json'
        }
    }

    const res = http.post(url, payload, param)

    check(res, {
            'Post status is 201 ': (r) => res.status === 201
        }
    )
    console.log(res)
}