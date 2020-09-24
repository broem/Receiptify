// import Vue from 'vue';
import axios from 'axios';

const BASE = 'http://localhost:8086';
const client = axios.create({
    baseURL: BASE,
    json: true
});


const API = {
    convertRaw(b64) {
        return this.do('post', '/api/v1/convertRaw', b64);
    },
    convertCsv(text) {
        return this.do('post', '/api/v1/generateCSV', text);
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