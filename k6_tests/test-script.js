import http from 'k6/http'
import { sleep } from 'k6'

export const options = {
   vus: 100,
   duration: '30s'
}

function generateRandomEmail() {
    const chars = 'abcdefghijklmnopqrstuvwxyz0123456789';
    let randomString = '';
    for (let i = 0; i < 10; i++) {
        randomString += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return `user_${randomString}@web.com`;
}

export default function() {
    const randomEmail = generateRandomEmail();

    const payload = JSON.stringify({
        "email": randomEmail,
        "password": "password#9090"
    })

    const headers = {'Content-Type': 'application/json'}
    http.post('http://localhost:3001/api/users',payload,{headers:headers})
    
}