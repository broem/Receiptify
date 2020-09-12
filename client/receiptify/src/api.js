import Vue from 'vue';
import axios from 'axios';

const BASE = 'http://localhost:8086';
const client = axios.create({
    baseURL: BASE,
    json: true
});


const API = {
    base64string(b64) {
        return this.do('post', '/base64string', b64);
    },

    async do(method, from, data) {
        return client({
            method,
            url: from,
            data
        }).then(req => {
            return req
        })
    }
}

export default API;