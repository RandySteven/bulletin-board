import {check, sleep} from 'k6'
import Login from "./http_auth.js";
import {Post} from "../api_script.js";

export default function CreateTask () {
    const endpoint = "/tasks"
    let token = Login()
    let payload = {
        task: {
            title: 'Kucing PoppyMAN',
            description: 'Kucing PoppyMAN adalah kucing yang dibentuk oleh legenda hidup',
            image: '',
            expired_date: '2025-01-01'
        },
        reward: {
            name: '2 Malam',
            image: '',
            description: 'Aju 1 malam kan aku sih 2 malam lah',
        },
        categories: {
            ids: [
                1
            ]
        }
    }

    const res = Post(201, endpoint, token, payload)
    console.log(res)
}